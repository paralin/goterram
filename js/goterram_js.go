package main

import (
	"fmt"

	"github.com/fuserobotics/gogame"
	"github.com/fuserobotics/goterram/factory"
	"github.com/gopherjs/gopherjs/js"
)

type TerramJsGlobal struct {
}

func (*TerramJsGlobal) BuildTerramGame(frontend *JsFrontend) *js.Object {
	game, err := factory.BuildTerramGame(frontend)

	if err != nil {
		fmt.Printf("Error building game: %s\n", err.Error())
		return nil
	}

	// return game
	return js.MakeWrapper(game)
}

func main() {
	fmt.Println("Initing goterram module...")
	exports := js.Module.Get("exports")
	exports.Set("TerramBuilder", js.MakeWrapper(&TerramJsGlobal{}))
	exports.Set("EGameOperatingMode", gogame.GameOperatingMode_value)
	exports.Set("EGameOperatingModeString", gogame.GameOperatingMode_name)
}
