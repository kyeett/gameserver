// +build js

package grpc

import (
	"context"
	"net/url"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"honnef.co/go/js/dom"

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

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func init() {
	document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
		log.Error("Document loaded!")
	})
}

func CorrectClient(serverAddr string, secure bool) (pbweb.BackendClient, error) {
	u2, err := url.Parse(document.DocumentURI())
	if err != nil {
		log.Error("unexpected error parsing URI", err)
		return nil, err
	}

	addr := "localhost:10001"
	if u2.Query().Get("addr") != "" {
		addr = u2.Query().Get("addr")

		if !strings.Contains(addr, "http") {
			addr = "https://" + addr
		}
	}

	log.Printf("Connecting to %s\n", addr)

	clas := pbweb.NewBackendClient(serverAddr)
	document.Body().SetInnerHTML(`<div><h2>GopherJS gRPC-Web is great!</h2></div>`)

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
