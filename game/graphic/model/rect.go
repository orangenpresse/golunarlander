package model

import (
	gl "github.com/go-gl/gl"
	"github.com/orangenpresse/golunarlander/game/graphic"
)

type Rect struct {
	graphic.GraphicsObject
}

func (r Rect) LoadToVram(shaderProgram gl.Program) *graphic.RenderObject {
	verticies := []float32{
		-1, -1, 0,
		1, -1, 0,
		-1, 1, 0,
		-1, 1, 0,
		1, 1, 0,
		1, -1, 0}
	return graphic.NewRenderObject(verticies, shaderProgram)
}

func NewRect(renderObject *graphic.RenderObject) *Rect {
	rect := new(Rect)
	rect.InitGraphic(renderObject)
	return rect
}
