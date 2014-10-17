package game

// import (
// 	"github.com/go-gl/mathgl/mgl32"
// )

// type Rect struct {
// 	modelMatrix mgl32.Mat4
// 	scale       mgl32.Vec3
// 	rotation    mgl32.Vec3
// 	translation mgl32.Vec3
// }

// func NewRect() *Rect {
// 	r := new(Rect)
// 	r.modelMatrix = mgl32.Ident4()
// }

// func (r *Rect) scale() {
// 	scaleMatrix := mgl32.Scale3D(0.06, 0.38*fuel, 1)
// 	model = scaleMatrix.Mul4(model)
// }

// func (r *Rect) rotate() {
// 	rotMatrix := mgl32.HomogRotate3D(0.0, mgl32.Vec3{0, 0, 1})
// 	model = rotMatrix.Mul4(model)
// }

// func (r *Rect) translate() {
// 	translationMatrix := mgl32.Translate3D(posX, (posY+0.38*fuel)-0.38, 0.0)
// 	model = translationMatrix.Mul4(model)

// }

// func (r *Rect) setModelUniform() {
// 	modelUniform := g.program.GetUniformLocation("model")
// 	modelUniform.UniformMatrix4fv(false, [16]float32(model))
// }

// func (r *Rect) setColor() {
// 	color := g.program.GetUniformLocation("color")
// 	color.Uniform4fv(1, []float32{0.3, 1, 0.3, 0})
// 	g.program.BindFragDataLocation(0, "outColor")
// }

// func (r *Rect) drawTriangles() {
// 	positionAttrib := g.program.GetAttribLocation("position")
// 	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
// 	positionAttrib.EnableArray()
// 	defer positionAttrib.DisableArray()

// 	gl.DrawArrays(gl.TRIANGLES, 0, 6)
// }

// func (r *Rect) Draw() {
// 	r.scale()
// 	r.rotate()
// 	r.translate()
// 	r.setModelUniform()
// 	r.setColor()
// 	r.drawTriangles()
// }
