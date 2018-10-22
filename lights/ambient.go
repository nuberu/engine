package lights

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/math"
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
