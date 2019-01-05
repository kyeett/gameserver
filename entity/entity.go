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
