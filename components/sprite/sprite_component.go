package sprite

import (
	"github.com/fuserobotics/gogame"
	"github.com/golang/protobuf/proto"
)

// terram components start at 1xx
var SpriteComponentMeta gogame.ComponentMeta = gogame.ComponentMeta{
	Id: 100,
}

type SpriteComponent struct {
	Entity   *gogame.Entity
	Data     SpriteData
	Frontend gogame.FrontendComponent
}

// Initialize a brand new transform component
func (tc *SpriteComponent) Init(ent *gogame.Entity) {
	tc.Entity = ent
}

// Initialize a remotely created transform component, over the network
func (tc *SpriteComponent) InitWithData(ent *gogame.Entity, data []byte) {
	tc.Entity = ent

	// parse data, handle error here somehow
	proto.Unmarshal(data, &tc.Data)
}

func (tc *SpriteComponent) InitLate() {
}

func (tc *SpriteComponent) ShouldUpdate() bool {
	return false
}

func (tc *SpriteComponent) Update() {
}

func (tc *SpriteComponent) Meta() gogame.ComponentMeta {
	return SpriteComponentMeta
}

func (tc *SpriteComponent) InitData() []byte {
	data, _ := proto.Marshal(&tc.Data)
	return data
}

func (tc *SpriteComponent) InitFrontend(fe gogame.FrontendComponent) {
	tc.Frontend = fe
}

func (tc *SpriteComponent) Destroy() {
}

// Factory to spawn transform components
type spriteComponentFactory struct {
}

func (tff *spriteComponentFactory) Meta() gogame.ComponentMeta {
	return SpriteComponentMeta
}

func (tff *spriteComponentFactory) New() gogame.Component {
	return &SpriteComponent{}
}

// Assert at compile time the component is valid
// This line will fail otherwise.
var SpriteComponentFactory gogame.ComponentFactory = &spriteComponentFactory{}
