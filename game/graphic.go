package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

func (lg *LunarLander) render() {
	lg.clearSurface()
	lg.drawMoonSurface()
	lg.drawLander()
	lg.drawHud()
}

func (lg *LunarLander) clearSurface() {
	rect := sdl.Rect{0, 0, int32(lg.Width), int32(lg.Height)}
	lg.surface.FillRect(&rect, 0x00000000)
}

func (lg *LunarLander) drawMoonSurface() {
	surfaceRect := sdl.Rect{0, 590, 800, 10}
	lg.surface.FillRect(&surfaceRect, 0x007a5345)
}

func (lg *LunarLander) drawHud() {
	posY := lg.Height - 125

	bg := sdl.Rect{int32(5), int32(posY), 20, 110}
	lg.surface.FillRect(&bg, 0x00878b88)

	lg.drawFuelBar(posY)
}

func (lg *LunarLander) drawFuelBar(posY int64) {
	fuelBorder := sdl.Rect{int32(8), int32(posY + 3), 14, 104}
	lg.surface.FillRect(&fuelBorder, 0x00c3c9c4)

	//fuel := int32(lg.Simulation.GetLander().GetLanderState().GetFuelLevel())
	fuel := int32(100)
	fuelBar := sdl.Rect{int32(10), (100 - fuel) + int32(posY+5), 10, fuel}
	lg.surface.FillRect(&fuelBar, 0x0000de3c)
}

func (lg *LunarLander) drawLander() {
	landerPos := lg.Simulation.GetLander().GetPosition()
	posX := int32(landerPos.X)
	posY := int32(lg.Height-25) - int32(landerPos.Y)

	landerRect := sdl.Rect{posX, posY, 10, 15}
	lg.surface.FillRect(&landerRect, 0x00007a79)

	lg.drawThrust(posX, posY)
	lg.drawExploded(posX, posY)
}

func (lg *LunarLander) drawThrust(posX int32, posY int32) {
	if lg.thrust {
		thrusterRect := sdl.Rect{posX + 3, posY + 16, 5, 3}
		lg.surface.FillRect(&thrusterRect, 0x00ff0000)
	}
}

func (lg *LunarLander) drawExploded(posX int32, posY int32) {
	//if lg.Simulation.GetLander().GetLanderState().Exploded {
	for x, y := 0, 0; x < 3; {
		p1x := x
		p1y := y
		p2x := 100 - x
		p2y := 100 - y

		renderer := lg.window.GetRenderer()
		renderer.SetDrawColor(0, 255, 0, 0)
		renderer.DrawLine(p1x, p1y, p2x, p2y)
		renderer.DrawLine(0, 0, 100, 100)
		x++
		y++
		//fmt.Println(x, y)
	}
	//}
}
