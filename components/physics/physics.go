package physics

import (
	"fmt"

	"github.com/fuserobotics/gogame"
	"github.com/golang/protobuf/proto"

	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"

	"github.com/fuserobotics/gogame/components/transform"
	"github.com/fuserobotics/goterram/common"
)

// terram components start at 1xx
var PhysicsComponentMeta gogame.ComponentMeta = gogame.ComponentMeta{
	Id: 101,
}

type PhysicsComponent struct {
	Entity     *gogame.Entity
	Data       PhysicsData
	Frontend   gogame.FrontendComponent
	Body       *chipmunk.Body
	Transform  *transform.TransformComponent
	GameCommon *common.TerramCommon
}

// Initialize a brand new sprite component
func (tc *PhysicsComponent) Init(ent *gogame.Entity) {
	tc.Entity = ent
	fmt.Println("Post entity init")
}

// Initialize a remotely created sprite component, over the network
func (tc *PhysicsComponent) InitWithData(ent *gogame.Entity, data []byte) {
	tc.Entity = ent

	// parse data, handle error here somehow
	proto.Unmarshal(data, &tc.Data)
}

func (tc *PhysicsComponent) InitLate() {
	fmt.Println("Pre gamecommon check")
	tc.GameCommon = tc.Entity.Game.GameCommon.(*common.TerramCommon)
	fmt.Println("Pre transform check")
	tc.Transform = tc.Entity.GetComponent(transform.TransformComponentMeta.Id).(*transform.TransformComponent)

	// this might be null
	pos := tc.Transform.Data.Position
	if pos == nil {
		fmt.Println("Nil position in initlate in physics")
		return
	}

	// Initialize the physics from the Data. For now, hardcode it.
	mass := 1
	radius := 25
	ball := chipmunk.NewCircle(vect.Vector_Zero, float32(radius))
	ball.SetElasticity(0.95) // no idea what this does

	tc.Body = chipmunk.NewBody(vect.Float(mass), ball.Moment(float32(mass)))
	tc.Body.SetPosition(vect.Vect{X: vect.Float(pos.X), Y: vect.Float(pos.Y)})
	tc.Body.AddShape(ball)

	tc.GameCommon.Space.AddBody(tc.Body)
}

func (tc *PhysicsComponent) ShouldUpdate() bool {
	return true
}

func (tc *PhysicsComponent) Update() {
	pos := tc.Body.Position()
	tc.Transform.Data.Position.X = float32(pos.X)
	tc.Transform.Data.Position.Y = float32(pos.Y)
	tc.Transform.SyncPosition()
}

func (tc *PhysicsComponent) Meta() gogame.ComponentMeta {
	return PhysicsComponentMeta
}

func (tc *PhysicsComponent) InitData() []byte {
	data, _ := proto.Marshal(&tc.Data)
	return data
}

func (tc *PhysicsComponent) InitFrontend(fe gogame.FrontendComponent) {
	tc.Frontend = fe
}

func (tc *PhysicsComponent) Destroy() {
	tc.GameCommon.Space.RemoveBody(tc.Body)
	tc.Body = nil
	tc.GameCommon = nil
}

// Factory to spawn transform components
type physicsComponentFactory struct {
}

func (tff *physicsComponentFactory) Meta() gogame.ComponentMeta {
	return PhysicsComponentMeta
}

func (tff *physicsComponentFactory) New() gogame.Component {
	return &PhysicsComponent{}
}

// Assert at compile time the component is valid
// This line will fail otherwise.
var PhysicsComponentFactory gogame.ComponentFactory = &physicsComponentFactory{}
