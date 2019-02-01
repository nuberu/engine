package lights

import (
	"github.com/nuberu/engine/core"
	"github.com/nuberu/engine/math"
)

type Directional struct {
	core.Light
}

func NewDirectional(color math.Color, intensity float32) *Directional {
	light := Directional{
		Light: *core.NewLight(color, intensity),
	}

	return &light
}
