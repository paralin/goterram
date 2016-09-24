package components

import (
	"github.com/fuserobotics/gogame"
	"github.com/fuserobotics/gogame/components"
	"github.com/fuserobotics/goterram/components/sprite"
)

func BuildTerramComponentTable() gogame.ComponentTable {
	res := components.NewComponentTable()

	// Register components here
	res[sprite.SpriteComponentMeta.Id] = sprite.SpriteComponentFactory

	return res
}

var TerramComponentTable gogame.ComponentTable = BuildTerramComponentTable()
