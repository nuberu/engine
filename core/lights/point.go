package lights

import "github.com/tokkenno/seed/core/math"

type Point struct {
	Light
}

func NewPoint(color math.Color, intensity float32) *Point {
	light := Point{
		Light: *NewLight(color, intensity),
	}

	return &light
}
