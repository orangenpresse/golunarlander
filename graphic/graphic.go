package graphic

import (
	"github.com/orangenpresse/golunarlander/simulation"
	"github.com/veandco/go-sdl2/sdl"
	_ "github.com/veandco/go-sdl2/sdl_ttf"
	"reflect"
	"time"
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
			if ev, _ := event.(*sdl.KeyDownEvent); ev.Keysym.Scancode == 82 {
				lg.thrust = true
			}

		case "*sdl.KeyUpEvent":
			if ev, _ := event.(*sdl.KeyUpEvent); ev.Keysym.Scancode == 82 {
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

func (lg *LanderGraphic) renderLander() {
	landerPos := lg.Simulation.GetLander().GetPosition()
	posX := int32(landerPos.X)
	posY := int32(lg.Height-25) - int32(landerPos.Y)

	landerRect := sdl.Rect{posX, posY, 10, 15}
	lg.surface.FillRect(&landerRect, 0x00007a79)

	//ttf.RenderText_Solid("meep", 0x00ff0000)

	if lg.thrust {
		thrusterRect := sdl.Rect{posX + 3, posY + 16, 5, 3}
		lg.surface.FillRect(&thrusterRect, 0x00ff0000)
	}
}
