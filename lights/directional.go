package lights

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/math"
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
