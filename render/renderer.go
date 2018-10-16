package render

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/core/cameras"
)

type Renderer interface {
	Render(scene *core.Scene, camera *cameras.Camera, target *Target, forceClear bool)
}
