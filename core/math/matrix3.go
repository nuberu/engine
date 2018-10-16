package math

import "math"

type Matrix3 struct {
	elements [9]float64
}

func Matrix3Identity() *Matrix3 {
	return &Matrix3{
		elements: [9]float64{
			1, 0, 0,
			0, 1, 0,
			0, 0, 1,
		},
	}
}

func (matrix *Matrix3) GetElements() [9]float64 {
	return matrix.elements
}

func (matrix *Matrix3) Set(n11, n12, n13, n21, n22, n23, n31, n32, n33 float64) {
	matrix.elements[0] = n11
	matrix.elements[1] = n12
	matrix.elements[2] = n13
	matrix.elements[3] = n21
	matrix.elements[4] = n22
	matrix.elements[5] = n23
	matrix.elements[6] = n31
	matrix.elements[7] = n32
	matrix.elements[8] = n33
}

func (matrix *Matrix3) SetFromMatrix4(m *Matrix4) {
	me := m.GetElements()
	matrix.Set(
		me[0], me[4], me[8],
		me[1], me[5], me[9],
		me[2], me[6], me[10],
	)
}

func (matrix *Matrix3) Clone() *Matrix3 {
	m := Matrix3{
		elements: [9]float64{},
	}
	copy(m.elements[0:], matrix.elements[0:])
	return &m
}

func (matrix *Matrix3) Multiply(m *Matrix3) {
	matrix.MultiplyMatrices(m, matrix)
}

func (matrix *Matrix3) MultiplyMatrices(ma *Matrix3, mb *Matrix3) {
	mae := ma.GetElements()
	mbe := mb.GetElements()

	matrix.elements[0] = mae[0]*mbe[0] + mae[3]*mbe[1] + mae[6]*mbe[2]
	matrix.elements[3] = mae[0]*mbe[3] + mae[3]*mbe[4] + mae[6]*mbe[5]
	matrix.elements[6] = mae[0]*mbe[6] + mae[3]*mbe[7] + mae[6]*mbe[8]

	matrix.elements[1] = mae[1]*mbe[0] + mae[4]*mbe[1] + mae[7]*mbe[2]
	matrix.elements[4] = mae[1]*mbe[3] + mae[4]*mbe[4] + mae[7]*mbe[5]
	matrix.elements[7] = mae[1]*mbe[6] + mae[4]*mbe[7] + mae[7]*mbe[8]

	matrix.elements[2] = mae[2]*mbe[0] + mae[5]*mbe[1] + mae[8]*mbe[2]
	matrix.elements[5] = mae[2]*mbe[3] + mae[5]*mbe[4] + mae[8]*mbe[5]
	matrix.elements[8] = mae[2]*mbe[6] + mae[5]*mbe[7] + mae[8]*mbe[8]
}

func (matrix *Matrix3) MultiplyScalar(scalar float64) {
	matrix.elements[0] *= scalar
	matrix.elements[1] *= scalar
	matrix.elements[2] *= scalar

	matrix.elements[3] *= scalar
	matrix.elements[4] *= scalar
	matrix.elements[5] *= scalar

	matrix.elements[6] *= scalar
	matrix.elements[7] *= scalar
	matrix.elements[8] *= scalar
}

func (matrix *Matrix3) Determinant() float64 {
	return matrix.elements[0]*matrix.elements[4]*matrix.elements[8] -
		matrix.elements[0]*matrix.elements[5]*matrix.elements[7] -
		matrix.elements[1]*matrix.elements[3]*matrix.elements[8] +
		matrix.elements[1]*matrix.elements[5]*matrix.elements[6] +
		matrix.elements[2]*matrix.elements[3]*matrix.elements[7] -
		matrix.elements[2]*matrix.elements[4]*matrix.elements[6]
}

func (matrix *Matrix3) GetInverse(ma *Matrix3) {

}

func (matrix *Matrix3) Transpose() {
	tmp := matrix.elements[1]
	matrix.elements[1] = matrix.elements[3]
	matrix.elements[3] = tmp

	tmp = matrix.elements[2]
	matrix.elements[2] = matrix.elements[6]
	matrix.elements[6] = tmp

	tmp = matrix.elements[5]
	matrix.elements[5] = matrix.elements[7]
	matrix.elements[7] = tmp
}

func (matrix *Matrix3) SetTranspose(ma *Matrix3) {
	matrix.elements[0] = ma.elements[0]
	matrix.elements[1] = ma.elements[3]
	matrix.elements[2] = ma.elements[6]

	matrix.elements[3] = ma.elements[1]
	matrix.elements[4] = ma.elements[4]
	matrix.elements[5] = ma.elements[7]

	matrix.elements[6] = ma.elements[2]
	matrix.elements[7] = ma.elements[5]
	matrix.elements[8] = ma.elements[8]
}

func (matrix *Matrix3) GetNormalMatrix(m4 *Matrix4) {
	matrix.SetFromMatrix4(m4)
	matrix.GetInverse(matrix)
	matrix.Transpose()
}

func (matrix *Matrix3) SetUvTransform(tx, ty, sx, sy, rotation, cx, cy float64) {
	cos := math.Cos(rotation)
	sin := math.Sin(rotation)

	matrix.Set(
		sx*cos, sx*sin, -sx*(cos*cx+sin*cy)+cx+tx,
		-sy*sin, sy*cos, -sy*(-sin*cx+cos*cy)+cy+ty,
		0, 0, 1,
	)
}

func (matrix *Matrix3) Scale(sx float64, sy float64) {
	matrix.elements[0] *= sx
	matrix.elements[1] *= sy

	matrix.elements[3] *= sx
	matrix.elements[4] *= sy

	matrix.elements[6] *= sx
	matrix.elements[7] *= sy
}

func (matrix *Matrix3) Rotate(rotation float64) {
	cos := math.Cos(rotation)
	sin := math.Sin(rotation)

	cpMatrix := matrix.Clone()

	matrix.elements[0] = cos*cpMatrix.elements[0] + sin*cpMatrix.elements[1]
	matrix.elements[3] = cos*cpMatrix.elements[3] + sin*cpMatrix.elements[4]
	matrix.elements[6] = cos*cpMatrix.elements[6] + sin*cpMatrix.elements[7]

	matrix.elements[1] = cos*cpMatrix.elements[0] + sin*cpMatrix.elements[1]
	matrix.elements[4] = cos*cpMatrix.elements[3] + sin*cpMatrix.elements[4]
	matrix.elements[7] = cos*cpMatrix.elements[6] + sin*cpMatrix.elements[7]
}

func (matrix *Matrix3) TranslateVector2(v *Vector2) {
	matrix.Translate(v.X, v.Y)
}

func (matrix *Matrix3) Translate(tx float64, ty float64) {
	matrix.elements[0] += tx * matrix.elements[2]
	matrix.elements[3] += tx * matrix.elements[5]
	matrix.elements[6] += tx * matrix.elements[8]
	matrix.elements[1] += ty * matrix.elements[2]
	matrix.elements[4] += ty * matrix.elements[5]
	matrix.elements[7] += ty * matrix.elements[8]
}


func (matrix *Matrix3) Equals(ma *Matrix3) bool {
	for ind := range matrix.elements {
		if matrix.elements[ind] != matrix.elements[ind] {
			return false
		}
	}
	return true
}