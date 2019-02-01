package buffer

import (
	coreConstant "github.com/nuberu/engine/core/constant"
	"github.com/nuberu/engine/renderers/webgl/constant"
	"github.com/nuberu/engine/renderers/webgl/js"
	"github.com/nuberu/webgl"
	"github.com/nuberu/webgl/types"
)

type Stencil struct {
	glContext               *webgl.RenderingContext
	locked                  bool
	currentStencilMaskInit  bool
	currentStencilMask      uint32
	currentStencilFuncInit  bool
	currentStencilFunc      coreConstant.Condition
	currentStencilRef       int
	currentStencilFuncMask  uint32
	currentStencilOpInit    bool
	currentStencilFail      constant.StencilFunc
	currentStencilZFail     constant.StencilFunc
	currentStencilZPass     constant.StencilFunc
	currentStencilClearInit bool
	currentStencilClear     int
}

func NewStencilBuffer(glContext *webgl.RenderingContext) *Stencil {
	return &Stencil{
		glContext:               glContext,
		locked:                  false,
		currentStencilMaskInit:  false,
		currentStencilFuncInit:  false,
		currentStencilOpInit:    false,
		currentStencilClearInit: false,
	}
}

func (ste *Stencil) SetMask(stencilMask uint32) {
	if (!ste.currentStencilMaskInit || ste.currentStencilMask != stencilMask) && !ste.locked {
		ste.currentStencilMaskInit = true
		ste.glContext.StencilMask(stencilMask)
		ste.currentStencilMask = stencilMask
	}
}

func (ste *Stencil) SetFunc(stencilFunc coreConstant.Condition, stencilRef int, stencilMask uint32) {
	if !ste.currentStencilFuncInit || ste.currentStencilFunc != stencilFunc || ste.currentStencilRef != stencilRef || ste.currentStencilFuncMask != stencilMask {
		ste.currentStencilFuncInit = true
		ste.glContext.StencilFunc(types.GLEnum(uint32(stencilFunc) + uint32(webgl.NEVER)), stencilRef, stencilMask)
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
