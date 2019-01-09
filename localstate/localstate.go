package localstate

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"sync"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"

	"github.com/kyeett/gameserver"
	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/types"
)

// Ensure struct implements interface
var _ gameserver.GameState = (*LocalState)(nil)

type LocalState struct {
	world    types.World
	entities []entity.Entity
	mu       *sync.RWMutex
}

func New(world types.World) gameserver.GameState {
	entities := []entity.Entity{
		entity.Entity{ID: newID(), Type: entity.Coin, Position: types.Position{types.Coord{2, 1}, 0}, Owner: ""},
		entity.Entity{ID: newID(), Type: entity.Coin, Position: types.Position{types.Coord{8, 2}, 0}, Owner: ""},
	}

	return &LocalState{
		world:    world,
		entities: entities,
		mu:       &sync.RWMutex{},
	}
}

func (s *LocalState) NewPlayer() (entity.Entity, error) {

	ID := newID()

	s.mu.Lock()
	defer s.mu.Unlock()

	var pos types.Position
	var i, maxAttempts = 0, 1000
	for ; i <= maxAttempts; i++ {
		pos = types.Position{types.Coord{rand.Intn(s.world.Width), rand.Intn(s.world.Height)}, 0}
		if s.world.ValidTarget(pos) {
			break
		}

		if i == maxAttempts {
			return entity.Entity{}, errors.Errorf("Failed to find a valid start position after %d attempts\n", i)
		}
	}

	e := entity.Entity{
		ID:       ID,
		Type:     entity.Character,
		Position: pos,
		Owner:    "",
	}

	s.entities = append(s.entities, e)
	log.Infof("New player with ID: %s joined", e.ID)
	return e, nil
}

func (s *LocalState) Entities() []entity.Entity {
	s.mu.RLock()
	tmp := make([]entity.Entity, len(s.entities))
	copy(tmp, s.entities)
	s.mu.RUnlock()
	return tmp
}

func (s *LocalState) World() types.World {
	return s.world
}

func (s *LocalState) checkCollisions(p entity.Entity) {
	log.Debug("check for collisions")
	// Check for collisions
	s.mu.RLock()
	for i, e := range s.entities {
		if p != e && p.Position.Coord == e.Position.Coord {
			log.Info("Object ", p, "destroys", e)
			s.entities[i] = e.Destroy(p.ID)
		}
	}
	s.mu.RUnlock()
}

func (s *LocalState) PerformAction(e entity.Entity, p types.Position) (entity.Entity, error) {

	log.Info("Perform action", e, p)
	e, err := s.moveTo(e, p)
	if err != nil {
		return e, err
	}

	s.checkCollisions(e)
	return e, nil
}

func (s *LocalState) bridgable(a entity.Entity, p types.Position) bool {
	Up := types.Coord{0, -1}
	Left := types.Coord{-1, 0}
	Down := types.Coord{0, 1}
	Right := types.Coord{1, 0}

	var next, prev types.Coord
	switch p.Theta {
	case 0:
		next = p.Coord.Add(Up)
		prev = p.Coord.Add(Down)
	case 3:
		next = p.Coord.Add(Left)
		prev = p.Coord.Add(Right)
	case 2:
		next = p.Coord.Add(Down)
		prev = p.Coord.Add(Up)
	case 1:
		next = p.Coord.Add(Right)
		prev = p.Coord.Add(Left)
	}

	if s.world.At(p.Coord) == types.Water && s.world.At(next) != types.Water && s.world.At(prev) != types.Water {
		return true
	}
	return false
}

// Todo: design the rules for entity interaction a bit better
func (s *LocalState) moveTo(a entity.Entity, c types.Position) (entity.Entity, error) {

	bridgable := s.bridgable(a, c)

	if s.world.ValidTarget(c) == false && !bridgable {
		return entity.Entity{}, errors.New("invalid move")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	var found = -1
	var blocked bool
	for i, e := range s.entities {
		if e.ID == a.ID {
			found = i
			continue
		}

		if e.Type == entity.Character && e.Position.Coord == c.Coord {
			blocked = true
		}
	}

	if found == -1 {
		log.Fatalf("Should not happend, moveTo got invalid entityID=%s, from entity=%s", a.ID, a)
	}

	if !blocked {
		s.entities[found].Position = c
	}

	return s.entities[found], nil
}

func newID() string {
	hash := md5.New()
	hash.Write([]byte(strconv.Itoa(rand.Intn(123456))))
	ID := hex.EncodeToString(hash.Sum(nil))[0:12]
	return ID
}
