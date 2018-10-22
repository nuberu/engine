package lights

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/math"
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
