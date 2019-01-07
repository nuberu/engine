package materials

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/core/constant"
	"github.com/tokkenno/seed/math"
)

type MeshBasic struct {
	core.Material

	color math.Color
	// lightMap
	lightMapIntensity float32
	// aoMap
	aoMapIntensity float32
	// specularMap
	// alphaMap
	// envMap

	combine         constant.Operation
	reflectivity    float32
	refractionRatio float32

	wireframe          bool
	wireframeLinewidth float32
	wireframeLinecap   string
	wireframeLinejoin  string

	skinning     bool
	morphTargets bool
}

func NewMeshBasic() *MeshBasic {
	mat := &MeshBasic{
		Material:           *core.NewMaterial(),
		color:              math.Color{R: 0xff, G: 0xff, B: 0xff},
		lightMapIntensity:  1.0,
		combine:            constant.MultiplyOperation,
		reflectivity:       1,
		refractionRatio:    0.98,
		wireframe:          false,
		wireframeLinewidth: 1,
		wireframeLinecap:   "round",
		wireframeLinejoin:  "round",
		skinning:           false,
		morphTargets:       false,
	}
	mat.Lights = false
	return mat
}
