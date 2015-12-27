package engine

import (
	"fmt"
	gl "github.com/go-gl/gl/v3.3-core/gl"
	"io/ioutil"
	"strings"
)

type Shader struct {
	filepath      string
	shaderVersion string
	shaderType    uint32
}

func NewShader(filepath string, shaderVersion string, shaderType uint32) (uint32, error) {
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

func (s *Shader) compile(source string) (uint32, error) {
	shader := gl.CreateShader(s.shaderType)
	csource := gl.Str(source + "\x00")
	gl.ShaderSource(shader, 1, &csource, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
