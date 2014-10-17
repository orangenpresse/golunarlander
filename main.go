package main

import (
	"github.com/orangenpresse/golunarlander/game"
)

func main() {
	game := game.LunarLanderGame{Width: 800, Height: 600}
	game.Start()
}
