package webgl

import (
	"github.com/tokkenno/seed/core/event"
	"github.com/tokkenno/seed/core/math"
	"github.com/tokkenno/seed/core/textures"
)

// TODO: Meter como texture options y componer las ultimas propiedades
type RenderTargetOptions struct {
	textures.Options

	GenerateMipmaps bool
	DepthBuffer     bool
	DepthTexture    *textures.Texture // type?
	StencilBuffer   bool
}

func DefaultRenderTargetOptions() *RenderTargetOptions {
	return &RenderTargetOptions{
		Options: textures.Options{
			MinFilter: textures.LinearMinFilter,
			MagFilter: textures.LinearMagFilter,
		},
		GenerateMipmaps: true,
	}
}

type RenderTarget struct {
	width         uint
	height        uint
	scissor       *math.Vector4
	scissorTest   bool
	viewport      *math.Vector4
	texture       *textures.Texture
	depthBuffer   bool
	depthTexture  *textures.Texture // type?
	stencilBuffer bool

	disposeEvent *event.Emitter
}

func NewRenderTarget(width uint, height uint, options *RenderTargetOptions) *RenderTarget {
	if options == nil {
		options = DefaultRenderTargetOptions()
	}

	rt := RenderTarget{
		width:       width,
		height:      height,
		scissor:     math.NewVector4(0, 0, float64(width), float64(height)),
		scissorTest: false,
		viewport:    math.NewVector4(0, 0, float64(width), float64(height)),
		texture:     textures.NewTexture(nil, 0, options.WrapS, options.WrapT, options.MagFilter, options.MinFilter, options.Format, options.TextureType, options.Anisotropy, options.Encoding),
	}

	rt.texture.GenerateMipmaps = options.GenerateMipmaps

	rt.stencilBuffer = options.StencilBuffer
	rt.depthBuffer = options.DepthBuffer
	rt.depthTexture = options.DepthTexture

	rt.disposeEvent = new(event.Emitter)

	return &rt
}

func (rt *RenderTarget) SetSize(width, height uint) {
	if rt.width != width || rt.height != height {
		rt.width = width
		rt.height = height
		rt.Dispose()
	}

	rt.viewport.Set(0, 0, float64(width), float64(height))
	rt.scissor.Set(0, 0, float64(width), float64(height))
}

func (rt *RenderTarget) Clone() *RenderTarget {
	return &RenderTarget{
		width:        rt.width,
		height:       rt.height,
		viewport:     rt.viewport.Clone(),
		texture:      rt.texture.Clone(),
		depthBuffer:  rt.depthBuffer,
		depthTexture: rt.depthTexture,
	}
}

func (rt *RenderTarget) Dispose() {
	rt.disposeEvent.Emit(nil)
}

func (rt *RenderTarget) OnDispose() *event.Handler {
	return rt.disposeEvent.GetHandler()
}
