package graphic

import (
	gl "github.com/go-gl/gl"
	"unsafe"
)

const FLOAT32_BYTE_SIZE int = int(unsafe.Sizeof(float32(0)))

type RenderObject struct {
	shaderProgram gl.Program
	vao           gl.VertexArray
	vbo           gl.Buffer
	vboLength     int
}

func NewRenderObject(verticies []float32, shaderProgram gl.Program) *RenderObject {
	renderObject := new(RenderObject)
	renderObject.shaderProgram = shaderProgram
	renderObject.createAndBindVbo(verticies)
	renderObject.createAndBindVoa()
	return renderObject
}

func (r *RenderObject) createAndBindVbo(verticies []float32) {
	r.vboLength = len(verticies)
	r.vbo = gl.GenBuffer()
	r.vbo.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*FLOAT32_BYTE_SIZE, verticies, gl.STATIC_DRAW)
}

func (r *RenderObject) createAndBindVoa() {
	r.vao = gl.GenVertexArray()
	r.vao.Bind()

	positionAttrib := r.shaderProgram.GetAttribLocation("position")
	positionAttrib.EnableArray()
	defer positionAttrib.DisableArray()
	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, 0)
}

func (r *RenderObject) GetShaderProgram() gl.Program {
	return r.shaderProgram
}

func (r *RenderObject) Draw() {
	r.vao.Bind()
	gl.DrawArrays(gl.TRIANGLES, 0, r.vboLength)
}
