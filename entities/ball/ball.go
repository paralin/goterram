package ball

import (
	"github.com/fuserobotics/gogame"
	"github.com/fuserobotics/gogame/components/transform"
)

type BallFactory struct {
}

func (*BallFactory) Spawn(id uint32) *gogame.Entity {
	ent := gogame.NewEntity(id, nil)
	// Add a transform
	ent.AddComponent(transform.TransformComponentFactory.New())
	return ent
}

// This will produce an error if it doesn't implement an entity properly
var factoryTypeAssertion gogame.EntityFactory = &BallFactory{}
