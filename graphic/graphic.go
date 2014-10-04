package graphic

import (
	_ "fmt"
	"github.com/orangenpresse/golunarlander/simulation"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
	"time"
)

const (
	KEY_UP    = 82
	KEY_DOWN  = 81
	KEY_LEFT  = 80
	KEY_RIGHT = 79
	KEY_R     = 21
	KEY_ESC   = 41
)

type Timer struct {
	oldTime int64
	delta   int64
}

func (t *Timer) Start() {
	t.oldTime = time.Now().UnixNano()
}

func (t *Timer) Update() {
	current := time.Now().UnixNano()
	t.delta = current - t.oldTime
	t.oldTime = current
}

func (t *Timer) GetDelta() int64 {
	return t.delta
}

type SdlLander interface {
	GetPosition() simulation.Vector2D
}

type Simulation interface {
	Start()
	Update(int64, bool)
	GetLander() SdlLander
}

type LanderGraphic struct {
	run        bool
	thrust     bool
	Width      int64
	Height     int64
	surface    *sdl.Surface
	window     *sdl.Window
	timer      Timer
	Simulation simulation.Simulation
}

func (lg *LanderGraphic) Start() {
	lg.run = true
	lg.timer.Start()
	lg.Simulation = simulation.Simulation{}
	lg.Simulation.Start()
	lg.window = sdl.CreateWindow("Lunar Lander", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int(lg.Width), int(lg.Height), sdl.WINDOW_SHOWN)
	lg.surface = lg.window.GetSurface()
	lg.render()
	lg.end()
}

func (lg *LanderGraphic) end() {
	lg.window.Destroy()
}

func (lg *LanderGraphic) render() {
	for lg.run == true {
		lg.timer.Update()
		lg.handleEvents()
		lg.Simulation.Update(lg.timer.GetDelta(), lg.thrust)
		lg.clearSurface()
		lg.renderMoonSurface()
		lg.renderLander()
		lg.drawHud()
		lg.window.UpdateSurface()
	}
}

func (lg *LanderGraphic) handleEvents() {
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

func (lg *LanderGraphic) clearSurface() {
	rect := sdl.Rect{0, 0, int32(lg.Width), int32(lg.Height)}
	lg.surface.FillRect(&rect, 0x00000000)
}

func (lg *LanderGraphic) renderMoonSurface() {
	surfaceRect := sdl.Rect{0, 590, 800, 10}
	lg.surface.FillRect(&surfaceRect, 0x007a5345)
}

func (lg *LanderGraphic) drawHud() {
	posY := lg.Height - 125

	bg := sdl.Rect{int32(5), int32(posY), 20, 110}
	lg.surface.FillRect(&bg, 0x00878b88)

	fuelBorder := sdl.Rect{int32(8), int32(posY + 3), 14, 104}
	lg.surface.FillRect(&fuelBorder, 0x00c3c9c4)

	var fuel int32 = 50

	fuelBar := sdl.Rect{int32(10), (100 - fuel) + int32(posY+5), 10, fuel}
	lg.surface.FillRect(&fuelBar, 0x0000de3c)
}

func (lg *LanderGraphic) renderLander() {
	landerPos := lg.Simulation.GetLander().GetPosition()
	posX := int32(landerPos.X)
	posY := int32(lg.Height-25) - int32(landerPos.Y)

	landerRect := sdl.Rect{posX, posY, 10, 15}
	lg.surface.FillRect(&landerRect, 0x00007a79)

	if lg.thrust {
		thrusterRect := sdl.Rect{posX + 3, posY + 16, 5, 3}
		lg.surface.FillRect(&thrusterRect, 0x00ff0000)
	}
}
