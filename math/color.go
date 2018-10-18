package math

import "image/color"

type Color struct {
	color.RGBA
}

func (col *Color) Clone() *Color {
	newCol := *col
	return &newCol
}

func (col *Color) SetIntensity(intensity float32) *Color {
	col.R = uint8(float32(col.R) * intensity)
	col.G = uint8(float32(col.G) * intensity)
	col.B = uint8(float32(col.B) * intensity)
	return col
}

func (col *Color) GetRGBVector() *Vector3 {
	return NewVector3(float64(col.R), float64(col.G), float64(col.B))
}
