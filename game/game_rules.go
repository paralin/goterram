package game

import (
	"fmt"

	"github.com/paralin/gogame"
	"github.com/paralin/goterram/common"
	"github.com/paralin/goterram/entities"

	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

// Main game tick rate in Hz
var TerramSettings gogame.GameSettings = gogame.GameSettings{
	Tick: 30,
}

var physicsStep float32 = float32(1.0) / float32(TerramSettings.Tick)

type TerramGameRules struct {
	Game  *gogame.Game
	Space *chipmunk.Space

	// For use when detecting changes
	opMode gogame.GameOperatingMode

	// Testing
	entIdCounter  uint32
	hasSpawnedEnt bool
}

func (gr *TerramGameRules) Init(game *gogame.Game) {
	fmt.Println("Initializing physics engine...")
	gr.Game = game
	gr.opMode = gogame.GameOperatingMode_LOCAL

	gr.Space = chipmunk.NewSpace()
	gr.Space.Gravity = vect.Vect{X: 0, Y: -9}

	game.GameCommon = &common.TerramCommon{
		Space: gr.Space,
	}
}

func (gr *TerramGameRules) Update() {
	if !gr.hasSpawnedEnt {
		gr.hasSpawnedEnt = true
		gr.Game.SpawnEntity(entities.TerramEntityFactories.Ball, nil)
	}

	// Step physics
	// Determine dt (in fraction of second, i.e. 1/60)
	gr.StepPhysics(physicsStep)
}

func (gr *TerramGameRules) StepPhysics(dt float32) {
	gr.Space.Step(vect.Float(dt))
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
