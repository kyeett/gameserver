package grpc

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/kyeett/gameserver"
	"github.com/kyeett/gameserver/entity"

	"github.com/kyeett/gameserver/types"

	pb "github.com/kyeett/gameserver/proto"
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

var cert = "../cert.pem"

func NewClient(serverAddr string, secure bool) (gameserver.GameServer, error) {
	log.SetLevel(log.DebugLevel)
	ctx := context.Background() // Todo move this out
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)

	opts := []grpc.DialOption{}

	switch secure {
	case true:
		creds, err := credentials.NewClientTLSFromFile(cert, "")
		if err != nil {
			return nil, errors.Wrapf(err, "could not load tls cert: %s", cert)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	case false:
		opts = append(opts, grpc.WithInsecure())
	}
	// opts = append(opts, grpc.WithInsecure())

	log.Debugf("dialing to %s\n", serverAddr)
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "new client")
	}

	client := pb.NewBackendClient(conn)
	resp, err := client.WorldRequest(ctx, &pb.Empty{}) // grpc.WaitForReady(true)
	if err != nil {
		log.Debugf("client failed to connect to %s\ndebug: %s", serverAddr, err)
		return nil, errors.Errorf("client failed to connect to %s", serverAddr)
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
	t := time.Now()
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	log.Debug("created a new player in", time.Since(t))
	resp, err := s.client.NewPlayer(ctx, &pb.Empty{})
	if err != nil {
		return entity.Entity{}, err
	}

	// Wait until the new player is available
	log.Debug("Wait until new player is available")
	ticker := time.NewTicker(2 * time.Millisecond)
	for {

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
