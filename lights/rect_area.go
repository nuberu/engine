package lights

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/math"
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
