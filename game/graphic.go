package game

import (
	"fmt"
	gl "github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/orangenpresse/golunarlander/simulation"
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
	program           gl.Program
	Lander            *simulation.Lander
}

func (lg *LunarLander) initGraphics() {
	g := Graphic{}
	lg.Graphic = &g
	g.Lander = lg.Simulation.GetLander()

	g.frameBufferWidht, g.frameBufferHeight = lg.window.GetFramebufferSize()

	gl.Init()
	g.initBuffers()
	g.compileShaders()
	verticies := []float32{
		-1, -1, 0,
		1, -1, 0,
		-1, 1, 0,
		-1, 1, 0,
		1, 1, 0,
		1, -1, 0}
	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*4, verticies, gl.STATIC_DRAW)
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
		if info := g.vertex_shader.GetInfoLog(); info != "" {
			fmt.Println(info)
		}
	}

	if data, err := ioutil.ReadFile("./game/shader/fragmentShader.glsl"); err != nil {
		fmt.Println("FragmentShader Read Error:" + err.Error())
	} else {
		g.fragment_shader = gl.CreateShader(gl.FRAGMENT_SHADER)
		g.fragment_shader.Source(string(data))
		g.fragment_shader.Compile()
		if info := g.fragment_shader.GetInfoLog(); info != "" {
			fmt.Println(info)
		}
	}

	g.program = gl.CreateProgram()
	g.program.AttachShader(g.vertex_shader)
	g.program.AttachShader(g.fragment_shader)
	g.program.Link()
	g.program.Use()

}

func (g *Graphic) end() {
	g.fragment_shader.Delete()
	g.vertex_shader.Delete()
	g.program.Delete()
}

func (g *Graphic) render() {
	g.clear()
	g.setPerspectiveAndCamera()
	g.drawMoonSurface()
	g.drawLander()
	// g.drawHud()

}

func (g *Graphic) clear() {
	gl.Viewport(0, 0, g.frameBufferWidht, g.frameBufferHeight)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

}

func (g *Graphic) setPerspectiveAndCamera() {
	projection := mgl32.Perspective(70.0, float32(800)/600, 0.1, 10.0)
	projectionUniform := g.program.GetUniformLocation("projection")
	projectionUniform.UniformMatrix4fv(false, [16]float32(projection))

	camera := mgl32.LookAtV(mgl32.Vec3{0, 0, 5}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cameraUniform := g.program.GetUniformLocation("camera")
	cameraUniform.UniformMatrix4fv(false, [16]float32(camera))
}

func (g *Graphic) drawMoonSurface() {
	model := mgl32.Ident4()

	translationMatrix := mgl32.Translate3D(0.0, -8.2, 0.0)
	model = translationMatrix.Mul4(model)

	rotMatrix := mgl32.HomogRotate3D(0.0, mgl32.Vec3{0, 0, 1})
	model = rotMatrix.Mul4(model)

	scaleMatrix := mgl32.Scale3D(10.0, 0.3, 1)
	model = scaleMatrix.Mul4(model)

	modelUniform := g.program.GetUniformLocation("model")
	modelUniform.UniformMatrix4fv(false, [16]float32(model))

	positionAttrib := g.program.GetAttribLocation("position")
	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	positionAttrib.EnableArray()
	defer positionAttrib.DisableArray()

	color := g.program.GetUniformLocation("color")
	color.Uniform4fv(1, []float32{0.7, 0.5, 0, 0})
	g.program.BindFragDataLocation(0, "outColor")

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
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

func (g *Graphic) drawLander() {
	landerPos := g.Lander.GetPosition()
	posY := landerPos.Y/10 - 35
	posX := 0

	model := mgl32.Ident4()

	translationMatrix := mgl32.Translate3D(float32(posX), float32(posY), 0.0)
	model = translationMatrix.Mul4(model)

	rotMatrix := mgl32.HomogRotate3D(0.0, mgl32.Vec3{0, 0, 1})
	model = rotMatrix.Mul4(model)

	scaleMatrix := mgl32.Scale3D(0.04, 0.06, 1)
	model = scaleMatrix.Mul4(model)

	modelUniform := g.program.GetUniformLocation("model")
	modelUniform.UniformMatrix4fv(false, [16]float32(model))

	positionAttrib := g.program.GetAttribLocation("position")
	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	positionAttrib.EnableArray()
	defer positionAttrib.DisableArray()

	color := g.program.GetUniformLocation("color")
	color.Uniform4fv(1, []float32{0.7, 0.5, 1, 0})
	g.program.BindFragDataLocation(0, "outColor")

	gl.DrawArrays(gl.TRIANGLES, 0, 6)

	g.drawThrust(model)

	// lg.drawExploded(posX, posY)
}

func (g *Graphic) drawThrust(model mgl32.Mat4) {
	if g.Lander.IsThrusting() {

		translationMatrix := mgl32.Translate3D(0.0, -0.1, 0.0)
		model = translationMatrix.Mul4(model)

		modelUniform := g.program.GetUniformLocation("model")
		modelUniform.UniformMatrix4fv(false, [16]float32(model))

		color := g.program.GetUniformLocation("color")
		color.Uniform4fv(1, []float32{1, 0, 0, 0})
		g.program.BindFragDataLocation(0, "outColor")

		gl.DrawArrays(gl.TRIANGLES, 0, 6)
	}
}

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
