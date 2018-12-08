package gl

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/renderers/gl_common"
)

type RenderState struct {
	lights       *gl_common.Lights
	lightsArray  []*core.Light
	shadowsArray []*core.Light
}

func NewRenderState() *RenderState {
	state := new(RenderState)
	return state
}

func (rs *RenderState) AddLight(light *core.Light) {
	rs.lightsArray = append(rs.lightsArray, light)
}

func (rs *RenderState) AddShadow(shadow *core.Light) {
	rs.shadowsArray = append(rs.shadowsArray, shadow)
}

func (rs *RenderState) SetupLights(camera *core.Camera) {
	rs.lights.Setup(rs.lightsArray, rs.shadowsArray, camera)
}