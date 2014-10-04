package main

import (
	"fmt"
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/orangenpresse/golunarlander/game"
	"runtime"
)

func main() {
	game := game.LunarLander{Width: 800, Height: 600}
	game.Start(createWindow())

}

func handleErrors(err glfw.ErrorCode, msg string) {
	fmt.Printf("GLFW ERROR: %v: %v\n", err, msg)
}

func createWindow() *glfw.Window {
	runtime.LockOSThread()

	glfw.SetErrorCallback(handleErrors)

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

	return window
}
