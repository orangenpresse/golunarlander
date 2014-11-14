package graphic

import (
	gl "github.com/go-gl/gl"
)

type ModelManager struct {
	models        []*RenderObject
	shaderProgram *gl.Program
}

func NewModelManager() *ModelManager {
	modelManger := new(ModelManager)
	modelManger.models = make([]*RenderObject, 0)
	return modelManger
}

func (m *ModelManager) RegisterModel(verticies *[]float32) *RenderObject {
	renderObject := new(RenderObject)
	m.models = append(m.models, renderObject)
	return renderObject
}
