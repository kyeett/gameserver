package grpc

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/kyeett/gameserver"
	"github.com/kyeett/gameserver/entity"

	"github.com/kyeett/gameserver/types"
)

// Ensure struct implements interface
var _ gameserver.GameServer = (*GrpcClient)(nil)

func NewClient(serverAddr string, secure bool) (gameserver.GameServer, error) {
	log.SetLevel(log.DebugLevel)
	ctx := context.Background() // Todo move this out
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)

	client, err := CorrectClient(serverAddr, secure)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client")
	}

	resp, err := client.WorldRequest(ctx, getEmpty()) // grpc.WaitForReady(true)
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
	stream, err := s.client.EntityStream(context.Background(), getEmpty())
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

func (s *GrpcClient) NewPlayer() (entity.Entity, error) {
	t := time.Now()
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	log.Debug("created a new player in", time.Since(t))
	resp, err := s.client.NewPlayer(ctx, getEmpty())
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
