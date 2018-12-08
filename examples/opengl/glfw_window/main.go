package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tokkenno/seed/renderers/gl"
)

var (
	renderer *gl.Renderer
)

type Window struct {
	glfw.Window
}

func (window *Window) RequestAnimationFrame(cb func()) {
	window.SetRefreshCallback(func(w *glfw.Window) {
		cb()
	})
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(1)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	windowPtr, err := glfw.CreateWindow(800, 600, "Test window", nil, nil)
	if err != nil {
		panic(2)
	}

	window := &Window{*windowPtr}

	window.MakeContextCurrent()

	renderer, err = gl.NewRenderer(window)

	if err != nil {
		panic(3)
	}

	window.RequestAnimationFrame(mainLoop)

	for !window.ShouldClose() {

	}
}

func mainLoop() {
	fmt.Println("Hello frame")
}
