package main

import (
	"github.com/orangenpresse/golunarlander/game"
)

func main() {
	game := game.LunarLander{Width: 1024, Height: 720}
	game.Start()
}
