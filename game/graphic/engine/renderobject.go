package engine

import (
	gl "github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

const FLOAT32_BYTE_SIZE int = int(unsafe.Sizeof(float32(0)))

type RenderObject struct {
	shaderProgram uint32
	vao           uint32
	vbo           uint32
	vboLength     int32
}

func NewRenderObject(verticies []float32, shaderProgram uint32) *RenderObject {
	renderObject := new(RenderObject)
	renderObject.shaderProgram = shaderProgram
	renderObject.createAndBindVbo(verticies)
	renderObject.createAndBindVoa()
	return renderObject
}

func (r *RenderObject) createAndBindVbo(verticies []float32) {
	r.vboLength = int32(len(verticies) / 3)
	gl.GenBuffers(1, &r.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*FLOAT32_BYTE_SIZE, gl.Ptr(verticies), gl.STATIC_DRAW)
}

func (r *RenderObject) createAndBindVoa() {
	gl.GenVertexArrays(1, &r.vao)
	gl.BindVertexArray(r.vao)

	positionAttrib := uint32(gl.GetAttribLocation(r.shaderProgram, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(positionAttrib)
	gl.VertexAttribPointer(positionAttrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
}

func (r *RenderObject) GetShaderProgram() uint32 {
	return r.shaderProgram
}

func (r *RenderObject) GetUniformLocation(name string) int32 {
	return gl.GetUniformLocation(r.shaderProgram, gl.Str(name+"\x00"))
}

func (r *RenderObject) BindFragDataLocation(color uint32, name string) {
	gl.BindFragDataLocation(r.shaderProgram, color, gl.Str(name+"\x00"))
}

func (r *RenderObject) Draw() {
	gl.BindVertexArray(r.vao)
	defer gl.BindVertexArray(0) // Unbind it
	gl.DrawArrays(gl.TRIANGLES, 0, r.vboLength)
}
