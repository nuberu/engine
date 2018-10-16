package gl

import (
	"image/color"
	"reflect"

	"github.com/tokkenno/seed/core/cameras"
	"github.com/tokkenno/seed/core/lights"
	"github.com/tokkenno/seed/core/math"
)

type Uniform struct {
	position         *math.Vector3
	direction        *math.Vector3
	color            math.Color
	skyColor         math.Color
	groundColor      math.Color
	distance         float64
	coneCos          float64
	penumbraCos      float64
	decay            float64
	shadow           bool
	shadowBias       uint
	shadowRadius     uint
	shadowMapSize    *math.Vector2
	shadowCameraNear float64
	shadowCameraFar  float64
	halfWidth        *math.Vector3
	halfHeight       *math.Vector3
}

type UniformsCache struct {
	lights map[uint64]*Uniform
}

func (cache *UniformsCache) Get(light *lights.Light) *Uniform {
	lightInstance, lightExists := cache.lights[light.GetId()]

	if lightExists {
		return lightInstance
	} else {
		switch reflect.TypeOf(light) {
		case reflect.TypeOf(&lights.Directional{}):
			return &Uniform{
				direction:     math.NewVector3(0, 0, 0),
				color:         color.White,
				shadow:        false,
				shadowBias:    0,
				shadowRadius:  0,
				shadowMapSize: math.NewVector2(0, 0),
			}
		case reflect.TypeOf(&lights.Spot{}):
			return &Uniform{
				position:      math.NewVector3(0, 0, 0),
				direction:     math.NewVector3(0, 0, 0),
				color:         color.White,
				distance:      0,
				coneCos:       0,
				penumbraCos:   0,
				decay:         0,
				shadow:        false,
				shadowBias:    0,
				shadowRadius:  1,
				shadowMapSize: math.NewVector2(0, 0),
			}
		case reflect.TypeOf(&lights.Point{}):
			return &Uniform{
				position:         math.NewVector3(0, 0, 0),
				color:            color.White,
				distance:         0,
				decay:            0,
				shadow:           false,
				shadowBias:       0,
				shadowRadius:     1,
				shadowMapSize:    math.NewVector2(0, 0),
				shadowCameraNear: 1,
				shadowCameraFar:  1000,
			}
		case reflect.TypeOf(&lights.Hemisphere{}):
			return &Uniform{
				direction:   math.NewVector3(0, 0, 0),
				skyColor:    color.White,
				groundColor: color.White,
			}
		case reflect.TypeOf(&lights.RectArea{}):
			return &Uniform{
				color:      color.White,
				position:   math.NewVector3(0, 0, 0),
				halfWidth:  math.NewVector3(0, 0, 0),
				halfHeight: math.NewVector3(0, 0, 0),
			}
		case reflect.TypeOf(&lights.Light{}):
		case reflect.TypeOf(&lights.Ambient{}):
		default:
			break
		}
	}

	return nil
}

type LightState struct {
	id      uint64
	ambient *math.Vector3
}

type Lights struct {
	cache *UniformsCache
	state LightState
}

func (l *Lights) Setup(lightList []*lights.Light, shadowList []*lights.Light, camera *cameras.Camera) {
	for _, light := range lightList {
		lightType := reflect.TypeOf(light)

		if lightType == reflect.TypeOf(&lights.Ambient{}) {
			l.state.ambient.Add(light.GetColor().Clone().SetIntensity(light.GetIntensity()).GetRGBVector())
			// TODO: https://github.com/mrdoob/three.js/blob/dev/src/renderers/webgl/WebGLLights.js#L166
		}
	}
}
