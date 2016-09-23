package game

import "github.com/fuserobotics/gogame"

type TerramGameRules struct {
	Game *gogame.Game

	// For use when detecting changes
	opMode gogame.GameOperatingMode
}

func (gr *TerramGameRules) Init(game *gogame.Game) {
	gr.Game = game
	gr.opMode = gogame.GameOperatingMode_LOCAL
}

func (gr *TerramGameRules) SetGameOperatingMode(opMode gogame.GameOperatingMode) {
	if opMode == gr.opMode {
		return
	}

	if gr.opMode == gogame.GameOperatingMode_REMOTE {
		// Delete all entities in prep for network sync
	}

	gr.opMode = opMode
}

func (gr *TerramGameRules) Destroy() {
}
