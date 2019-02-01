package lights

import (
	"github.com/nuberu/engine/core"
	"github.com/nuberu/engine/math"
)

type Spot struct {
	core.Light
}

func NewSpot(color math.Color, intensity float32) *Spot {
	light := Spot{
		Light: *core.NewLight(color, intensity),
	}

	return &light
}
