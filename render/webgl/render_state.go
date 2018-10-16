package webgl

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/core/cameras"
	"github.com/tokkenno/seed/render/webgl/gl"
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