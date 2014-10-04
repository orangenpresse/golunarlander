package game

import (
	glfw "github.com/go-gl/glfw3"
)

func (lg *LunarLander) handleEvents(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		lg.run = false
	}

	if key == glfw.KeyUp && action == glfw.Press {
		lg.thrust = true
	} else if key == glfw.KeyUp && action == glfw.Release {
		lg.thrust = false
	}

	if key == glfw.KeyR {
		lg.Simulation.Start()
		lg.Graphic.Lander = lg.Simulation.GetLander()
	}
}
