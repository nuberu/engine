package main

import (
	"github.com/tokkenno/seed/math"
	"github.com/tokkenno/seed/renderers/gl"
)

func main() {
	target, err := gl.NewWindow(math.NewVector2(800, 600), "Test window")

	if err != nil {
		panic(1)
	}

	renderer := gl.NewRenderer(target)

	renderer.NextFrame(mainLoop)
}

func mainLoop() {

}