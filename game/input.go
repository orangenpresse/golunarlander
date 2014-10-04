package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

const (
	KEY_UP    = 82
	KEY_DOWN  = 81
	KEY_LEFT  = 80
	KEY_RIGHT = 79
	KEY_R     = 21
	KEY_ESC   = 41
)

func (lg *LunarLander) handleEvents() {
	event := sdl.PollEvent()
	if event != nil {
		eventType := reflect.TypeOf(event).String()

		switch eventType {

		case "*sdl.QuitEvent":
			lg.run = false

		case "*sdl.KeyDownEvent":
			ev, _ := event.(*sdl.KeyDownEvent)

			//fmt.Println(ev.Keysym.Scancode)

			if ev.Keysym.Scancode == KEY_ESC {
				lg.run = false
			}
			if ev.Keysym.Scancode == KEY_UP {
				lg.thrust = true
			}
			if ev.Keysym.Scancode == KEY_R {
				lg.Simulation.Start()
			}

		case "*sdl.KeyUpEvent":
			if ev, _ := event.(*sdl.KeyUpEvent); ev.Keysym.Scancode == KEY_UP {
				lg.thrust = false
			}
		}

	}
}
