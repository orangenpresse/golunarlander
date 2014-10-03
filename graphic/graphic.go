package graphic

import (
	_ "fmt"
	"github.com/orangenpresse/golunarlander/simulation"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type SdlLander interface {
	GetPosition() simulation.Vector2D
}

type Simulation interface {
	Start()
	Update(int64)
	GetLander() SdlLander
}

type LanderGraphic struct {
	run        bool
	Width      int64
	Height     int64
	surface    *sdl.Surface
	window     *sdl.Window
	Simulation Simulation
}

func (lg *LanderGraphic) Start() {
	lg.run = true
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
		if lg.Simulation != nil {
			lg.Simulation.Update(100)
		}
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
