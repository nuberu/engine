package lights

import (
	"github.com/nuberu/engine/core"
	"github.com/nuberu/engine/math"
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
