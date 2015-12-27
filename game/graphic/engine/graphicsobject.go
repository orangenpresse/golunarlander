package engine

import (
	gl "github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type GraphicsObject struct {
	renderObject *RenderObject
	modelMatrix  mgl32.Mat4
	scaling      mgl32.Vec3
	rotation     mgl32.Vec4
	translation  mgl32.Vec3
	color        mgl32.Vec4
}

func (g *GraphicsObject) InitGraphic(renderObject *RenderObject) {
	g.renderObject = renderObject
}

func (g *GraphicsObject) SetScale(x float32, y float32, z float32) {
	g.scaling = mgl32.Vec3{x, y, z}
}

func (g *GraphicsObject) SetRotation(deg float32, x float32, y float32, z float32) {
	g.rotation = mgl32.Vec4{x, y, z, deg}
}

func (g *GraphicsObject) SetTranslation(x float32, y float32, z float32) {
	g.translation = mgl32.Vec3{x, y, z}
}

func (g *GraphicsObject) SetColor(red float32, green float32, blue float32, alpha float32) {
	g.color = mgl32.Vec4{red, green, blue, alpha}
}

func (g *GraphicsObject) reset() {
	g.modelMatrix = mgl32.Ident4()
}

func (g *GraphicsObject) scale() {
	scaleVector := g.scaling
	scaleMatrix := mgl32.Scale3D(scaleVector.X(), scaleVector.Y(), scaleVector.Z())
	g.modelMatrix = scaleMatrix.Mul4(g.modelMatrix)
}

func (g *GraphicsObject) rotate() {
	rotationVector := g.rotation
	rotMatrix := mgl32.HomogRotate3D(rotationVector.W(), rotationVector.Vec3())
	g.modelMatrix = rotMatrix.Mul4(g.modelMatrix)
}

func (g *GraphicsObject) translate() {
	translationVector := g.translation
	translationMatrix := mgl32.Translate3D(translationVector.X(), translationVector.Y(), translationVector.Z())
	g.modelMatrix = translationMatrix.Mul4(g.modelMatrix)
}

func (g *GraphicsObject) applyModelUniform() {
	location := g.renderObject.GetUniformLocation("model")
	gl.UniformMatrix4fv(location, 1, false, &g.modelMatrix[0])

}

func (g *GraphicsObject) applyColor() {
	colorVector := g.color
	location := g.renderObject.GetUniformLocation("color")
	color := []float32{colorVector.X(), colorVector.Y(), colorVector.Z(), colorVector.W()}
	gl.Uniform4fv(location, 1, &color[0])
	g.renderObject.BindFragDataLocation(0, "outColor")
}

func (g *GraphicsObject) drawTriangles() {
	g.renderObject.Draw()
}

func (g *GraphicsObject) Draw() {
	g.reset()
	g.scale()
	g.rotate()
	g.translate()
	g.applyModelUniform()
	g.applyColor()
	g.drawTriangles()
}
