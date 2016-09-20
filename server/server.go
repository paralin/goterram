package main

import (
	"fmt"
	"github.com/fuserobotics/goterram/factory"
)

func main() {
	_, err := factory.BuildTerramGame()
	if err != nil {
		fmt.Printf("Error making game: %s\n", err)
		return
	}
	fmt.Printf("Built game successfully.\n")
}
