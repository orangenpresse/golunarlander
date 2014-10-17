package game

import (
	"fmt"
	gl "github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/orangenpresse/golunarlander/simulation"
	"io/ioutil"
)

type Graphic struct {
	shaderVersion     string
	verticies         []float32
	vao               gl.VertexArray
	vbo               gl.Buffer
	vertex_shader     gl.Shader
	fragment_shader   gl.Shader
	frameBufferHeight int
	frameBufferWidth  int
	program           gl.Program
	Lander            *simulation.Lander
}

func (lg *LunarLanderGame) initGraphics(shaderVersion string) {
	lg.Graphic = new(Graphic)
	g := lg.Graphic
	g.shaderVersion = shaderVersion
	g.Lander = lg.Simulation.GetLander()

	g.frameBufferWidth, g.frameBufferHeight = lg.window.GetFramebufferSize()

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

func (g *Graphic) getShaders() (vertexShader string, fragmentShader string) {
	if data, err := ioutil.ReadFile("./game/shader/vertexShader" + g.shaderVersion + ".glsl"); err != nil {
		fmt.Println("VertexShader Read Error:" + err.Error())
		panic("VertexShader not found")
	} else {
		vertexShader = string(data)
	}

	if data, err := ioutil.ReadFile("./game/shader/fragmentShader" + g.shaderVersion + ".glsl"); err != nil {
		fmt.Println("FragmentShader Read Error:" + err.Error())
		panic("FragmentShader not found")
	} else {
		fragmentShader = string(data)
	}

	return vertexShader, fragmentShader
}

func (g *Graphic) compileShaders() {
	vertexShader, fragmentShader := g.getShaders()

	g.vertex_shader = gl.CreateShader(gl.VERTEX_SHADER)
	g.vertex_shader.Source(vertexShader)
	g.vertex_shader.Compile()
	if info := g.vertex_shader.GetInfoLog(); info != "" {
		fmt.Println(info)
	}

	g.fragment_shader = gl.CreateShader(gl.FRAGMENT_SHADER)
	g.fragment_shader.Source(fragmentShader)
	g.fragment_shader.Compile()
	if info := g.fragment_shader.GetInfoLog(); info != "" {
		fmt.Println(info)
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
	g.drawHud()

}

func (g *Graphic) clear() {
	gl.Viewport(0, 0, g.frameBufferWidth, g.frameBufferHeight)
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

func (g *Graphic) drawHud() {
	posY := float32(-1.68)
	posX := float32(-3)

	// Draw outer
	model := mgl32.Ident4()

	scaleMatrix := mgl32.Scale3D(0.08, 0.4, 1)
	model = scaleMatrix.Mul4(model)

	translationMatrix := mgl32.Translate3D(posX, posY, 0.0)
	model = translationMatrix.Mul4(model)

	modelUniform := g.program.GetUniformLocation("model")
	modelUniform.UniformMatrix4fv(false, [16]float32(model))

	positionAttrib := g.program.GetAttribLocation("position")
	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	positionAttrib.EnableArray()
	defer positionAttrib.DisableArray()

	color := g.program.GetUniformLocation("color")
	color.Uniform4fv(1, []float32{0.3, 0.3, 0.3, 0})
	g.program.BindFragDataLocation(0, "outColor")

	gl.DrawArrays(gl.TRIANGLES, 0, 6)

	// Draw inner
	model2 := mgl32.Ident4()

	scaleMatrix2 := mgl32.Scale3D(0.06, 0.38, 1)
	model2 = scaleMatrix2.Mul4(model2)

	translationMatrix2 := mgl32.Translate3D(posX, posY, 0.0)
	model2 = translationMatrix2.Mul4(model2)

	modelUniform2 := g.program.GetUniformLocation("model")
	modelUniform2.UniformMatrix4fv(false, [16]float32(model2))

	positionAttrib2 := g.program.GetAttribLocation("position")
	positionAttrib2.AttribPointer(3, gl.FLOAT, false, 0, nil)
	positionAttrib2.EnableArray()
	defer positionAttrib2.DisableArray()

	color2 := g.program.GetUniformLocation("color")
	color2.Uniform4fv(1, []float32{0.5, 0.5, 0.5, 0})
	g.program.BindFragDataLocation(0, "outColor")

	gl.DrawArrays(gl.TRIANGLES, 0, 6)

	g.drawFuelBar(posX, posY)
}

func (g *Graphic) drawFuelBar(posX float32, posY float32) {
	fuel := float32(g.Lander.GetFuelLevel()) / 100

	model := mgl32.Ident4()

	scaleMatrix := mgl32.Scale3D(0.06, 0.38*fuel, 1)
	model = scaleMatrix.Mul4(model)

	translationMatrix := mgl32.Translate3D(posX, (posY+0.38*fuel)-0.38, 0.0)
	model = translationMatrix.Mul4(model)

	modelUniform := g.program.GetUniformLocation("model")
	modelUniform.UniformMatrix4fv(false, [16]float32(model))

	positionAttrib := g.program.GetAttribLocation("position")
	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	positionAttrib.EnableArray()
	defer positionAttrib.DisableArray()

	color := g.program.GetUniformLocation("color")
	color.Uniform4fv(1, []float32{0.3, 1, 0.3, 0})
	g.program.BindFragDataLocation(0, "outColor")

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

func (g *Graphic) drawLander() {
	landerPos := g.Lander.GetPosition()
	posY := float32(landerPos.Y/200) - 2.1
	posX := float32(landerPos.X / 200)

	model := mgl32.Ident4()

	scaleMatrix := mgl32.Scale3D(0.04, 0.06, 1)
	model = scaleMatrix.Mul4(model)

	translationMatrix := mgl32.Translate3D(posX, posY, 0.0)
	model = translationMatrix.Mul4(model)

	rotMatrix := mgl32.HomogRotate3D(0.0, mgl32.Vec3{0, 0, 1})
	model = rotMatrix.Mul4(model)

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

	g.drawThrust(posX, posY)
	g.drawExploded(posX, posY)
}

func (g *Graphic) drawThrust(posX float32, posY float32) {
	thrusterState := g.Lander.IsThrusting()

	if thrusterState.Bottom {
		g.drawThrusterFlame(posX, posY-0.08)
	}

	if thrusterState.Left {
		g.drawThrusterFlame(posX-0.05, posY)
	}

	if thrusterState.Right {
		g.drawThrusterFlame(posX+0.05, posY)
	}
}

func (g *Graphic) drawThrusterFlame(posX float32, posY float32) {
	model := mgl32.Ident4()

	scaleMatrix := mgl32.Scale3D(0.03, 0.01, 1)
	model = scaleMatrix.Mul4(model)

	translationMatrix := mgl32.Translate3D(posX, posY, 0.0)
	model = translationMatrix.Mul4(model)

	modelUniform := g.program.GetUniformLocation("model")
	modelUniform.UniformMatrix4fv(false, [16]float32(model))

	color := g.program.GetUniformLocation("color")
	color.Uniform4fv(1, []float32{1, 0, 0, 0})
	g.program.BindFragDataLocation(0, "outColor")

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

func (g *Graphic) drawExploded(posX float32, posY float32) {
	if g.Lander.IsExploded() {
		model := mgl32.Ident4()

		scaleMatrix := mgl32.Scale3D(0.1, 0.1, 1)
		model = scaleMatrix.Mul4(model)

		translationMatrix := mgl32.Translate3D(posX, posY, 0.0)
		model = translationMatrix.Mul4(model)

		modelUniform := g.program.GetUniformLocation("model")
		modelUniform.UniformMatrix4fv(false, [16]float32(model))

		color := g.program.GetUniformLocation("color")
		color.Uniform4fv(1, []float32{1, 1, 0, 0})
		g.program.BindFragDataLocation(0, "outColor")

		gl.DrawArrays(gl.TRIANGLES, 0, 6)
	}
}
