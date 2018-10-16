package lights

import "github.com/tokkenno/seed/core/math"

type Ambient struct {
	Light
}

func NewAmbient(color math.Color, intensity float32) *Ambient {
	light := Ambient{
		Light: *NewLight(color, intensity),
	}

	return &light
}
