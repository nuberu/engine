package renderer

import (
	gl "github.com/nuberu/webgl"
	"errors"
	"github.com/nuberu/engine/core"
	"github.com/nuberu/engine/core/constant"
	"github.com/nuberu/engine/core/render"
	"github.com/nuberu/engine/math"
	"github.com/nuberu/engine/renderers/webgl"
	"github.com/nuberu/engine/renderers/webgl/program"
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

	PhysicallyCorrectLights bool
	ToneMapping             constant.ToneMapping
	ToneMappingExposure     float32
	ToneMappingWhitePoint   float32

	MaxMorphTargets int
	MaxMorphNormals int

	isContextLost bool

	//frameBuffer

	renderStates *States
	currentRenderState *State
	//currentRenderTarget
	//currentFrameBuffer
	currentMaterialId int64

	currentGeometryProgram *program.Geometry

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

	contextAttributes *webgl.Settings

	extensions *webgl.Extensions
	capabilities *webgl.Capabilities

	glContext          *gl.RenderingContext
	onContextLostJs    js.Callback
	onContextRestoreJs js.Callback
}

func NewRenderer(settings *webgl.Settings) (*Renderer, error) {
	if settings == nil {
		settings = webgl.DefaultSettings()
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
		PhysicallyCorrectLights: false,
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

		contextAttributes: settings,
	}

	// Event listeners must be registered before WebGL context is created, see Three.js #12753
	renderer.onContextLostJs = js.NewCallback(func(args []js.Value) { renderer.onContextLost() })
	renderer.onContextRestoreJs = js.NewCallback(func(args []js.Value) { renderer.onContextRestore() })

	canvas.Call("addEventListener", "webglcontextlost", renderer.onContextLostJs, false)
	canvas.Call("addEventListener", "webglcontextrestored", renderer.onContextRestoreJs, false)

	context, err := gl.FromCanvas(canvas)
	if err != nil {
		return nil, err
	}
	renderer.glContext = context

	// InitGLContext
	renderer.extensions = webgl.NewExtensions(renderer.glContext)
	renderer.capabilities = webgl.NewCapabilities(renderer.glContext, renderer.extensions, settings)

	if renderer.capabilities.GetWebGLVersion() < 2 {
		renderer.extensions.Get("WEBGL_depth_texture")
		renderer.extensions.Get("OES_texture_float")
		renderer.extensions.Get("OES_texture_half_float")
		renderer.extensions.Get("OES_texture_half_float_linear")
		renderer.extensions.Get("OES_standard_derivatives")
		renderer.extensions.Get("OES_element_index_uint")
		renderer.extensions.Get("ANGLE_instanced_arrays")
	}

	renderer.extensions.Get("OES_texture_float_linear")

	// TODO: WebGlUtils: line 265

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
	renderer.currentGeometryProgram.Geometry = nil
	renderer.currentGeometryProgram.Program = 0
	renderer.currentGeometryProgram.WireFrame = false
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

	renderer.currentRenderState = renderer.renderStates.Get(scene, camera)
	renderer.currentRenderState.Restart()

	// scene.onBeforeRender( _this, scene, camera, renderTarget ); Line 1065

	return nil
}

func (renderer *Renderer) GetTargetPixelRatio() float32 {
	return 0 // TODO:
}

func (renderer *Renderer) onContextLost() {

}

func (renderer *Renderer) onContextRestore() {

}
