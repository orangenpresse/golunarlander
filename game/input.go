package game

import (
	glfw "github.com/go-gl/glfw3"
)

func (lg *LunarLander) handleEvents(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		lg.run = false
	}
	if key == glfw.KeyUp {
		lg.thrust = true
	}
	if key == glfw.KeyR {
		lg.Simulation.Start()
	}
}
