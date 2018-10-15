package render

import (
	"../core"
	"../core/cameras"
)

type Renderer interface {
	Render(scene *core.Scene, camera *cameras.Camera, target *Target, forceClear bool)
}
