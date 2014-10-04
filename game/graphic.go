package game

import "github.com/veandco/go-sdl2/sdl"

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

	lg.drawFuelBar()
}

func (lg *LunarLander) drawFuelBar() {
	fuelBorder := sdl.Rect{int32(8), int32(posY + 3), 14, 104}
	lg.surface.FillRect(&fuelBorder, 0x00c3c9c4)

	fuel := lg.Simulation.Lander.GetLanderState().fuel

	fuelBar := sdl.Rect{int32(10), (100 - fuel) + int32(posY+5), 10, fuel}
	lg.surface.FillRect(&fuelBar, 0x0000de3c)
}

func (lg *LunarLander) drawLander() {
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
