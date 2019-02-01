package main

import (
	"github.com/nuberu/engine/core"
	"github.com/nuberu/engine/geometries"
	"github.com/nuberu/engine/materials"
	"github.com/nuberu/engine/objects"
	"github.com/nuberu/engine/renderers/webgl/renderer"
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")
	container := document.Call("getElementById", "container")

	renderer, _ := renderer.NewRenderer(nil)
	container.Call("appendChild", renderer.DomElement)

	var renderFrame js.Callback

	var scene *core.Scene
	var camera *core.Camera

	cube := objects.NewMesh(geometries.NewBox(1, 1, 1), materials.NewMeshBasic())

	renderFrame = js.NewCallback(func(args []js.Value) {
		cube.Rotation.SetX(cube.Rotation.GetX() + 0.01)
		cube.Rotation.SetY(cube.Rotation.GetY() + 0.01)

		renderer.Render(scene, camera, nil, true)

		js.Global().Call("requestAnimationFrame", renderFrame)
	})

	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	done := make(chan struct{}, 0)
	<-done
}
