package engine

import (
	"fmt"
	gl "github.com/go-gl/gl"
	"io/ioutil"
)

type Shader struct {
	filepath      string
	shaderVersion string
	shaderType    gl.GLenum
}

func NewShader(filepath string, shaderVersion string, shaderType gl.GLenum) gl.Shader {
	shader := new(Shader)
	shader.filepath = filepath
	shader.shaderVersion = shaderVersion
	shader.shaderType = shaderType

	source := shader.read()
	return shader.compile(source)
}

func (s *Shader) read() string {
	if data, err := ioutil.ReadFile(s.filepath + s.shaderVersion + ".glsl"); err != nil {
		fmt.Println("[" + s.filepath + "]Shader Read Error:" + err.Error())
		panic("Shader not found")
	} else {
		return string(data)
	}
}

func (s *Shader) compile(source string) gl.Shader {
	shader := gl.CreateShader(s.shaderType)
	shader.Source(source)
	shader.Compile()
	if info := shader.GetInfoLog(); info != "" {
		fmt.Println(info)
	}
	return shader
}
