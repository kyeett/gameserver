package localserver

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"

	"github.com/kyeett/gameserver"
	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/types"
)

// Ensure struct implements interface
var _ gameserver.GameServer = (*LocalServer)(nil)

type LocalServer struct {
	world    types.World
	entities []entity.Entity
}

func New() gameserver.GameServer {
	entities := []entity.Entity{
		entity.Entity{NewID(), entity.Coin, types.Position{types.Coord{2, 1}, 0}},
		entity.Entity{NewID(), entity.Coin, types.Position{types.Coord{8, 2}, 0}},
	}

	return &LocalServer{
		world:    types.FirstWorld,
		entities: entities,
	}
}

func (s *LocalServer) NewPlayer() (entity.Entity, error) {

	ID := NewID()
	e := entity.Entity{
		ID,
		entity.Character,
		types.Position{types.Coord{2, 2}, 0},
	}

	s.entities = append(s.entities, e)
	return e, nil
}

func (s *LocalServer) PerformAction(e entity.Entity, p types.Position) (entity.Entity, error) {
	e = s.moveTo(e, p)
	s.checkCollisions()
	return e, nil
}

func (s *LocalServer) Entities() []entity.Entity {
	entities := []entity.Entity{}
	return entities
}

func (s *LocalServer) World() types.World {
	return s.world
}

func (s *LocalServer) checkCollisions() {

	// Check for collisions
	// for i, e := range s.entities {
	// 	if p != e && p.Coord == e.Coord {
	// 		s.entities[i] = e.Destroy()
	// 		s.score = calculateScore(s.entities)
	// 	}
	// }
}

func NewID() string {
	hash := md5.New()
	hash.Write([]byte(strconv.Itoa(rand.Intn(123456))))
	ID := hex.EncodeToString(hash.Sum(nil))[0:12]
	return ID
}

func (s *LocalServer) moveTo(a entity.Entity, c types.Position) entity.Entity {
	if s.world.ValidTarget(c) == false {
		return a
	}

	for i, e := range s.entities {
		if e == a {
			s.entities[i].Position = c
			return s.entities[i]
		}
	}
	return a
}
