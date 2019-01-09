package renderer

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/renderers/webgl"
)

type State struct {
	lights       *webgl.Lights
	lightsArray  []*core.Light
	shadowsArray []*core.Light
}

func NewState() *State {
	return &State{
		lightsArray:  make([]*core.Light, 0),
		shadowsArray: make([]*core.Light, 0),
	}
}

func (state *State) Restart() {
	state.lightsArray = make([]*core.Light, 0)
	state.shadowsArray = make([]*core.Light, 0)
}

func (state *State) AddLight(light *core.Light) {
	state.lightsArray = append(state.lightsArray, light)
}

func (state *State) AddShadow(shadow *core.Light) {
	state.shadowsArray = append(state.shadowsArray, shadow)
}

func (state *State) SetupLights(camera *core.Camera) {
	state.lights.Setup(state.lightsArray, state.shadowsArray, camera)
}