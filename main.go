package main

import (
	"fmt"
	"github.com/orangenpresse/golunarlander/graphic"
)

func main() {
	game := graphic.LanderGrapic{Width: 800, Height: 600}
	game.Start()
}
