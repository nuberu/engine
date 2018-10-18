package math

import "math"

func Clamp(value float64, min float64, max float64) float64 {
	return math.Max(min, math.Min(max, value))
}
