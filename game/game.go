package game

import (
	_ "fmt"
	"github.com/orangenpresse/golunarlander/simulation"
	"github.com/veandco/go-sdl2/sdl"
)

type LunarLander struct {
	run        bool
	thrust     bool
	Width      int64
	Height     int64
	surface    *sdl.Surface
	window     *sdl.Window
	timer      Timer
	Simulation simulation.Simulation
}

func (lg *LunarLander) Start() {
	lg.run = true
	lg.timer.Start()
	lg.Simulation = simulation.Simulation{}
	lg.Simulation.Start()
	lg.window = sdl.CreateWindow("Lunar Lander", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int(lg.Width), int(lg.Height), sdl.WINDOW_SHOWN)
	lg.surface = lg.window.GetSurface()
	lg.mainLoop()
	lg.end()
}

func (lg *LunarLander) end() {
	lg.window.Destroy()
}

func (lg *LunarLander) mainLoop() {
	for lg.run == true {
		lg.timer.Update()
		lg.handleEvents()
		lg.Simulation.Update(lg.timer.GetDelta(), lg.thrust)
		lg.render()
		lg.window.UpdateSurface()
	}
}
