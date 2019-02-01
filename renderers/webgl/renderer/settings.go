package renderer

import "github.com/nuberu/engine/renderers/webgl/constant"

type Settings struct {
	Alpha                  bool
	Depth                  bool
	Stencil                bool
	Antialias              bool
	PremultipliedAlpha     bool
	PreserveDrawingBuffer  bool
	PowerPreference        string
	Precision              constant.Precision
	LogarithmicDepthBuffer bool
}

func DefaultSettings() *Settings {
	settings := new(Settings)
	settings.Alpha = false
	settings.Depth = true
	settings.Stencil = true
	settings.Antialias = false
	settings.PremultipliedAlpha = true
	settings.PreserveDrawingBuffer = false
	settings.PowerPreference = "default"
	settings.Precision = constant.HighPrecision
	settings.LogarithmicDepthBuffer = false
	return settings
}
