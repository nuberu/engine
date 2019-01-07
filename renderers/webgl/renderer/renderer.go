package renderer

import (
	"errors"
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/core/constant"
	"github.com/tokkenno/seed/core/render"
	"github.com/tokkenno/seed/math"
	"github.com/tokkenno/seed/renderers/webgl"
	"syscall/js"
)

type Renderer struct {
	DomElement js.Value

	AutoClear        bool
	AutoClearColor   bool
	AutoClearDepth   bool
	AutoClearStencil bool

	SortObjects bool

	//ClippingPlanes       []
	LocalClippingEnabled bool

	GammaFactor float32
	GammaInput  bool
	GammaOutput bool

	PhysicallyCorrentLights bool
	ToneMapping             constant.ToneMapping
	ToneMappingExposure     float32
	ToneMappingWhitePoint   float32

	MaxMorphTargets int
	MaxMorphNormals int

	isContextLost bool

	//frameBuffer

	//currentRenderTarget
	//currentFrameBuffer
	currentMaterialId int64

	//currentGeometryProgram *GeometryProgram

	currentCamera *core.Camera
	//currentArrayCamera

	currentViewport    *math.Vector4
	currentScissor     *math.Vector4
	currentScissorTest bool

	usedTextureUnits int

	width  int
	height int

	pixelRatio float32

	viewPort    *math.Vector4
	scissor     *math.Vector4
	scissorTest bool

	frustum *math.Frustum
	//clipping             *gl.Clipping
	clippingEnabled      bool
	localClippingEnabled bool

	projectScreenMatrix *math.Matrix4

	vector3 *math.Vector3

	contextAttributes *Settings

	extensions *webgl.Extensions

	glContext          js.Value
	onContextLostJs    js.Callback
	onContextRestoreJs js.Callback
}

func NewRenderer(options *Settings) (*Renderer, error) {
	if options == nil {
		options = DefaultSettings()
	}

	canvas := js.Global().Get("document").Call("createElement", "canvas")

	width := canvas.Get("width").Int()
	height := canvas.Get("height").Int()

	renderer := Renderer{
		AutoClear:               true,
		AutoClearColor:          true,
		AutoClearDepth:          true,
		AutoClearStencil:        true,
		SortObjects:             true,
		localClippingEnabled:    true,
		GammaFactor:             2.0,
		GammaInput:              false,
		GammaOutput:             false,
		PhysicallyCorrentLights: false,
		ToneMapping:             constant.LinearToneMapping,
		ToneMappingExposure:     1.0,
		ToneMappingWhitePoint:   1.0,
		MaxMorphTargets:         8,
		MaxMorphNormals:         4,
		DomElement:              canvas,

		isContextLost:        false,
		currentMaterialId:    -1,
		currentViewport:      math.NewVector4(0, 0, 0, 0),
		currentScissor:       math.NewVector4(0, 0, 0, 0),
		usedTextureUnits:     0,
		width:                width,
		height:               height,
		pixelRatio:           1,
		viewPort:             math.NewVector4(0, 0, float32(width), float32(height)),
		scissor:              math.NewVector4(0, 0, float32(width), float32(height)),
		scissorTest:          false,
		frustum:              &math.Frustum{},
		clippingEnabled:      false,
		LocalClippingEnabled: false,
		projectScreenMatrix:  math.NewDefaultMatrix4(),
		vector3:              math.NewDefaultVector3(),

		contextAttributes: options,
	}

	// Event listeners must be registered before WebGL context is created, see Three.js #12753
	renderer.onContextLostJs = js.NewCallback(func(args []js.Value) { renderer.onContextLost() })
	renderer.onContextRestoreJs = js.NewCallback(func(args []js.Value) { renderer.onContextRestore() })

	canvas.Call("addEventListener", "webglcontextlost", renderer.onContextLostJs, false)
	canvas.Call("addEventListener", "webglcontextrestored", renderer.onContextRestoreJs, false)

	renderer.glContext = canvas.Call("getContext", "webgl")
	if renderer.glContext == js.Undefined() {
		renderer.glContext = canvas.Call("getContext", "experimental-webgl")
	}
	if renderer.glContext == js.Undefined() {
		return nil, errors.New("WebGL context can't be created. Maybe the browser don't support then")
	}

	renderer.extensions = webgl.NewExtensions(renderer.glContext)

	return &renderer, nil
}

func (renderer *Renderer) Close() {
	defer renderer.onContextLostJs.Release()
	defer renderer.onContextRestoreJs.Release()
}

func (renderer *Renderer) Render(scene *core.Scene, camera *core.Camera, target *render.Target, forceClear bool) error {
	if camera == nil {
		return errors.New("the camera can't be null")
	}

	if renderer.isContextLost {
		return nil
	}

	// reset caching for this frame
	//renderer.currentGeometryProgram.geometry = nil
	//renderer.currentGeometryProgram.program = nil
	//renderer.currentGeometryProgram.wireFrame = false
	renderer.currentMaterialId = -1
	renderer.currentCamera = nil

	// update scene graph
	if scene.AutoUpdateRender {
		scene.UpdateMatrixWorld(false)
	}

	// update camera matrices and frustum
	if camera.GetParent() == nil {
		camera.UpdateMatrixWorld(false)
	}

	//renderer.currentRenderState = renderer.renderStates.Get(scene, camera)
	//renderer.currentRenderState.Init()
	return nil
}

func (renderer *Renderer) GetTargetPixelRatio() float32 {
	return 0 // TODO:
}

func (renderer *Renderer) onContextLost() {

}

func (renderer *Renderer) onContextRestore() {

}
