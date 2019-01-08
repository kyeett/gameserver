// +build js

package grpc

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/kyeett/gameserver/entity"
	pbweb "github.com/kyeett/gameserver/proto/web"
	"github.com/kyeett/gameserver/types"

	webempty "github.com/johanbrandhorst/protobuf/ptypes/empty"
)

type GrpcClient struct {
	world     types.World
	entities  []entity.Entity
	client    pbweb.BackendClient
	mu        *sync.RWMutex
	once      sync.Once
	startedCh chan struct{} // Todo, better method for knowing if started?
}

func CorrectClient(serverAddr string, secure bool) (pbweb.BackendClient, error) {

	clas := pbweb.NewBackendClient(serverAddr)
	return clas, nil
}

func getEmpty() *webempty.Empty {
	return &webempty.Empty{}
}

func (s *GrpcClient) PerformAction(e entity.Entity, p types.Position) (entity.Entity, error) {
	log.Debugf("%s, perform action", e.ID)
	ent := pbweb.Entity{
		ID:    e.ID,
		X:     int32(p.X),
		Y:     int32(p.Y),
		Theta: int32(p.Theta),
	}

	resp, err := s.client.PerformAction(context.Background(), &pbweb.ActionRequest{Entity: &ent})
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
