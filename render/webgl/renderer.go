package webgl

import (
	"../../core"
	"../../core/cameras"
	"../../render"
	"./dom"
)

type Renderer struct {
	Canvas *dom.Canvas
}

func NewRenderer(options *RendererOptions) {
	renderer := new(Renderer)
	renderer.Canvas = options.Canvas
}

func (renderer *Renderer) Render(scene *core.Scene, camera *cameras.Camera, target *render.Target, forceClear bool) {
	if scene.AutoUpdateRender {
		scene.UpdateMatrixWorld(false)
	}

	if camera.Parent == nil {
		camera.UpdateMatrixWorld(false)
	}

	currentRenderState
}
