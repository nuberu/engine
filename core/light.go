package core

import (
	"github.com/nuberu/engine/math"
)

type Light struct {
	Object3

	color     *math.Color
	intensity float32
}

func NewLight(color math.Color, intensity float32) *Light {
	light := new(Light)

	light.color = &color
	light.intensity = intensity

	return light
}

func (light *Light) GetColor() *math.Color {
	return light.color.Clone()
}

func (light *Light) GetIntensity() float32 {
	return light.intensity
}
