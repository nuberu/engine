package webgl

type RendererOptions struct {
	Alpha                 bool
	Depth                 bool
	Stencil               bool
	Antialias             bool
	PremultipliedAlpha    bool
	PreserveDrawingBuffer bool
	PowerPreference       string
}

func NewDefaultOptions() *RendererOptions {
	options := new(RendererOptions)
	options.Alpha = false
	options.Depth = true
	options.Stencil = true
	options.Antialias = false
	options.PremultipliedAlpha = true
	options.PreserveDrawingBuffer = false
	options.PowerPreference = "default"
	return options
}
