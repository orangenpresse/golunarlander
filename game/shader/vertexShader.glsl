#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec2 position;

void main() {
    gl_Position = projection * camera * model * vec4(position, 0.0, 1.0);
}