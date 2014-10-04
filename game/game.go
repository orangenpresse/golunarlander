package game

import (
	"fmt"
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/orangenpresse/golunarlander/simulation"
	"runtime"
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

func (lg *LunarLander) CreateWindow() {
	runtime.LockOSThread()

	glfw.SetErrorCallback(lg.handleErrors)

	glfw.Init()
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	window, err := glfw.CreateWindow(800, 600, "Lunar Lander", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	// use vsync
	glfw.SwapInterval(1)

	// have to call this here or some OpenGL calls like CreateProgram will cause segfault
	if gl.Init() != 0 {
		panic("glew init failed")
	}
	gl.GetError() // ignore INVALID_ENUM that GLEW raises when using OpenGL 3.2+

	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()

	lg.window = window
	lg.start()
}

func (lg *LunarLander) handleErrors(err glfw.ErrorCode, msg string) {
	fmt.Printf("GLFW ERROR: %v: %v\n", err, msg)
}

func (lg *LunarLander) start() {
	lg.run = true
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
	for lg.run == true {
		lg.timer.Update()
		glfw.PollEvents()
		lg.Simulation.Update(lg.timer.GetDelta(), lg.thrust)
		lg.render()
	}
}
