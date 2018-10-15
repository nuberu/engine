package lights

import "../math"

type Hemisphere struct {
	Light
}

func NewHemisphere(color math.Color, intensity float32) *Hemisphere {
	light := Hemisphere{
		Light: *NewLight(color, intensity),
	}

	return &light
}
