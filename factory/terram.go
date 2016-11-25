package factory

import (
	"github.com/paralin/gogame"

	"github.com/paralin/goterram/components"
	"github.com/paralin/goterram/game"
)

/* Build a terram game instance. */
func BuildTerramGame(frontend gogame.Frontend) (*gogame.Game, error) {
	componentTable := components.BuildTerramComponentTable()
	gameRules := &game.TerramGameRules{}

	return gogame.BuildGame(game.TerramSettings, componentTable, gameRules, frontend)
}
