package math

import (
	nativeMath "math"
)

const (
	Pi = float32(nativeMath.Pi)
)

func Abs(v float32) float32 {
	return float32(nativeMath.Abs(float64(v)))
}

func Sqrt(v float32) float32 {
	return float32(nativeMath.Sqrt(float64(v)))
}

func Cos(x float32) float32 {
	return float32(nativeMath.Cos(float64(x)))
}

func Sin(x float32) float32 {
	return float32(nativeMath.Sin(float64(x)))
}

func Asin(x float32) float32 {
	return float32(nativeMath.Asin(float64(x)))
}

func Acos(x float32) float32 {
	return float32(nativeMath.Acos(float64(x)))
}

func Atan2(y, x float32) float32 {
	return float32(nativeMath.Atan2(float64(y), float64(x)))
}

func Max(y, x float32) float32 {
	return float32(nativeMath.Max(float64(y), float64(x)))
}

func Min(y, x float32) float32 {
	return float32(nativeMath.Min(float64(y), float64(x)))
}

func Round(x float32) float32 {
	return float32(nativeMath.Round(float64(x)))
}

func Pow(x, y float32) float32 {
	return float32(nativeMath.Pow(float64(x), float64(y)))
}

func Floor(x float32) float32 {
	return float32(nativeMath.Floor(float64(x)))
}

func Ceil(x float32) float32 {
	return float32(nativeMath.Ceil(float64(x)))
}

func Clamp(value float32, min float32, max float32) float32 {
	return Max(min, Min(max, value))
}
