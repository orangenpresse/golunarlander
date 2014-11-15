#version 130

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 position;

void main() {
    gl_Position = projection * camera * model * vec4(position, 1.0);
}