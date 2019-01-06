package gameserver

import (
	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/types"
)

// Todo: rename to GameState?
type GameServer interface {
	NewPlayer() (entity.Entity, error)
	PerformAction(entity.Entity, types.Position) (entity.Entity, error)
	World() types.World
	Entities() []entity.Entity
}
