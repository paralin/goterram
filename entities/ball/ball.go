package ball

import (
	"github.com/paralin/gogame"
	"github.com/paralin/gogame/components/transform"

	"github.com/paralin/goterram/components/physics"
	"github.com/paralin/goterram/components/sprite"
)

type BallFactory struct {
}

func (*BallFactory) Spawn(id uint32, game *gogame.Game) *gogame.Entity {
	ent := gogame.NewEntity(id, game)
	// Add a transform
	ent.AddComponent(transform.TransformComponentFactory.New())
	// Add a sprite
	ent.AddComponent(sprite.SpriteComponentFactory.New())
	// Add physics
	ent.AddComponent(physics.PhysicsComponentFactory.New())
	return ent
}

// This will produce an error if it doesn't implement an entity properly
var factoryTypeAssertion gogame.EntityFactory = &BallFactory{}
