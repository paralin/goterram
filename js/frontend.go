package main

import (
	"github.com/fuserobotics/gogame"
	"github.com/gopherjs/gopherjs/js"
)

// Wraps a IFrontend in TypeScript
type JsFrontend struct {
	*js.Object
}

func (fe *JsFrontend) Init() {
	if fe.Object.Get("init") == nil {
		return
	}
	fe.Object.Call("init")
}

func (fe *JsFrontend) AddEntity(entity *gogame.Entity) gogame.FrontendEntity {
	if fe.Object.Get("addEntity") == nil {
		return nil
	}
	res := fe.Object.Call("addEntity", entity.ToNetworkInit())
	if res == nil {
		return nil
	}
	return &JsFrontendEntity{Object: res}
}

func (fe *JsFrontend) Destroy() {
	if fe.Object.Get("destroy") == nil {
		return
	}
	fe.Object.Call("destroy")
}