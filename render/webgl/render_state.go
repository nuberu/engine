package webgl

import (
	"../../core"
	"../../core/cameras"
)

type RenderState struct {
	lights       *gl.Lights
	lightsArray  []*gl.Lights
	shadowsArray []*gl.Lights
}

func NewRenderState() *RenderState {
	state := new(RenderState)
	return state
}

func (rs *RenderState) AddLight(light *gl.Lights) {
	rs.lightsArray = append(rs.lightsArray, light)
}

func (rs *RenderState) AddShadow(shadow *gl.Lights) {
	rs.shadowsArray = append(rs.shadowsArray, shadow)
}

func (rs *RenderState) SetupLights(camera *cameras.Camera) {
	rs.lights.Setup(rs.lightsArray, rs.shadowsArray, camera)
}

type RenderStates struct {
	states map[uint64]map[uint64]*RenderState
}

func NewRenderStates() *RenderStates {
	states := new(RenderStates)
	states.states = make(map[uint64]map[uint64]*RenderState)
	return states
}

func (rs *RenderStates) Get(scene *core.Scene, camera *cameras.Camera) *RenderState {
	var renderState *RenderState = nil

	sceneState, sceneStateExists := rs.states[scene.GetId()]

	if sceneStateExists {
		state, cameraStateExists := sceneState[camera.GetId()]

		if cameraStateExists {
			return state
		} else {
			renderState = NewRenderState()
			sceneState[camera.GetId()] = renderState
			return renderState
		}
	} else {
		renderState = NewRenderState()
		rs.states[scene.GetId()] = make(map[uint64]*RenderState)
		rs.states[scene.GetId()][camera.GetId()] = renderState
		return renderState
	}
}

func (rs *RenderStates) Dispose() {
	for k := range rs.states {
		delete(rs.states, k)
	}
}
