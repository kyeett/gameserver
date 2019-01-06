package grpc

import (
	"context"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/localserver"
	"github.com/kyeett/gameserver/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/kyeett/gameserver/grpc/proto"
)

var _ pb.BackendServer = (*GrpcServer)(nil)

type GrpcServer struct {
	local *localserver.LocalServer
	mu    *sync.RWMutex
}

func NewServer(w types.World) (*GrpcServer, error) {
	log.Info("new remote server created")
	l := localserver.New(w).(*localserver.LocalServer)

	return &GrpcServer{
		l,
		&sync.RWMutex{},
	}, nil
}

// Todo clean up
func (s *GrpcServer) Run(ctx context.Context) {

	port := ":10001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ss := grpc.NewServer()
	pb.RegisterBackendServer(ss, s)
	// Register reflection service on gRPC server.
	reflection.Register(ss)

	go func() {
		<-ctx.Done()
		ss.GracefulStop()
	}()

	log.Infof("starting server at %s", port)
	if err := ss.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *GrpcServer) NewPlayer(ctx context.Context, _ *pb.Empty) (*pb.PlayerID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
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
	ticker := time.NewTicker(20 * time.Millisecond)
	for {
		// log.Debug("push entities")
		i++
		<-ticker.C

		s.mu.Lock()
		e := s.local.Entities()
		payload, err := GobMarshal(&e)
		s.mu.Unlock()
		if err != nil {
			return err
		}

		stream.Send(&pb.EntityResponse{
			Payload: payload,
		})
	}
}

//Todo: refactor this code.
func (s *GrpcServer) PerformAction(ctx context.Context, req *pb.ActionRequest) (*pb.ActionResponse, error) {
	tmpEntity := req.GetEntity()

	s.mu.Lock()
	defer s.mu.Unlock()
	e, err := s.local.PerformAction(entity.Entity{ID: tmpEntity.GetID()},
		types.Position{
			Coord: types.Coord{X: int(tmpEntity.GetX()), Y: int(tmpEntity.GetY())},
			Theta: int(tmpEntity.GetTheta())})
	if err != nil {
		return nil, err
	}
	ent := pb.Entity{
		ID:    e.ID,
		X:     int32(e.Position.X),
		Y:     int32(e.Position.Y),
		Theta: int32(e.Position.Theta),
	}
	return &pb.ActionResponse{Entity: &ent}, nil
}
