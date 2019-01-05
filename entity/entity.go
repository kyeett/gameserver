package entity

import (
	"github.com/kyeett/gameserver/types"
)

type Type int

const (
	Coin Type = iota + 1
	Score
	Character
	Empty
)

// Entity is a generic object that has a ID, type and a generic Position in space
type Entity struct {
	ID       string
	Type     Type
	Position types.Position
}

func (e Entity) Destroy() Entity {
	switch e.Type {
	case Coin:
		return Entity{
			Position: types.Position{types.Coord{-1, -1}, 0},
			Type:     Score,
		}
	}

	return Entity{
		Position: types.Position{types.Coord{-1, -1}, 0},
		Type:     Empty,
	}
}
