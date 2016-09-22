package main

import (
	"github.com/gopherjs/gopherjs/js"
)

type JsFrontendComponent struct {
	*js.Object
}

func (fc *JsFrontendComponent) Init() {
	if fc.Object.Get("init") == nil {
		return
	}
	fc.Object.Call("init")
}

func (fc *JsFrontendComponent) Destroy() {
	if fc.Object.Get("destroy") == nil {
		return
	}
	fc.Object.Call("destroy")
}
