package engine

import (
	gl "github.com/go-gl/gl"
)

type ModelManager struct {
	models        map[string]*RenderObject
	shaderProgram *gl.Program
}

func NewModelManager() *ModelManager {
	modelManger := new(ModelManager)
	modelManger.models = make(map[string]*RenderObject)
	return modelManger
}

func (m *ModelManager) RegisterModel(name string, renderObject *RenderObject) {
	m.models[name] = renderObject
}

func (m *ModelManager) GetRenderObject(name string) *RenderObject {
	return m.models[name]
}
