package lights

import "github.com/tokkenno/seed/core/math"

type Spot struct {
	Light
}

func NewSpot(color math.Color, intensity float32) *Spot {
	light := Spot{
		Light: *NewLight(color, intensity),
	}

	return &light
}
