package webgl

import (
	"github.com/tokkenno/seed/render/webgl/dom"
)

type RendererOptions struct {
	Canvas                *dom.Canvas
	Alpha                 bool
	Depth                 bool
	Stencil               bool
	Antialias             bool
	PremultipliedAlpha    bool
	PreserveDrawingBuffer bool
	PowerPreference       string
}

func DefaultOptions() *RendererOptions {
	options := new(RendererOptions)
	options.Canvas = dom.NewCanvas()
	options.Alpha = false
	options.Depth = true
	options.Stencil = true
	options.Antialias = false
	options.PremultipliedAlpha = true
	options.PreserveDrawingBuffer = false
	options.PowerPreference = "default"
	return options
}
