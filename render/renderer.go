package render

import (
	"github.com/tokkenno/seed/core"
)

type Renderer interface {
	Render(scene *core.Scene, camera *core.Camera, target *Target, forceClear bool)
}
