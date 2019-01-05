package grpc

import (
	"context"
	"fmt"
	"math/rand"

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
}

func NewClient() (gameserver.GameServer, error) {
	// entities := []entity.Entity{
	// 	entity.Entity{NewID(), entity.Coin, types.Position{types.Coord{2, 1}, 0}},
	// 	entity.Entity{NewID(), entity.Coin, types.Position{types.Coord{8, 2}, 0}},
	// }

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

	return &GrpcClient{
		client:   client,
		world:    w,
		entities: []entity.Entity{},
	}, nil
}

func (s *GrpcClient) PerformAction(e entity.Entity, p types.Position) (entity.Entity, error) {
	return entity.Entity{}, nil
}

func (s *GrpcClient) NewPlayer() (entity.Entity, error) {
	resp, err := s.client.NewPlayer(context.Background(), &pb.Empty{})
	if err != nil {
		return entity.Entity{}, err
	}
	e := entity.Entity{
		resp.GetID(),
		entity.Character,
		types.Position{types.Coord{rand.Intn(3), rand.Intn(3)}, 0},
	}
	return e, nil
}

func (s *GrpcClient) Entities() []entity.Entity {
	return s.entities
}

func (s *GrpcClient) World() types.World {
	return s.world
}
