package game

import (
	gl "github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
	"unsafe"
)

const FLOAT32_BYTE_SIZE int = int(unsafe.Sizeof(float32(0)))

type Rect struct {
	modelMatrix   mgl32.Mat4
	scaling       mgl32.Vec3
	rotation      mgl32.Vec4
	translation   mgl32.Vec3
	color         mgl32.Vec4
	shaderProgram *gl.Program
}

func LoadToRam() {
	verticies := []float32{
		-1, -1, 0,
		1, -1, 0,
		-1, 1, 0,
		-1, 1, 0,
		1, 1, 0,
		1, -1, 0}
	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*FLOAT32_BYTE_SIZE, verticies, gl.STATIC_DRAW)
}

func NewRect() *Rect {
	r := new(Rect)
	r.modelMatrix = mgl32.Ident4()
	return r
}

func (r *Rect) SetScale(x float32, y float32, z float32) {
	r.scaling = mgl32.Vec3{x, y, z}
}

func (r *Rect) SetRotation(deg float32, x float32, y float32, z float32) {
	r.rotation = mgl32.Vec4{x, y, z, deg}
}

func (r *Rect) SetTranslation(x float32, y float32, z float32) {
	r.translation = mgl32.Vec3{x, y, z}
}

func (r *Rect) SetColor(red float32, green float32, blue float32, alpha float32) {
	r.color = mgl32.Vec4{red, green, blue, alpha}
}

func (r *Rect) scale() {
	scaleVector := r.scaling
	scaleMatrix := mgl32.Scale3D(scaleVector.X(), scaleVector.Y(), scaleVector.Z())
	r.modelMatrix = scaleMatrix.Mul4(r.modelMatrix)
}

func (r *Rect) rotate() {
	rotationVector := r.rotation
	rotMatrix := mgl32.HomogRotate3D(rotationVector.W(), rotationVector.Vec3())
	r.modelMatrix = rotMatrix.Mul4(r.modelMatrix)
}

func (r *Rect) translate() {
	translationVector := r.translation
	translationMatrix := mgl32.Translate3D(translationVector.X(), translationVector.Y(), translationVector.Z())
	r.modelMatrix = translationMatrix.Mul4(r.modelMatrix)
}

func (r *Rect) applyModelUniform() {
	modelUniform := r.shaderProgram.GetUniformLocation("model")
	modelUniform.UniformMatrix4fv(false, [16]float32(r.modelMatrix))
}

func (r *Rect) applyColor() {
	colorVector := r.color
	color := r.shaderProgram.GetUniformLocation("color")
	color.Uniform4fv(1, []float32{colorVector.X(), colorVector.Y(), colorVector.Z(), colorVector.W()})
	r.shaderProgram.BindFragDataLocation(0, "outColor")
}

func (r *Rect) drawTriangles() {
	positionAttrib := r.shaderProgram.GetAttribLocation("position")
	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	positionAttrib.EnableArray()
	defer positionAttrib.DisableArray()

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

func (r *Rect) Draw() {
	r.scale()
	r.rotate()
	r.translate()
	r.applyModelUniform()
	r.applyColor()
	r.drawTriangles()
}
