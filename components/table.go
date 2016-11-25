package components

import (
	"github.com/paralin/gogame"
	"github.com/paralin/gogame/components"

	"github.com/paralin/goterram/components/physics"
	"github.com/paralin/goterram/components/sprite"
)

func BuildTerramComponentTable() gogame.ComponentTable {
	res := components.NewComponentTable()

	// Register components here
	res[sprite.SpriteComponentMeta.Id] = sprite.SpriteComponentFactory
	res[physics.PhysicsComponentMeta.Id] = physics.PhysicsComponentFactory

	return res
}

var TerramComponentTable gogame.ComponentTable = BuildTerramComponentTable()
