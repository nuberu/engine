package math

import "math"

func Clamp(value float32, min float32, max float32) float32 {
	return math.Max(min, math.Min(max, value))
}
