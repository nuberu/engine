package buffer

import (
	"github.com/nuberu/engine/math"
	"github.com/nuberu/webgl"
)

type Color struct {
	glContext         *webgl.RenderingContext
	locked            bool
	color             *math.Vector4
	currentColorMask  *math.Vector4
	currentColorClear *math.Vector4
}

func NewColorBuffer(glContext *webgl.RenderingContext) *Color {
	return &Color{
		glContext:         glContext,
		locked:            false,
		color:             math.NewVector4(0, 0, 0, 1),
		currentColorMask:  nil,
		currentColorClear: math.NewVector4(0, 0, 0, 0),
	}
}

func (col *Color) SetMask(colorMask *math.Vector4) {
	if col.currentColorMask != colorMask && !col.locked {
		col.glContext.ColorMask(colorMask.X, colorMask.Y, colorMask.Z, colorMask.W)
		col.currentColorMask = colorMask
	}
}

func (col *Color) SetLocked(lock bool) {
	col.locked = lock
}

func (col *Color) SetClear(r float32, g float32, b float32, a float32, premultipliedAlpha bool) {
	if premultipliedAlpha {
		r *= a
		g *= a
		b *= a
	}

	col.color.Set(r, g, b, a)

	if !col.currentColorClear.Equals(col.color) {
		col.glContext.ClearColor(r, g, b, a)
		col.currentColorClear.Copy(col.color)
	}
}

func (col *Color) Reset() {
	col.locked = false
	col.currentColorMask = nil
	col.currentColorClear.Set(-1, 0, 0, 0)
}
