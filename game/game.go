package game

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
	"github.com/orangenpresse/golunarlander/simulation"
	"runtime"
)

type LunarLanderGame struct {
	run        bool
	thrust     simulation.ThrusterState
	Width      int
	Height     int
	timer      Timer
	window     *glfw.Window
	Simulation simulation.Simulation
	Graphic    *Graphic
	Options    simulation.Options
}

func (lg *LunarLanderGame) CreateWindow() (shaderVersion string) {
	runtime.LockOSThread()

	glfw.SetErrorCallback(lg.handleErrors)

	glfw.Init()

	version := glfw.GetVersionString()
	if version == "3.0.4 Cocoa NSGL chdir menubar" {
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 3)
		glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
		glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
		shaderVersion = "330"
	} else {
		shaderVersion = "130"
	}

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
	return shaderVersion
}

func (lg *LunarLanderGame) handleErrors(err glfw.ErrorCode, msg string) {
	fmt.Printf("GLFW ERROR: %v: %v\n", err, msg)
}

func (lg *LunarLanderGame) Start() {
	shaderVersion := lg.CreateWindow()
	lg.run = true
	lg.window.SetKeyCallback(lg.handleEvents)
	lg.timer.Start()
	lg.Options = simulation.Options{false}
	lg.Simulation = simulation.Simulation{}
	lg.Simulation.Start(&lg.Options)
	lg.initGraphics(shaderVersion)
	lg.mainLoop()
	lg.end()
}

func (lg *LunarLanderGame) end() {
	lg.Graphic.end()
	lg.window.SetShouldClose(true)
	lg.window.Destroy()
	glfw.Terminate()
}

func (lg *LunarLanderGame) mainLoop() {
	for lg.run == true {
		lg.timer.Update()
		glfw.PollEvents()
		lg.Simulation.Update(lg.timer.GetDelta(), lg.thrust)
		lg.Graphic.render()
		lg.window.SwapBuffers()
	}
}
