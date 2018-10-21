package core

import (
	"github.com/tokkenno/seed/core/constant"
	"github.com/tokkenno/seed/core/types"
	"github.com/tokkenno/seed/math"
)

var (
	materialIdGenerator = new(IdGenerator)
)

type Material struct {
	Object3

	fog                 bool
	lights              bool
	blending            constant.Blending
	side                constant.Side
	flatShading         bool
	vertexColors        constant.Color
	opacity             types.Subset
	transparent         bool
	blendSrc            constant.Factor
	blendDst            constant.Factor
	blendEquation       constant.Equation
	blendSrcAlpha       constant.Factor
	blendDstAlpha       constant.Factor
	blendEquationAlpha  constant.Equation
	depthFunc           constant.Depth
	depthTest           bool
	depthWrite          bool
	clippingPlanes      []math.Plane
	clipIntersection    bool
	clipShadows         bool
	shadowSide          constant.Side
	colorWrite          bool
	precision           uint // override the renderer's default precision for this material
	polygonOffset       bool
	polygonOffsetFactor uint
	polygonOffsetUnits  uint
	dithering           bool
	alphaTest           uint
	preMultipliedAlpha  bool
	overdraw            uint // Overdrawn pixels (typically between 0 and 1) for fixing antialiasing gaps in CanvasRenderer
	needsUpdate         bool
}

func NewMaterial() *Material {
	mat := &Material{
		Object3:             *newObjectWithId(materialIdGenerator.Next()),
		fog:                 true,
		lights:              true,
		blending:            constant.NormalBlending,
		side:                constant.FrontSide,
		flatShading:         false,
		vertexColors:        constant.NoColors,
		opacity:             types.NewSubset(0, 1, 1),
		transparent:         false,
		blendSrc:            constant.SrcAlphaFactor,
		blendDst:            constant.OneMinusSrcAlphaFactor,
		blendEquation:       constant.AddEquation,
		blendSrcAlpha:       constant.NilFactor,
		blendDstAlpha:       constant.NilFactor,
		blendEquationAlpha:  constant.NilEquation,
		depthFunc:           constant.LessEqualDepth,
		depthTest:           true,
		depthWrite:          true,
		clippingPlanes:      []math.Plane{},
		clipIntersection:    false,
		clipShadows:         false,
		shadowSide:          constant.NoSide,
		colorWrite:          true,
		precision:           0,
		polygonOffset:       false,
		polygonOffsetFactor: 0,
		polygonOffsetUnits:  0,
		dithering:           false,
		alphaTest:           0,
		preMultipliedAlpha:  false,
		overdraw:            0,
		needsUpdate:         true,
	}

	mat.Visible = true

	return mat
}

func (mat *Material) Clone() *Material {
	newMat := NewMaterial()
	newMat.Copy(mat)
	return newMat
}

func (mat *Material) Copy(source *Material) {
	mat.fog = source.fog
	mat.lights = source.lights

	mat.blending = source.blending
	mat.side = source.side
	mat.flatShading = source.flatShading
	mat.vertexColors = source.vertexColors

	mat.opacity = source.opacity.Clone()
	mat.transparent = source.transparent

	mat.blendSrc = source.blendSrc
	mat.blendDst = source.blendDst
	mat.blendEquation = source.blendEquation
	mat.blendSrcAlpha = source.blendSrcAlpha
	mat.blendDstAlpha = source.blendDstAlpha
	mat.blendEquationAlpha = source.blendEquationAlpha

	mat.depthFunc = source.depthFunc
	mat.depthTest = source.depthTest
	mat.depthWrite = source.depthWrite

	mat.colorWrite = source.colorWrite

	mat.precision = source.precision

	mat.polygonOffset = source.polygonOffset
	mat.polygonOffsetFactor = source.polygonOffsetFactor
	mat.polygonOffsetUnits = source.polygonOffsetUnits

	mat.dithering = source.dithering

	mat.alphaTest = source.alphaTest
	mat.preMultipliedAlpha = source.preMultipliedAlpha

	mat.overdraw = source.overdraw

	mat.Visible = source.Visible

	mat.clipShadows = source.clipShadows
	mat.clipIntersection = source.clipIntersection



	clippingPlanesCpy := make([]math.Plane, len(source.clippingPlanes))
	for i, plane := range source.clippingPlanes {
		clippingPlanesCpy[i] = *plane.Clone()
	}

	mat.clippingPlanes = clippingPlanesCpy

	mat.shadowSide = source.shadowSide
}