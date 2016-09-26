package components

import (
	"github.com/fuserobotics/gogame"
	"github.com/fuserobotics/gogame/components"

	"github.com/fuserobotics/goterram/components/physics"
	"github.com/fuserobotics/goterram/components/sprite"
)

func BuildTerramComponentTable() gogame.ComponentTable {
	res := components.NewComponentTable()

	// Register components here
	res[sprite.SpriteComponentMeta.Id] = sprite.SpriteComponentFactory
	res[physics.PhysicsComponentMeta.Id] = physics.PhysicsComponentFactory

	return res
}

var TerramComponentTable gogame.ComponentTable = BuildTerramComponentTable()
