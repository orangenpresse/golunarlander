package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

type Gametimer struct {
	oldTime int64
	delta   int64
}

func (g *Gametimer) update() {
	time := time.Time{}
	currentTime := time.UnixNano()
	g.delta = currentTime - g.oldTime
	g.oldTime = currentTime
}

type Game struct {
	g      float64
	height float64
	timer  Gametimer
}

type Vector2d struct {
	x float64
	y float64
}

type Lander struct {
	mass     float64
	velocity Vector2d
	position Vector2d
	force    float64
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	for {

		if termbox.PollEvent().Key == termbox.KeyEsc {
			os.Exit(0)
		}
	}
	// for {
	// 	// timer update
	// 	// apply physics
	// 	// render
	// 	// check exit
	// }
}
