package grpc

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/kyeett/gameserver/localserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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

// Todo clean up
func (s *GrpcServer) Run() {
	port := ":10001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ss := grpc.NewServer()
	pb.RegisterBackendServer(ss, s)
	// Register reflection service on gRPC server.
	reflection.Register(ss)
	if err := ss.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *GrpcServer) NewPlayer(ctx context.Context, _ *pb.Empty) (*pb.PlayerID, error) {
	p, err := s.local.NewPlayer()
	if err != nil {
		return nil, err
	}

	return &pb.PlayerID{ID: p.ID}, nil
}

// Todo, decide format to send over wire
func (s *GrpcServer) WorldRequest(ctx context.Context, _ *pb.Empty) (*pb.WorldResponse, error) {

	return &pb.WorldResponse{
		Tiles:  s.local.World().TileBytes(),
		Width:  int32(s.local.World().Width),
		Height: int32(s.local.World().Height),
	}, nil

}
func (s *GrpcServer) EntityStream(_ *pb.Empty, stream pb.Backend_EntityStreamServer) error {

	i := 0
	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		i++
		<-ticker.C

		e := s.local.Entities()
		payload, err := GobMarshal(&e)
		if err != nil {
			return err
		}

		stream.Send(&pb.EntityResponse{
			Payload: payload,
		})
	}
}
