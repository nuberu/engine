package lights

import (
	"github.com/nuberu/engine/core"
	"github.com/nuberu/engine/math"
)

type Hemisphere struct {
	core.Light
}

func NewHemisphere(color math.Color, intensity float32) *Hemisphere {
	light := Hemisphere{
		Light: *core.NewLight(color, intensity),
	}

	return &light
}
