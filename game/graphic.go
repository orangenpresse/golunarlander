package game

import (
	"fmt"
	gl "github.com/go-gl/gl"
	"io/ioutil"
)

type Graphic struct {
	verticies         []float32
	vao               gl.VertexArray
	vbo               gl.Buffer
	vertex_shader     gl.Shader
	fragment_shader   gl.Shader
	frameBufferHeight int
	frameBufferWidht  int
}

func (lg *LunarLander) initGraphics() {
	g := Graphic{}
	lg.Graphic = &g

	g.frameBufferWidht, g.frameBufferHeight = lg.window.GetFramebufferSize()

	gl.Init()
	g.initBuffers()
	g.compileShaders()
	verticies := []float32{
		0, 1, 0,
		-1, -1, 0,
		1, -1, 0,
		0, -1, 0,
		1, 1, 0,
		-1, 1, 0}
	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*6, verticies, gl.STATIC_DRAW)
}

func (g *Graphic) initBuffers() {
	g.vao = gl.GenVertexArray()
	g.vao.Bind()

	g.vbo = gl.GenBuffer()
	g.vbo.Bind(gl.ARRAY_BUFFER)
}

func (g *Graphic) compileShaders() {
	if data, err := ioutil.ReadFile("./game/shader/vertexShader.glsl"); err != nil {
		fmt.Println("VertexShader Read Error:" + err.Error())
	} else {
		g.vertex_shader = gl.CreateShader(gl.VERTEX_SHADER)
		g.vertex_shader.Source(string(data))
		g.vertex_shader.Compile()
		fmt.Println(g.vertex_shader.GetInfoLog())
	}

	if data, err := ioutil.ReadFile("./game/shader/fragmentShader.glsl"); err != nil {
		fmt.Println("FragmentShader Read Error:" + err.Error())
	} else {
		g.fragment_shader = gl.CreateShader(gl.FRAGMENT_SHADER)
		g.fragment_shader.Source(string(data))
		g.fragment_shader.Compile()
		fmt.Println(g.fragment_shader.GetInfoLog())
	}

}

func (g *Graphic) end() {
	g.fragment_shader.Delete()
	g.vertex_shader.Delete()
}

func (g *Graphic) render() {
	g.clear()
	g.drawMoonSurface()
	// lg.drawLander()
	// lg.drawHud()

}

func (g *Graphic) clear() {
	gl.Viewport(0, 0, g.frameBufferWidht, g.frameBufferHeight)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (g *Graphic) drawMoonSurface() {
	program := gl.CreateProgram()
	program.AttachShader(g.vertex_shader)
	program.AttachShader(g.fragment_shader)

	program.BindFragDataLocation(0, "outColor")
	program.Link()
	program.Use()
	defer program.Delete()

	positionAttrib := program.GetAttribLocation("position")
	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	positionAttrib.EnableArray()
	defer positionAttrib.DisableArray()

	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	// surfaceRect := sdl.Rect{0, 590, 800, 10}
	// lg.surface.FillRect(&surfaceRect, 0x007a5345)
}

// func (lg *Graphic) drawHud() {
// 	posY := lg.Height - 125

// 	bg := sdl.Rect{int32(5), int32(posY), 20, 110}
// 	lg.surface.FillRect(&bg, 0x00878b88)

// 	lg.drawFuelBar(posY)
// }

// func (lg *Graphic) drawFuelBar(posY int64) {
// 	fuelBorder := sdl.Rect{int32(8), int32(posY + 3), 14, 104}
// 	lg.surface.FillRect(&fuelBorder, 0x00c3c9c4)

// 	fuel := int32(lg.Simulation.GetLander().GetFuelLevel())

// 	fuelBar := sdl.Rect{int32(10), (100 - fuel) + int32(posY+5), 10, fuel}
// 	lg.surface.FillRect(&fuelBar, 0x0000de3c)
// }

// func (lg *Graphic) drawLander() {
// 	landerPos := lg.Simulation.GetLander().GetPosition()
// 	posX := int32(landerPos.X)
// 	posY := int32(lg.Height-25) - int32(landerPos.Y)

// 	landerRect := sdl.Rect{posX, posY, 10, 15}
// 	lg.surface.FillRect(&landerRect, 0x00007a79)

// 	lg.drawThrust(posX, posY)
// 	lg.drawExploded(posX, posY)
// }

// func (lg *Graphic) drawThrust(posX int32, posY int32) {
// 	if lg.Simulation.GetLander().IsThrusting() {
// 		thrusterRect := sdl.Rect{posX + 3, posY + 16, 5, 3}
// 		lg.surface.FillRect(&thrusterRect, 0x00ff0000)
// 	}
// }

// func (lg *Graphic) drawExploded(posX int32, posY int32) {
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
