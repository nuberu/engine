package lights

import (
	"../math"
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
