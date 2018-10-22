package core

import (
	"github.com/tokkenno/seed/math"
)

type Face3 struct {
	A             math.Vector3
	B             math.Vector3
	C             math.Vector3
	Normal        math.Vector3
	VertexNormals [3]math.Vector3
	Color         math.Color
	MaterialIndex uint
}

func NewDefaultFace3(a, b, c *math.Vector3) *Face3 {
	return NewFace3(a, b, c, math.NewDefaultVector3(), math.ColorBlack, 0)
}

func NewFace3(a, b, c, normal *math.Vector3, color math.Color, materialIndex uint) *Face3 {
	return &Face3{
		A:             *a.Clone(),
		B:             *b.Clone(),
		C:             *c.Clone(),
		Normal:        *normal.Clone(),
		VertexNormals: [3]math.Vector3{},
		Color:         color,
		MaterialIndex: materialIndex,
	}
}

func (face *Face3) Clone() *Face3 {
	f := &Face3{}
	f.Copy(face)
	return f
}

func (face *Face3) Copy(f *Face3) {
	face.A = *f.A.Clone()
	face.B = *f.B.Clone()
	face.C = *f.C.Clone()
	face.Normal = *f.Normal.Clone()
	face.Color = math.Color{R: f.Color.R, G: f.Color.G, B: f.Color.B, A: f.Color.A}
	face.MaterialIndex = f.MaterialIndex
}
