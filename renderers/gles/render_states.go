package webgl

import (
	"github.com/tokkenno/seed/core"
)

type RenderStates struct {
	states map[uint64]map[uint64]*RenderState
}

func NewRenderStates() *RenderStates {
	states := new(RenderStates)
	states.states = make(map[uint64]map[uint64]*RenderState)
	return states
}

func (rs *RenderStates) Get(scene *core.Scene, camera *core.Camera) *RenderState {
	var renderState *RenderState = nil

	sceneState, sceneStateExists := rs.states[uint64(scene.GetId())]

	if sceneStateExists {
		state, cameraStateExists := sceneState[uint64(camera.GetId())]

		if cameraStateExists {
			return state
		} else {
			renderState = NewRenderState()
			sceneState[uint64(camera.GetId())] = renderState
			return renderState
		}
	} else {
		renderState = NewRenderState()
		rs.states[uint64(scene.GetId())] = make(map[uint64]*RenderState)
		rs.states[uint64(scene.GetId())][uint64(camera.GetId())] = renderState
		return renderState
	}
}

func (rs *RenderStates) Dispose() {
	for k := range rs.states {
		delete(rs.states, k)
	}
}

