package localserver

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
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
		entity.Entity{ID: newID(), Type: entity.Coin, Position: types.Position{types.Coord{2, 1}, 0}, Owner: ""},
		entity.Entity{ID: newID(), Type: entity.Coin, Position: types.Position{types.Coord{8, 2}, 0}, Owner: ""},
	}

	return &LocalServer{
		world:    types.FirstWorld,
		entities: entities,
	}
}

func (s *LocalServer) NewPlayer() (entity.Entity, error) {

	ID := newID()
	e := entity.Entity{
		ID:       ID,
		Type:     entity.Character,
		Position: types.Position{types.Coord{rand.Intn(3), rand.Intn(3)}, 0},
		Owner:    "",
	}

	s.entities = append(s.entities, e)
	fmt.Println(s.entities)
	return e, nil
}

func (s *LocalServer) Entities() []entity.Entity {
	return s.entities
}

func (s *LocalServer) World() types.World {
	return s.world
}

func (s *LocalServer) checkCollisions(p entity.Entity) {

	// Check for collisions
	for i, e := range s.entities {
		if p != e && p.Position.Coord == e.Position.Coord {
			fmt.Println(p, "destroy", e)
			s.entities[i] = e.Destroy(p.ID)
		}
	}
}

func (s *LocalServer) PerformAction(e entity.Entity, p types.Position) (entity.Entity, error) {
	fmt.Println("Perform action", e, p)
	e = s.moveTo(e, p)
	fmt.Println("Performed action", e)
	s.checkCollisions(e)
	return e, nil
}

// Todo: design the rules for entity interaction a bit better
func (s *LocalServer) moveTo(a entity.Entity, c types.Position) entity.Entity {
	if s.world.ValidTarget(c) == false {
		return a
	}

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

	return s.entities[found]
}

func newID() string {
	hash := md5.New()
	hash.Write([]byte(strconv.Itoa(rand.Intn(123456))))
	ID := hex.EncodeToString(hash.Sum(nil))[0:12]
	return ID
}
