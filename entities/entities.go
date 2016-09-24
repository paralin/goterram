package entities

import (
	"github.com/fuserobotics/gogame"
	"github.com/fuserobotics/goterram/entities/ball"
)

var TerramEntityFactories = struct {
	Ball gogame.EntityFactory
}{
	Ball: &ball.BallFactory{},
}
