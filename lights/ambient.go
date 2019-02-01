package lights

import (
	"github.com/nuberu/engine/core"
	"github.com/nuberu/engine/math"
)

type Ambient struct {
	core.Light
}

func NewAmbient(color math.Color, intensity float32) *Ambient {
	light := Ambient{
		Light: *core.NewLight(color, intensity),
	}

	return &light
}
