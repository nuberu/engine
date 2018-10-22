package math

import "image/color"

var (
	ColorBlack = Color{R: 0, G: 0, B: 0, A: 0xffffffff}
)

type Color struct {
	color.Color

	R uint16
	G uint16
	B uint16
	A uint16
}

func (col *Color) RGBA() (r, g, b, a uint32) {
	return uint32(col.R), uint32(col.G), uint32(col.B), uint32(col.A)
}

func (col *Color) Clone() *Color {
	newCol := *col
	return &newCol
}

func (col *Color) SetIntensity(intensity float32) *Color {
	col.R = uint16(float32(col.R) * intensity)
	col.G = uint16(float32(col.G) * intensity)
	col.B = uint16(float32(col.B) * intensity)
	return col
}

func (col *Color) GetRGBVector() *Vector3 {
	return NewVector3(float32(col.R), float32(col.G), float32(col.B))
}