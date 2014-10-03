package graphic

import (
	"fmt"
	"github.com/orangenpresse/golunarlander/simulation"
	"github.com/veandco/go-sdl2/sdl"
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
	Width      int64
	Height     int64
	surface    *sdl.Surface
	window     *sdl.Window
	timer      Timer
	Simulation Simulation
}

func (lg *LanderGraphic) Start() {
	lg.run = true
	lg.timer.Start()
	if lg.Simulation != nil {
		lg.Simulation.Start()
	}
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
		if lg.Simulation != nil {
			lg.Simulation.Update(lg.timer.GetDelta(), true)
		}
		//fmt.Print(lg.timer.GetDelta())
		lg.clearSurface()
		lg.renderMoonSurface()
		lg.renderLander()
		lg.window.UpdateSurface()

		event := sdl.PollEvent()
		if event != nil {
			eventType := reflect.TypeOf(event).String()
			switch eventType {
			case "*sdl.QuitEvent":
				lg.run = false
			default:
				fmt.Println(eventType)
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
	lg.surface.FillRect(&surfaceRect, 0x7a534500)
}

func (lg *LanderGraphic) renderLander() {
	if lg.Simulation == nil {
		return
	}
	landerPos := lg.Simulation.GetLander().GetPosition()
	landerRect := sdl.Rect{10, 10, int32(landerPos.X), int32(landerPos.Y)}
	lg.surface.FillRect(&landerRect, 0x007a7900)
}
