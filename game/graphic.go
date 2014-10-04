package game

import (
//gl "github.com/go-gl/gl"
)

func (lg *LunarLander) render() {
	//width, height := lg.window.GetFramebufferSize()
	// gl.Viewport(0, 0, width, height)
	// gl.Clear(gl.COLOR_BUFFER_BIT)

	//gl.DrawArrays(gl.TRIANGLE_STRIP, 0, len(triangleVertices))

	//lg.window.SwapBuffers()

	// lg.clearSurface()
	// lg.drawMoonSurface()
	// lg.drawLander()
	// lg.drawHud()
}

// func (lg *LunarLander) clearSurface() {
// 	rect := sdl.Rect{0, 0, int32(lg.Width), int32(lg.Height)}
// 	lg.surface.FillRect(&rect, 0x00000000)
// }

// func (lg *LunarLander) drawMoonSurface() {
// 	surfaceRect := sdl.Rect{0, 590, 800, 10}
// 	lg.surface.FillRect(&surfaceRect, 0x007a5345)
// }

// func (lg *LunarLander) drawHud() {
// 	posY := lg.Height - 125

// 	bg := sdl.Rect{int32(5), int32(posY), 20, 110}
// 	lg.surface.FillRect(&bg, 0x00878b88)

// 	lg.drawFuelBar(posY)
// }

// func (lg *LunarLander) drawFuelBar(posY int64) {
// 	fuelBorder := sdl.Rect{int32(8), int32(posY + 3), 14, 104}
// 	lg.surface.FillRect(&fuelBorder, 0x00c3c9c4)

// 	fuel := int32(lg.Simulation.GetLander().GetFuelLevel())

// 	fuelBar := sdl.Rect{int32(10), (100 - fuel) + int32(posY+5), 10, fuel}
// 	lg.surface.FillRect(&fuelBar, 0x0000de3c)
// }

// func (lg *LunarLander) drawLander() {
// 	landerPos := lg.Simulation.GetLander().GetPosition()
// 	posX := int32(landerPos.X)
// 	posY := int32(lg.Height-25) - int32(landerPos.Y)

// 	landerRect := sdl.Rect{posX, posY, 10, 15}
// 	lg.surface.FillRect(&landerRect, 0x00007a79)

// 	lg.drawThrust(posX, posY)
// 	lg.drawExploded(posX, posY)
// }

// func (lg *LunarLander) drawThrust(posX int32, posY int32) {
// 	if lg.Simulation.GetLander().IsThrusting() {
// 		thrusterRect := sdl.Rect{posX + 3, posY + 16, 5, 3}
// 		lg.surface.FillRect(&thrusterRect, 0x00ff0000)
// 	}
// }

// func (lg *LunarLander) drawExploded(posX int32, posY int32) {
// 	if lg.Simulation.GetLander().IsExploded() {
// 		p1x, p1y, p2x, p2y := 0, 0, int(lg.Width), int(lg.Height)

// 		renderer := lg.window.GetRenderer()
// 		renderer.SetDrawColor(255, 128, 0, 0)
// 		renderer.DrawLine(p1x, p1y, p2x, p2y)
// 		renderer.DrawLine(p1x, p2y, p2x, p1y)
// 		renderer.DrawLine(p2x/2, p1y, p2x/2, p2y)

// 		// TODO ALL TO RENDERER
// 		renderer.Present()
// 	}
// }
