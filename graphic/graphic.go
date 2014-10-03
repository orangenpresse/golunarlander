package graphic

import (
	_ "fmt"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type Vector2D struct {
	X float64
	Y float64
}

type SdlLander interface {
	GetPosition() Vector2D
}

type LanderGraphic struct {
	run     bool
	Width   int64
	Height  int64
	surface *sdl.Surface
	window  *sdl.Window
	Lander  SdlLander
}

func (lg *LanderGraphic) Start() {
	lg.run = true
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
	if lg.Lander == nil {
		return
	}
	landerPos := lg.Lander.GetPosition()
	landerRect := sdl.Rect{10, 10, int32(landerPos.X), int32(landerPos.Y)}
	lg.surface.FillRect(&landerRect, 0x007a7900)
}

// func main() {
// 	landergraphic := LanderGraphic{Width: 800, Height: 600}
// 	landergraphic.start()
// }
