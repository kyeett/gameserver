package entity

import (
	"fmt"

	"github.com/kyeett/gameserver/types"
)

type Type int

const (
	Coin Type = iota + 1
	Score
	Character
	Bridge
	Empty
)

func (t Type) String() string {
	switch t {
	case Coin:
		return "Coin"
	case Score:
		return "Score"
	case Character:
		return "Character"
	case Bridge:
		return "Bridge"
	case Empty:
		return "Empty"
	}
	return "Unknown"
}

// Entity is a generic object that has a ID, type and a generic Position in space
type Entity struct {
	ID       string
	Type     Type
	Position types.Position
	Owner    string
}

func (e Entity) Destroy(destoryerID string) Entity {
	switch e.Type {
	case Coin:
		return Entity{
			Position: types.Position{types.Coord{-1, -1}, 0},
			Type:     Score,
			Owner:    destoryerID,
		}
	}

	return Entity{
		Position: types.Position{types.Coord{-1, -1}, 0},
		Type:     Empty,
	}
}

func (e Entity) String() string {
	return fmt.Sprintf("% 10s:%s at (% 5s)", e.Type, e.ID, e.Position.Coord)
}
