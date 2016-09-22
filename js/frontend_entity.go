package main

import (
	"github.com/fuserobotics/gogame"
	"github.com/gopherjs/gopherjs/js"
)

type JsFrontendEntity struct {
	*js.Object
}

func (fe *JsFrontendEntity) Init() {
	if fe.Object.Get("init") == nil {
		return
	}
	fe.Object.Call("init")
}

func (fe *JsFrontendEntity) AddComponent(id uint32) gogame.FrontendComponent {
	if fe.Object.Get("addComponent") == nil {
		return nil
	}
	res := fe.Object.Call("addComponent", id)
	if res == nil {
		return nil
	}
	return &JsFrontendComponent{Object: res}
}

func (fe *JsFrontendEntity) InitLate() {
	if fe.Object.Get("initLate") == nil {
		return
	}
	fe.Object.Call("initLate")
}

func (fe *JsFrontendEntity) Destroy() {
	if fe.Object.Get("destroy") == nil {
		return
	}
	fe.Object.Call("destroy")
}
