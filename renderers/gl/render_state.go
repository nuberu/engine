package gl

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/renderers/gl_common"
)

type RenderState struct {
	lights       *gl_common.Lights
	lightsArray  []*gl_common.Lights
	shadowsArray []*gl_common.Lights
}

func NewRenderState() *RenderState {
	state := new(RenderState)
	return state
}

func (rs *RenderState) AddLight(light *gl_common.Lights) {
	rs.lightsArray = append(rs.lightsArray, light)
}

func (rs *RenderState) AddShadow(shadow *gl_common.Lights) {
	rs.shadowsArray = append(rs.shadowsArray, shadow)
}

func (rs *RenderState) SetupLights(camera *core.Camera) {
	rs.lights.Setup(rs.lightsArray, rs.shadowsArray, camera)
}