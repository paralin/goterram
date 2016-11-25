package entities

import (
	"github.com/paralin/gogame"
	"github.com/paralin/goterram/entities/ball"
)

var TerramEntityFactories = struct {
	Ball gogame.EntityFactory
}{
	Ball: &ball.BallFactory{},
}
