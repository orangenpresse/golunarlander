package game

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
	"github.com/orangenpresse/golunarlander/simulation"
	"runtime"
)

type LunarLander struct {
	run        bool
	thrust     bool
	Width      int
	Height     int
	timer      Timer
	window     *glfw.Window
	Simulation simulation.Simulation
	Graphic    *Graphic
}

func (lg *LunarLander) CreateWindow() {
	runtime.LockOSThread()

	glfw.SetErrorCallback(lg.handleErrors)

	glfw.Init()
	//glfw.WindowHint(glfw.ContextVersionMajor, 3)
	//glfw.WindowHint(glfw.ContextVersionMinor, 0)
	//glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	//glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	window, err := glfw.CreateWindow(lg.Width, lg.Height, "Lunar Lander", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetCloseCallback(func(window *glfw.Window) {
		lg.run = false
	})

	// use vsync
	glfw.SwapInterval(1)

	lg.window = window
}

func (lg *LunarLander) handleErrors(err glfw.ErrorCode, msg string) {
	fmt.Printf("GLFW ERROR: %v: %v\n", err, msg)
}

func (lg *LunarLander) Start() {
	lg.CreateWindow()
	lg.run = true
	lg.window.SetKeyCallback(lg.handleEvents)
	lg.timer.Start()
	lg.Simulation = simulation.Simulation{}
	lg.Simulation.Start()
	lg.initGraphics()
	lg.mainLoop()
	lg.end()
}

func (lg *LunarLander) end() {
	lg.Graphic.end()
	lg.window.SetShouldClose(true)
	lg.window.Destroy()
	glfw.Terminate()
}

func (lg *LunarLander) mainLoop() {
	for lg.run == true {
		lg.timer.Update()
		glfw.PollEvents()
		lg.Simulation.Update(lg.timer.GetDelta(), lg.thrust)
		lg.Graphic.render()
		lg.window.SwapBuffers()
	}
}
