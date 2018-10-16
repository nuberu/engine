package lights

import (
	"github.com/tokkenno/seed/core/math"
)

type RectArea struct {
	Light
}

func NewRectArea(color math.Color, intensity float32) *RectArea {
	light := RectArea{
		Light: *NewLight(color, intensity),
	}

	return &light
}
