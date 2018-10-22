package lights

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/math"
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
