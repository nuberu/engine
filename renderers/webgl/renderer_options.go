package webgl

import (
	GlES "github.com/tokkenno/seed/renderers/gles"
	"github.com/tokkenno/seed/renderers/webgl/dom"
)

type RendererOptions struct {
	GlES.RendererOptions
	Canvas                *dom.Canvas
}

func DefaultOptions() *RendererOptions {
	options := GlES.NewDefaultOptions()
	options.Canvas = dom.NewCanvas()
	return options
}
