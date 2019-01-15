package buffer

import (
	coreConstant "github.com/tokkenno/seed/core/constant"
	"github.com/tokkenno/seed/renderers/webgl/constant"
	"github.com/tokkenno/seed/renderers/webgl/js"
)

type Stencil struct {
	glContext               *js.WebGLRenderingContext
	locked                  bool
	currentStencilMaskInit  bool
	currentStencilMask      uint
	currentStencilFuncInit  bool
	currentStencilFunc      coreConstant.Condition
	currentStencilRef       int
	currentStencilFuncMask  uint
	currentStencilOpInit    bool
	currentStencilFail      constant.StencilFunc
	currentStencilZFail     constant.StencilFunc
	currentStencilZPass     constant.StencilFunc
	currentStencilClearInit bool
	currentStencilClear     int
}

func NewStencilBuffer(glContext *js.WebGLRenderingContext) *Stencil {
	return &Stencil{
		glContext:               glContext,
		locked:                  false,
		currentStencilMaskInit:  false,
		currentStencilFuncInit:  false,
		currentStencilOpInit:    false,
		currentStencilClearInit: false,
	}
}

func (ste *Stencil) SetMask(stencilMask uint) {
	if (!ste.currentStencilMaskInit || ste.currentStencilMask != stencilMask) && !ste.locked {
		ste.currentStencilMaskInit = true
		ste.glContext.StencilMask(stencilMask)
		ste.currentStencilMask = stencilMask
	}
}

func (ste *Stencil) SetFunc(stencilFunc coreConstant.Condition, stencilRef int, stencilMask uint) {
	if !ste.currentStencilFuncInit || ste.currentStencilFunc != stencilFunc || ste.currentStencilRef != stencilRef || ste.currentStencilFuncMask != stencilMask {
		ste.currentStencilFuncInit = true
		ste.glContext.StencilFunc(stencilFunc, stencilRef, stencilMask)
		ste.currentStencilFunc = stencilFunc
		ste.currentStencilRef = stencilRef
		ste.currentStencilFuncMask = stencilMask
	}
}

func (ste *Stencil) SetOp(stencilFail constant.StencilFunc, stencilZFail constant.StencilFunc, stencilZPass constant.StencilFunc) {
	if !ste.currentStencilOpInit || ste.currentStencilFail != stencilFail || ste.currentStencilZFail != stencilZFail || ste.currentStencilZPass != stencilZPass {
		ste.currentStencilOpInit = true
		ste.glContext.StencilOp(stencilFail, stencilZFail, stencilZPass)
		ste.currentStencilFail = stencilFail
		ste.currentStencilZFail = stencilZFail
		ste.currentStencilZPass = stencilZPass
	}
}

func (ste *Stencil) SetClear(stencil int) {
	if !ste.currentStencilClearInit || ste.currentStencilClear != stencil {
		ste.currentStencilClearInit = true
		ste.glContext.ClearStencil(stencil)
		ste.currentStencilClear = stencil
	}
}

func (ste *Stencil) SetLocked(lock bool) {
	ste.locked = lock
}

func (ste *Stencil) Reset() {
	ste.locked = false
	ste.currentStencilMaskInit = false
	ste.currentStencilFuncInit = false
	ste.currentStencilOpInit = false
	ste.currentStencilClearInit = false
}
