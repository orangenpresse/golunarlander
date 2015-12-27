package graphic

import (
	gl "github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	data "github.com/orangenpresse/golunarlander/dataObjects"
	"github.com/orangenpresse/golunarlander/game/graphic/engine"
	"github.com/orangenpresse/golunarlander/game/graphic/model"
	"github.com/orangenpresse/golunarlander/lander"
)

type Graphic struct {
	modelManager    *engine.ModelManager
	shaderVersion   string
	vertex_shader   uint32
	fragment_shader uint32
	program         uint32
	Lander          lander.LanderInterface
	OtherLanders    []lander.LanderInterface
	Options         *data.Options
	rect            *model.Rect
}

func NewGraphic(options *data.Options, shaderVersion string, lander lander.LanderInterface) *Graphic {
	gl.Init()
	g := new(Graphic)
	g.Options = options
	g.shaderVersion = shaderVersion
	g.Lander = lander
	g.createProgram()
	g.initModels()
	return g
}

func (g *Graphic) createProgram() {
	var err error
	g.vertex_shader, err = engine.NewShader("./game/graphic/shader/vertexShader", g.shaderVersion, gl.VERTEX_SHADER)
	if err != nil {
		panic(err.Error())
	}

	g.fragment_shader, err = engine.NewShader("./game/graphic/shader/fragmentShader", g.shaderVersion, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err.Error())
	}

	g.program = gl.CreateProgram()
	gl.AttachShader(g.program, g.vertex_shader)
	gl.AttachShader(g.program, g.fragment_shader)
	gl.LinkProgram(g.program)
	gl.UseProgram(g.program)
}

func (g *Graphic) initModels() {
	g.modelManager = engine.NewModelManager()
	g.modelManager.RegisterModel("rect", model.Rect{}.LoadToVram(g.program))
	g.rect = model.NewRect(g.modelManager.GetRenderObject("rect"))
}

func (g *Graphic) End() {
	gl.DeleteShader(g.fragment_shader)
	gl.DeleteShader(g.vertex_shader)
	gl.DeleteProgram(g.program)
}

func (g *Graphic) Render(landers []lander.LanderInterface) {
	g.clear()
	g.setPerspectiveAndCamera()
	g.drawMoonSurface()
	g.drawLander(nil)

	for _, playerLander := range landers {
		g.drawLander(playerLander)
	}

	g.drawHud()

}

func (g *Graphic) clear() {
	gl.Viewport(0, 0, g.Options.BufferWidth, g.Options.BufferHeight)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (g *Graphic) setPerspectiveAndCamera() {
	projection := mgl32.Perspective(70.0, float32(800)/600, 0.1, 10.0)
	projectionUniform := gl.GetUniformLocation(g.program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	camera := mgl32.LookAtV(mgl32.Vec3{0, 0, 5}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cameraUniform := gl.GetUniformLocation(g.program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])
}

func (g *Graphic) drawTest() {
	r := model.NewRect(g.modelManager.GetRenderObject("rect"))
	r.SetColor(1.0, 0, 0, 0)
	r.SetScale(1, 1, 1)
	r.SetTranslation(0, 0, 0)
	r.Draw()
}

func (g *Graphic) drawMoonSurface() {
	g.rect.SetColor(0.7, 0.5, 0, 0)
	g.rect.SetScale(10.0, 0.3, 1)
	g.rect.SetTranslation(0.0, -2.45, 0.0)
	g.rect.Draw()
}

func (g *Graphic) drawHud() {
	posY := float32(-1.68)
	posX := float32(-3)

	// Draw Outer
	if g.Options.DebugMode {
		g.rect.SetColor(0.0, 0.3, 1.0, 0)
	} else {
		g.rect.SetColor(0.3, 0.3, 0.3, 0)
	}
	g.rect.SetScale(0.08, 0.4, 1)
	g.rect.SetTranslation(posX, posY, 0)
	g.rect.Draw()

	// Draw inner
	g.rect.SetColor(0.5, 0.5, 0.5, 0)
	g.rect.SetScale(0.06, 0.38, 1)
	g.rect.SetTranslation(posX, posY, 0.0)
	g.rect.Draw()

	g.drawFuelBar(posX, posY)
}

func (g *Graphic) drawFuelBar(posX float32, posY float32) {
	fuel := float32(g.Lander.GetFuelLevel()) / 100
	factor := float32(0.38)

	g.rect.SetScale(0.06, factor*fuel, 1)
	g.rect.SetTranslation(posX, (posY+factor*fuel)-factor, 0.0)
	g.rect.SetColor(1-fuel, fuel, 0.1, 0)
	g.rect.Draw()
}

func (g *Graphic) drawLander(lander lander.LanderInterface) {
	var landerPos data.Vector2D
	if lander != nil {
		landerPos = lander.GetPosition()
	} else {
		landerPos = g.Lander.GetPosition()
	}

	posY := float32(landerPos.Y/200) - 2.1
	posX := float32(landerPos.X / 200)

	g.rect.SetColor(0.7, 0.5, 1, 0)
	g.rect.SetScale(0.04, 0.06, 1)
	g.rect.SetTranslation(posX, posY, 0.0)
	g.rect.SetRotation(0.0, 0, 0, 1)
	g.rect.Draw()

	// Reset
	g.rect.SetRotation(0, 0, 0, 0)

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
	g.rect.SetScale(0.03, 0.01, 1)
	g.rect.SetTranslation(posX, posY, 0.0)
	g.rect.SetColor(1, 0, 0, 0)
	g.rect.Draw()

}

func (g *Graphic) drawExploded(posX float32, posY float32) {
	if g.Lander.IsExploded() {
		g.rect.SetScale(0.1, 0.1, 1)
		g.rect.SetTranslation(posX, posY, 0.0)
		g.rect.SetColor(1, 1, 0, 0)
		g.rect.Draw()
	}
}
