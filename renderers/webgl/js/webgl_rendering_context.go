package js

import (
	coreConstant "github.com/tokkenno/seed/core/constant"
	"github.com/tokkenno/seed/renderers/webgl/constant"
	"syscall/js"
)

// WebGL context wrapper
type WebGLRenderingContext struct {
	loaded bool
	Js     js.Value
}

func (c *WebGLRenderingContext) mapConstant(constant string) js.Value {
	return c.Js.Get(constant)
}

func (c *WebGLRenderingContext) mapCondition(condition coreConstant.Condition) js.Value {
	switch condition {
	case coreConstant.Never:
		return c.Js.Get("NEVER")
	case coreConstant.Always:
		return c.Js.Get("ALWAYS")
	case coreConstant.Less:
		return c.Js.Get("LESS")
	case coreConstant.Equal:
		return c.Js.Get("EQUAL")
	case coreConstant.GreaterEqual:
		return c.Js.Get("GEQUAL")
	case coreConstant.Greater:
		return c.Js.Get("GREATER")
	case coreConstant.NotEqual:
		return c.Js.Get("NOTEQUAL")
	case coreConstant.LessEqual:
		return c.Js.Get("LEQUAL")
	}
	return c.Js.Get("NEVER")
}

func (c *WebGLRenderingContext) ClearColor(r float32, g float32, b float32, a float32) {
	c.Js.Call("clearColor", r, g, b, a)
}

func (c *WebGLRenderingContext) ClearDepth(depth float32) {
	c.Js.Call("clearDepth", depth)
}

func (c *WebGLRenderingContext) ClearStencil(index int) {
	c.Js.Call("clearStencil", index)
}

func (c *WebGLRenderingContext) ColorMask(r float32, g float32, b float32, a float32) {
	c.Js.Call("colorMask", r, g, b, a)
}

func (c *WebGLRenderingContext) DepthFunc(depth coreConstant.Condition) {
	c.Js.Call("depthFunc", c.mapCondition(depth))
}

func (c *WebGLRenderingContext) DepthMask(flag bool) {
	c.Js.Call("depthMask", flag)
}

func (c *WebGLRenderingContext) StencilMask(mask uint) {
	c.Js.Call("stencilMask", mask)
}

func (c *WebGLRenderingContext) StencilFunc(function coreConstant.Condition, ref int, mask uint) {
	c.Js.Call("stencilFunc", c.mapCondition(function), ref, mask)
}

func (c *WebGLRenderingContext) StencilOp(fail constant.StencilFunc, zfail constant.StencilFunc, zpass constant.StencilFunc) {
	c.Js.Call("stencilOp", c.mapConstant(string(fail)), c.mapConstant(string(zfail)), c.mapConstant(string(zpass)))
}
