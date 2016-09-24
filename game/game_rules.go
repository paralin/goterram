package game

import (
	"github.com/fuserobotics/gogame"
	"github.com/fuserobotics/goterram/entities"
)

// Main game tick rate in Hz
var TerramSettings gogame.GameSettings = gogame.GameSettings{
	Tick: 30,
}

type TerramGameRules struct {
	Game *gogame.Game

	// For use when detecting changes
	opMode gogame.GameOperatingMode

	// Testing
	entIdCounter  uint32
	hasSpawnedEnt bool
}

func (gr *TerramGameRules) Init(game *gogame.Game) {
	gr.Game = game
	gr.opMode = gogame.GameOperatingMode_LOCAL
}

func (gr *TerramGameRules) Update() {
	if !gr.hasSpawnedEnt {
		gr.hasSpawnedEnt = true
		gr.Game.SpawnEntity(entities.TerramEntityFactories.Ball, nil)
	}
}

func (gr *TerramGameRules) NextEntityId() uint32 {
	gr.entIdCounter++
	return gr.entIdCounter
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
