package render

import (
	"github.com/nuberu/engine/core"
	"github.com/nuberu/engine/math"
)

type Renderer interface {
	Dispose()
	Init() error
	IsInitiated() bool
	NextFrame(func())
	GetPixelRatio() float32
	SetPixelRatio(ratio float32)
	GetSize() (x uint, y uint)
	SetSize(x uint, y uint)
	Render(scene *core.Scene, camera *core.Camera, target *Target, forceClear bool)
	SetScissor(area *math.Box2)
	SetViewport(area *math.Box2)
}
