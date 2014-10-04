package game

import (
	glfw "github.com/go-gl/glfw3"
)

const (
	KEY_UP    = 82
	KEY_DOWN  = 81
	KEY_LEFT  = 80
	KEY_RIGHT = 79
	KEY_R     = 21
	KEY_ESC   = 41
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
