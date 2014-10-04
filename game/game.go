package game

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/orangenpresse/golunarlander/simulation"
)

type LunarLander struct {
	run        bool
	thrust     bool
	Width      int64
	Height     int64
	timer      Timer
	window     *glfw.Window
	Simulation simulation.Simulation
}

func (lg *LunarLander) Start(window *glfw.Window) {
	lg.run = true
	lg.window = window
	lg.window.SetKeyCallback(lg.handleEvents)
	lg.timer.Start()
	lg.Simulation = simulation.Simulation{}
	lg.Simulation.Start()
	lg.mainLoop()
	lg.end()
}

func (lg *LunarLander) end() {
	lg.window.SetShouldClose(true)
}

func (lg *LunarLander) mainLoop() {
	glfw.Init()
	for lg.run == true {
		lg.timer.Update()
		glfw.PollEvents()
		lg.Simulation.Update(lg.timer.GetDelta(), lg.thrust)
		lg.render()
	}
}
