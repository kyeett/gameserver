package grpc

import (
	"context"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/kyeett/gameserver"
	"github.com/kyeett/gameserver/entity"
	pb "github.com/kyeett/gameserver/grpc/proto"
	"github.com/kyeett/gameserver/types"

	log "github.com/sirupsen/logrus"
)

// Ensure struct implements interface
var _ gameserver.GameServer = (*GrpcClient)(nil)

type GrpcClient struct {
	world     types.World
	entities  []entity.Entity
	client    pb.BackendClient
	mu        *sync.RWMutex
	once      sync.Once
	startedCh chan struct{} // Todo, better method for knowing if started?
}

func NewClient(serverAddr string) (gameserver.GameServer, error) {

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewBackendClient(conn)
	resp, err := client.WorldRequest(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}

	var ts []types.Tile
	for _, b := range resp.GetTiles() {
		ts = append(ts, types.Tile(b))
	}
	w := types.NewWorld(ts, int(resp.GetWidth()), int(resp.GetHeight()))

	c := GrpcClient{
		client:    client,
		world:     w,
		entities:  []entity.Entity{},
		mu:        &sync.RWMutex{},
		startedCh: make(chan struct{}),
	}

	go c.recieveEntityUpdates()

	<-c.startedCh
	log.Debugf("Client connected successfully\n")
	return &c, nil
}

func (s *GrpcClient) recieveEntityUpdates() {
	log.Debug("start entity stream")
	stream, err := s.client.EntityStream(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var entities []entity.Entity
		err = GobUnmarshal(resp.GetPayload(), &entities)
		if err != nil {
			log.Fatal(err)
		}

		s.mu.Lock()
		s.entities = entities
		s.mu.Unlock()

		s.once.Do(func() {
			close(s.startedCh)
		})
	}
}

func (s *GrpcClient) PerformAction(e entity.Entity, p types.Position) (entity.Entity, error) {
	log.Debugf("%s, perform action", e.ID)
	ent := pb.Entity{
		ID:    e.ID,
		X:     int32(p.X),
		Y:     int32(p.Y),
		Theta: int32(p.Theta),
	}

	resp, err := s.client.PerformAction(context.Background(), &pb.ActionRequest{Entity: &ent})
	if err != nil {
		return entity.Entity{}, err
	}

	return entity.Entity{
		ID: resp.Entity.GetID(),
		Position: types.Position{
			types.Coord{int(resp.Entity.GetX()),
				int(resp.Entity.GetY())},
			int(resp.Entity.GetTheta()),
		},
		Owner: "",
	}, nil
}

func (s *GrpcClient) NewPlayer() (entity.Entity, error) {
	log.Debug("create new player")
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
	resp, err := s.client.NewPlayer(ctx, &pb.Empty{})
	if err != nil {
		return entity.Entity{}, err
	}

	// Wait until the new player is available
	log.Debug("Wait until new player is available")
	ticker := time.NewTicker(20 * time.Millisecond)
	for {
		log.Debug("Trololo")

		s.mu.RLock()
		for _, e := range s.entities {
			if resp.GetID() == e.ID {
				s.mu.RUnlock()
				return e, nil
			}
		}
		s.mu.RUnlock()
		<-ticker.C
	}
}

func (s *GrpcClient) Entities() []entity.Entity {
	s.mu.RLock()
	e := s.entities
	s.mu.RUnlock()
	return e
}

func (s *GrpcClient) World() types.World {
	return s.world
}
