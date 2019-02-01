package lights

import (
	"github.com/nuberu/engine/core"
	"github.com/nuberu/engine/math"
)

type RectArea struct {
	core.Light
}

func NewRectArea(color math.Color, intensity float32) *RectArea {
	light := RectArea{
		Light: *core.NewLight(color, intensity),
	}

	return &light
}
