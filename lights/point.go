package lights

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/math"
)

type Point struct {
	core.Light
}

func NewPoint(color math.Color, intensity float32) *Point {
	light := Point{
		Light: *core.NewLight(color, intensity),
	}

	return &light
}
