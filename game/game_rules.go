package game

import "github.com/fuserobotics/gogame"

type TerramGameRules struct {
	Game *gogame.Game
}

func (gr *TerramGameRules) Init(game *gogame.Game) {
	gr.Game = game
}
