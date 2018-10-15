package math

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(x float64, y float64) *Vector2 {
	return &Vector2{
		X: x,
		Y: y,
	}
}
