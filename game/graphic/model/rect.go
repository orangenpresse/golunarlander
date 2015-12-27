package model

import (
	"github.com/orangenpresse/golunarlander/game/graphic/engine"
)

type Rect struct {
	engine.GraphicsObject
}

func (r Rect) LoadToVram(shaderProgram uint32) *engine.RenderObject {
	verticies := []float32{
		-1, -1, 0,
		1, -1, 0,
		-1, 1, 0,
		-1, 1, 0,
		1, 1, 0,
		1, -1, 0}
	return engine.NewRenderObject(verticies, shaderProgram)
}

func NewRect(renderObject *engine.RenderObject) *Rect {
	rect := new(Rect)
	rect.InitGraphic(renderObject)
	return rect
}
