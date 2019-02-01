package buffer

import (
	"github.com/nuberu/engine/core/constant"
	"github.com/nuberu/webgl"
	"github.com/nuberu/webgl/types"
)

type Depth struct {
	glContext             *webgl.RenderingContext
	locked                bool
	currentDepthMaskInit  bool
	currentDepthMask      bool
	currentDepthFuncInit  bool
	currentDepthFunc      constant.Condition
	currentDepthClearInit bool
	currentDepthClear     float32
}

func NewDepthBuffer(glContext *webgl.RenderingContext) *Depth {
	return &Depth{
		glContext:             glContext,
		locked:                false,
		currentDepthMaskInit:  false,
		currentDepthFuncInit:  false,
		currentDepthClearInit: false,
	}
}

func (dep *Depth) SetMask(depthMask bool) {
	if (!dep.currentDepthMaskInit || dep.currentDepthMask != depthMask) && !dep.locked {
		dep.currentDepthMaskInit = true
		dep.glContext.DepthMask(depthMask)
		dep.currentDepthMask = depthMask
	}
}

func (dep *Depth) SetFunc(depthFunc constant.Condition) {
	if !dep.currentDepthFuncInit || dep.currentDepthFunc != depthFunc {
		dep.currentDepthFuncInit = true
		dep.glContext.DepthFunc(types.GLEnum(uint32(depthFunc) + uint32(webgl.NEVER)))
	}
}

func (dep *Depth) SetClear(depth float32) {
	if !dep.currentDepthClearInit || dep.currentDepthClear != depth {
		dep.currentDepthClearInit = true
		dep.glContext.ClearDepth(depth)
		dep.currentDepthClear = depth
	}
}

func (dep *Depth) SetLocked(lock bool) {
	dep.locked = lock
}

func (dep *Depth) Reset() {
	dep.locked = false
	dep.currentDepthMaskInit = false
	dep.currentDepthFuncInit = false
	dep.currentDepthClearInit = false
}
