package main

import (
	"github.com/orangenpresse/golunarlander/game"
)

func main() {
	game := game.LunarLander{Width: 800, Height: 600}
	game.Start()
}
