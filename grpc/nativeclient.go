// +build !js

package grpc

import (
	"context"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kyeett/gameserver/entity"
	pb "github.com/kyeett/gameserver/proto"
	"github.com/kyeett/gameserver/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var cert = "../insecure/cert.pem"

type GrpcClient struct {
	world     types.World
	entities  []entity.Entity
	client    pb.BackendClient
	mu        *sync.RWMutex
	once      sync.Once
	startedCh chan struct{} // Todo, better method for knowing if started?
}

func CorrectClient(serverAddr string, secure bool) (pb.BackendClient, error) {

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

	return pb.NewBackendClient(conn), nil
}

func getEmpty() *empty.Empty {
	return &empty.Empty{}
}

//Todo: make common between jsclient and nativeclient
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
