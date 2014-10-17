package game

import (
	glfw "github.com/go-gl/glfw3"
)

func (lg *LunarLanderGame) handleEvents(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		lg.run = false
	}

	// Down
	if key == glfw.KeyUp && action == glfw.Press {
		lg.thrust.Bottom = true
	} else if key == glfw.KeyUp && action == glfw.Release {
		lg.thrust.Bottom = false
	}

	// Left
	if key == glfw.KeyLeft && action == glfw.Press {
		lg.thrust.Right = true
	} else if key == glfw.KeyLeft && action == glfw.Release {
		lg.thrust.Right = false
	}

	// Right
	if key == glfw.KeyRight && action == glfw.Press {
		lg.thrust.Left = true
	} else if key == glfw.KeyRight && action == glfw.Release {
		lg.thrust.Left = false
	}

	// R
	if key == glfw.KeyR {
		lg.Simulation.Start()
		lg.Graphic.Lander = lg.Simulation.GetLander()
	}
}
