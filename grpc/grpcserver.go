package grpc

import (
	"context"

	"github.com/kyeett/gameserver/localserver"

	pb "github.com/kyeett/gameserver/grpc/proto"
)

var _ pb.BackendServer = (*GrpcServer)(nil)

type GrpcServer struct {
	local *localserver.LocalServer
}

func NewServer() (*GrpcServer, error) {
	l := localserver.New().(*localserver.LocalServer)

	return &GrpcServer{
		l,
	}, nil
}

func (s *GrpcServer) NewPlayer(ctx context.Context, _ *pb.Empty) (*pb.PlayerID, error) {
	return &pb.PlayerID{ID: "123asd"}, nil
}
func (s *GrpcServer) WorldRequest(ctx context.Context, _ *pb.Empty) (*pb.WorldResponse, error) {
	return &pb.WorldResponse{
		Tiles: []byte{}, Width: 1, Height: 1,
	}, nil

}
func (s *GrpcServer) EntityStream(_ *pb.Empty, strream pb.Backend_EntityStreamServer) error {
	return nil
}
