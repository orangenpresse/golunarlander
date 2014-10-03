package graphic

import (
	_ "fmt"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type Vector2D struct {
	x float64
	y float64
}

type Lander interface {
	GetPosition() Vector2D
}

type SdlLander struct {
	Lander
}

type LanderGrapic struct {
	run     bool
	width   int64
	height  int64
	surface *sdl.Surface
	window  *sdl.Window
	Lander  *SdlLander
}

func (lg *LanderGrapic) start() {
	lg.run = true
	lg.window = sdl.CreateWindow("Lunar Lander", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int(lg.width), int(lg.height), sdl.WINDOW_SHOWN)
	lg.surface = lg.window.GetSurface()
	lg.render()
	lg.end()
}

func (lg *LanderGrapic) end() {
	lg.window.Destroy()
}

func (lg *LanderGrapic) render() {
	for lg.run == true {
		lg.clearSurface()
		lg.RenderMoonSurface()
		//lg.RenderLander()
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

func (lg *LanderGrapic) clearSurface() {
	rect := sdl.Rect{0, 0, int32(lg.width), int32(lg.height)}
	lg.surface.FillRect(&rect, 0x00000000)
}

func (lg *LanderGrapic) RenderMoonSurface() {
	surfaceRect := sdl.Rect{0, 590, 800, 10}
	lg.surface.FillRect(&surfaceRect, 0x7a534500)
}

func (lg *LanderGrapic) RenderLander() {
	landerPos := lg.Lander.GetPosition()
	landerRect := sdl.Rect{10, 10, int32(landerPos.x), int32(landerPos.y)}
	lg.surface.FillRect(&landerRect, 0x007a7900)
}

// func main() {
// 	landergraphic := LanderGrapic{width: 800, height: 600}
// 	landergraphic.start()
// }
