package core

import (
	"github.com/nuberu/engine/math"
)

type IGeometry interface {
	ApplyMatrix(m *math.Matrix4)
	RotateX(a math.Angle)
	RotateY(a math.Angle)
	RotateZ(a math.Angle)
	Translate(x, y, z float32)
	Scale(x, y, z float32)
	LookAt(v *math.Vector3)
	Center()
	Normalize()
}

// The basic geometry
type Geometry struct {
	IGeometry

	vertices []*math.Vector3
	faces    []*Face3

	boundingBox    *math.Box3
	boundingSphere *math.Sphere

	verticesNeedUpdate bool
	normalsNeedUpdate  bool
}

func (geo *Geometry) GetBoundingBox() *math.Box3 {
	if geo.boundingBox == nil {
		geo.computeBoundingBox()
	}
	return geo.boundingBox
}

func (geo *Geometry) GetBoundingSphere() *math.Sphere {
	if geo.boundingSphere == nil {
		geo.computeBoundingSphere()
	}
	return geo.boundingSphere
}

func (geo *Geometry) ApplyMatrix(matrix *math.Matrix4) {
	normalMatrix := math.NewDefaultMatrix3()
	normalMatrix.SetNormalMatrix(matrix)

	for i := 0; i < len(geo.vertices); i++ {
		vertex := geo.vertices[i]
		vertex.ApplyMatrix4(matrix)
	}

	for i := 0; i < len(geo.faces); i ++ {
		face := geo.faces[i]
		face.Normal.ApplyMatrix3(normalMatrix)
		face.Normal.Normalize()

		for j := 0; j < len(face.VertexNormals); j++ {
			face.VertexNormals[j].ApplyMatrix3(normalMatrix)
			face.VertexNormals[j].Normalize()
		}
	}

	if geo.boundingBox != nil {
		geo.computeBoundingBox()
	}

	if geo.boundingSphere != nil {
		geo.computeBoundingSphere()
	}

	geo.verticesNeedUpdate = true
	geo.normalsNeedUpdate = true
}

func (geo *Geometry) RotateX(angle math.Angle) {
	m1 := math.NewMatrix4RotationX(angle)
	geo.ApplyMatrix(m1)
}

func (geo *Geometry) RotateY(angle math.Angle) {
	m1 := math.NewMatrix4RotationY(angle)
	geo.ApplyMatrix(m1)
}

func (geo *Geometry) RotateZ(angle math.Angle) {
	m1 := math.NewMatrix4RotationZ(angle)
	geo.ApplyMatrix(m1)
}

func (geo *Geometry) Translate(x, y, z float32) {
	m1 := math.NewMatrix4Translation(x, y, z)
	geo.ApplyMatrix(m1)
}

func (geo *Geometry) TranslateVector3(v *math.Vector3) {
	geo.Translate(v.X, v.Y, v.Z)
}

func (geo *Geometry) Scale(x, y, z float32) {
	m1 := math.NewMatrix4Scale(x, y, z)
	geo.ApplyMatrix(m1)
}

func (geo *Geometry) LookAt(vector *math.Vector3) {
	obj := NewObject()
	obj.LookAt(vector)
	obj.UpdateMatrix()
}

func (geo *Geometry) Center() {
	geo.computeBoundingBox()

	offset := geo.boundingBox.GetCenter()
	offset.Negate()

	geo.TranslateVector3(offset)
}

func (geo *Geometry) Normalize() {
	geo.computeBoundingSphere()

	var center = geo.boundingSphere.Center
	var radius = geo.boundingSphere.Radius

	s := float32(0)
	if radius == 0 {
		s = 1.0 / radius
	}

	var matrix = math.NewDefaultMatrix4()
	matrix.Set(
		s, 0, 0, - s*center.X,
		0, s, 0, - s*center.Y,
		0, 0, s, - s*center.Z,
		0, 0, 0, 1,
	)

	geo.ApplyMatrix(matrix)
}

func (geo *Geometry) computeBoundingBox() {
	if geo.boundingBox == nil {
		geo.boundingBox = math.NewDefaultBox3()
	}

	geo.boundingBox.SetFromPoints(geo.vertices)
}

func (geo *Geometry) computeBoundingSphere() {
	if geo.boundingSphere == nil {
		geo.boundingSphere = math.NewDefaultSphere()
	}

	geo.boundingSphere.SetFromPoints(geo.vertices)
}