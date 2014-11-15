package game

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
	data "github.com/orangenpresse/golunarlander/dataObjects"
	"github.com/orangenpresse/golunarlander/game/graphic"
	"github.com/orangenpresse/golunarlander/simulation"
	"runtime"
)

type LunarLanderGame struct {
	run        bool
	thrust     data.ThrusterState
	Width      int
	Height     int
	timer      Timer
	window     *glfw.Window
	Simulation simulation.Simulation
	Graphic    *graphic.Graphic
	Options    data.Options
}

func (lg *LunarLanderGame) CreateWindow() (shaderVersion string) {
	runtime.LockOSThread()

	glfw.SetErrorCallback(lg.handleErrors)

	glfw.Init()

	version := glfw.GetVersionString()
	fmt.Println(version)
	if version == "3.0.4 Cocoa NSGL chdir menubar" || version == "3.1.0 Cocoa NSGL chdir menubar retina dynamic" {
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

func (lg *LunarLanderGame) initGraphics(shaderVersion string) {
	lg.Options.BufferWidth, lg.Options.BufferHeight = lg.window.GetFramebufferSize()
	lg.Graphic = graphic.NewGraphic(&lg.Options, shaderVersion, lg.Simulation.GetLander())
}

func (lg *LunarLanderGame) handleErrors(err glfw.ErrorCode, msg string) {
	fmt.Printf("GLFW ERROR: %v: %v\n", err, msg)
}

func (lg *LunarLanderGame) Start() {
	shaderVersion := lg.CreateWindow()
	lg.run = true
	lg.window.SetKeyCallback(lg.handleEvents)
	lg.timer.Start()
	lg.Options = data.Options{}
	lg.Options.DebugMode = false
	lg.Simulation = simulation.Simulation{}
	lg.Simulation.Start(&lg.Options)
	lg.initGraphics(shaderVersion)
	lg.mainLoop()
	lg.end()
}

func (lg *LunarLanderGame) end() {
	lg.Graphic.End()
	lg.window.SetShouldClose(true)
	lg.window.Destroy()
	glfw.Terminate()
}

func (lg *LunarLanderGame) mainLoop() {
	for lg.run == true {
		lg.timer.Update()
		glfw.PollEvents()
		lg.Simulation.Update(lg.timer.GetDelta(), lg.thrust)
		lg.Graphic.Render()
		lg.window.SwapBuffers()
	}
}
