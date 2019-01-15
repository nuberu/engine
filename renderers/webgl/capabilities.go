package webgl

import (
	"github.com/tokkenno/seed/renderers/webgl/constant"
	"log"
	"syscall/js"
)

type Capabilities struct {
	glContext js.Value

	maxAnisotropy          float64
	precision              constant.Precision
	logarithmicDepthBuffer bool
	maxTextures            int
	maxVertexTextures      int
	maxTextureSize         int
	maxCubemapSize         int
	maxAttributes          int
	maxVertexUniforms      int
	maxVaryings            int
	maxFragmentUniforms    int
	vertexTextures         bool
	floatFragmentTextures  bool
	floatVertexTextures    bool
	webGLVersion           int
}

func NewCapabilities(glContext js.Value, extensions *Extensions, settings *Settings) *Capabilities {
	capabilities := &Capabilities{
		glContext:    glContext,
		webGLVersion: 1, // TODO: Add support to 2
	}

	// Max Anisotropy
	extension := extensions.Get("EXT_texture_filter_anisotropic")
	if extension != js.Undefined() {
		capabilities.maxAnisotropy = glContext.Call("getParameter", constant.MaxTextureMaxAnisotropyExt).Float()
	} else {
		capabilities.maxAnisotropy = 0
	}

	// Precision
	if glContext.Call("getShaderPrecisionFormat", glContext.Get("VERTEX_SHADER"), glContext.Get("HIGH_FLOAT ")).Get("precision").Int() > 0 &&
		glContext.Call("getShaderPrecisionFormat", glContext.Get("FRAGMENT_SHADER"), glContext.Get("HIGH_FLOAT ")).Get("precision").Int() > 0 {
		capabilities.precision = constant.HighPrecision
	} else if glContext.Call("getShaderPrecisionFormat", glContext.Get("VERTEX_SHADER"), glContext.Get("MEDIUM_FLOAT  ")).Get("precision").Int() > 0 &&
		glContext.Call("getShaderPrecisionFormat", glContext.Get("FRAGMENT_SHADER"), glContext.Get("MEDIUM_FLOAT  ")).Get("precision").Int() > 0 {
		capabilities.precision = constant.MediumPrecision
	} else {
		capabilities.precision = constant.LowPrecision
	}

	if settings.Precision != capabilities.precision {
		log.Printf("%s not supported, using %s", settings.Precision, capabilities.precision)
		settings.Precision = capabilities.precision
	}

	capabilities.logarithmicDepthBuffer = settings.LogarithmicDepthBuffer

	capabilities.maxTextures = glContext.Call("getParameter", glContext.Get("MAX_TEXTURE_IMAGE_UNITS")).Int()
	capabilities.maxVertexTextures = glContext.Call("getParameter", glContext.Get("MAX_VERTEX_TEXTURE_IMAGE_UNITS")).Int()
	capabilities.maxTextureSize = glContext.Call("getParameter", glContext.Get("MAX_TEXTURE_SIZE")).Int()
	capabilities.maxCubemapSize = glContext.Call("getParameter", glContext.Get("MAX_CUBE_MAP_TEXTURE_SIZE")).Int()

	capabilities.maxAttributes = glContext.Call("getParameter", glContext.Get("MAX_VERTEX_ATTRIBS")).Int()
	capabilities.maxVertexUniforms = glContext.Call("getParameter", glContext.Get("MAX_VERTEX_UNIFORM_VECTORS")).Int()
	capabilities.maxVaryings = glContext.Call("getParameter", glContext.Get("MAX_VARYING_VECTORS")).Int()
	capabilities.maxFragmentUniforms = glContext.Call("getParameter", glContext.Get("MAX_FRAGMENT_UNIFORM_VECTORS")).Int()

	capabilities.vertexTextures = capabilities.maxVertexTextures > 0
	capabilities.floatFragmentTextures = capabilities.webGLVersion > 1 || extensions.Get("OES_texture_float") != js.Undefined()
	capabilities.floatVertexTextures = capabilities.vertexTextures && capabilities.floatFragmentTextures

	return capabilities
}

func (cap *Capabilities) GetMaxAnisotropy() float64 {
	return cap.maxAnisotropy
}

func (cap *Capabilities) GetWebGLVersion() int {
	return cap.webGLVersion
}

func (cap *Capabilities) GetMaxPrecision(precision constant.Precision) constant.Precision {
	if precision == constant.HighPrecision && cap.precision == constant.HighPrecision {
		return constant.HighPrecision
	} else if precision == constant.MediumPrecision && (cap.precision == constant.HighPrecision || cap.precision == constant.MediumPrecision) {
		return constant.MediumPrecision
	} else {
		return constant.LowPrecision
	}
}
