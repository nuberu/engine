package main

import (
	"github.com/tokkenno/seed/math"
	"github.com/tokkenno/seed/renderers/gl"
)

var renderer *gl.Renderer

func main() {
	target, err := gl.NewWindow(math.NewVector2(800, 600), "Test window")

	if err != nil {
		panic(1)
	}

	renderer, err = gl.NewRenderer(target)

	if err != nil {
		panic(2)
	}

	renderer.NextFrame(mainLoop)
}

func mainLoop() {
	renderer.NextFrame(mainLoop)
}
