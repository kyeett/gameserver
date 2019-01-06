package grpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/kyeett/gameserver"
	"github.com/kyeett/gameserver/entity"
	pb "github.com/kyeett/gameserver/grpc/proto"
	"github.com/kyeett/gameserver/types"
)

// Ensure struct implements interface
var _ gameserver.GameServer = (*GrpcClient)(nil)

type GrpcClient struct {
	world    types.World
	entities []entity.Entity
	client   pb.BackendClient
	mu       *sync.RWMutex
}

func NewClient() (gameserver.GameServer, error) {
	port := 10001
	serverAddr := fmt.Sprintf("localhost:%d", port)
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
		client:   client,
		world:    w,
		entities: []entity.Entity{},
		mu:       &sync.RWMutex{},
	}

	go c.recieveEntityUpdates()

	fmt.Println(w, resp.GetTiles())

	return &c, nil
}

func (s *GrpcClient) recieveEntityUpdates() {
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
	}
}

func (s *GrpcClient) PerformAction(e entity.Entity, p types.Position) (entity.Entity, error) {
	ent := pb.Entity{
		ID:    e.ID,
		X:     int32(p.X),
		Y:     int32(p.Y),
		Theta: int32(p.Theta),
	}
	fmt.Println("Perform action client", e)
	fmt.Println("Perform action client", ent)

	resp, err := s.client.PerformAction(context.Background(), &pb.ActionRequest{Entity: &ent})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

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
	resp, err := s.client.NewPlayer(context.Background(), &pb.Empty{})
	if err != nil {
		return entity.Entity{}, err
	}

	// Wait until the new player is available
	ticker := time.NewTicker(20 * time.Millisecond)
	for {

		s.mu.RLock()
		for _, e := range s.entities {
			if resp.GetID() == e.ID {
				return e, nil
			}
		}
		s.mu.RUnlock()
		<-ticker.C
	}

	// e := entity.Entity{
	// 	resp.GetID(),
	// 	entity.Character,
	// 	types.Position{types.Coord{rand.Intn(3), rand.Intn(3)}, 0},
	// 	"",
	// }
	// s.world
	// return e, nil
}

func (s *GrpcClient) Entities() []entity.Entity {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.entities
}

func (s *GrpcClient) World() types.World {
	return s.world
}
