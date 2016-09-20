package factory

import (
	"github.com/fuserobotics/gogame"

	"github.com/fuserobotics/goterram/components"
	"github.com/fuserobotics/goterram/game"
)

/* Build a terram game instance. */
func BuildTerramGame() (*gogame.Game, error) {
	componentTable := components.BuildTerramComponentTable()
	gameRules := &game.TerramGameRules{}

	return gogame.BuildGame(componentTable, gameRules)
}
