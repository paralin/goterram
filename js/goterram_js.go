package main

import (
	"fmt"
	"github.com/fuserobotics/goterram/factory"
	"github.com/gopherjs/gopherjs/js"
)

type TerramJsGlobal struct {
}

func (*TerramJsGlobal) BuildTerramGame() *js.Object {
	game, err := factory.BuildTerramGame()
	if err != nil {
		fmt.Printf("Error building game: %s\n", err.Error())
		return nil
	}
	// return game
	return js.MakeWrapper(game)
}

func main() {
	js.Module.Set("exports", js.MakeWrapper(&TerramJsGlobal{}))
}
