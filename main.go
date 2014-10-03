package main

import (
	"github.com/orangenpresse/golunarlander/graphic"
)

func main() {
	game := graphic.LanderGraphic{Width: 800, Height: 600}
	game.Start()
}
