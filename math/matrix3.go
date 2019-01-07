package math

import (
	"errors"
	"fmt"
	"log"
)

type Matrix3 struct {
	elements [9]float32
}

func NewDefaultMatrix3() *Matrix3 {
	matrix := &Matrix3{
		elements: [9]float32{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	matrix.SetIdentity()
	return matrix
}

func NewMatrix3(n11, n12, n13, n21, n22, n23, n31, n32, n33 float32) *Matrix3 {
	matrix := &Matrix3{
		elements: [9]float32{n11, n12, n13, n21, n22, n23, n31, n32, n33},
	}
	return matrix
}

func (matrix *Matrix3) GetElements() [9]float32 {
	return matrix.elements
}

func (matrix *Matrix3) Set(n11, n12, n13, n21, n22, n23, n31, n32, n33 float32) {
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

func (matrix *Matrix3) SetIdentity() {
	matrix.elements = [9]float32{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func (matrix *Matrix3) Clone() *Matrix3 {
	m := Matrix3{
		elements: [9]float32{},
	}
	copy(m.elements[0:], matrix.elements[0:])
	return &m
}

func (matrix *Matrix3) Copy(other *Matrix3) {
	copy(matrix.elements[0:], other.elements[0:])
}

func (matrix *Matrix3) Multiply(m *Matrix3) {
	matrix.MultiplyMatrices(matrix, m)
}

func (matrix *Matrix3) PreMultiply(m *Matrix3) {
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

func (matrix *Matrix3) MultiplyScalar(scalar float32) {
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

func (matrix *Matrix3) Determinant() float32 {
	return matrix.elements[0]*matrix.elements[4]*matrix.elements[8] -
		matrix.elements[0]*matrix.elements[5]*matrix.elements[7] -
		matrix.elements[1]*matrix.elements[3]*matrix.elements[8] +
		matrix.elements[1]*matrix.elements[5]*matrix.elements[6] +
		matrix.elements[2]*matrix.elements[3]*matrix.elements[7] -
		matrix.elements[2]*matrix.elements[4]*matrix.elements[6]
}

func (matrix *Matrix3) Inverse() error {
	return matrix.SetInverse(matrix.Clone(), true)
}

// Set the inverse of [ma] in the current [matrix]
func (matrix *Matrix3) SetInverse(ma *Matrix3, errorOnDegenerate bool) error {
	t11 := ma.elements[8]*ma.elements[4] - ma.elements[5]*ma.elements[7]
	t12 := ma.elements[5]*ma.elements[6] - ma.elements[8]*ma.elements[3]
	t13 := ma.elements[7]*ma.elements[3] - ma.elements[4]*ma.elements[6]

	det := ma.elements[0]*t11 + ma.elements[1]*t12 + ma.elements[2]*t13

	if det == 0 {
		if errorOnDegenerate == true {
			return errors.New(".SetInverse() can't invert matrix, determinant is 0")
		} else {
			log.Println(".SetInverse() can't invert matrix, determinant is 0")
			matrix.SetIdentity()
			return nil
		}
	}

	detInv := 1 / det

	matrix.elements[0] = t11 * detInv
	matrix.elements[1] = (ma.elements[2]*ma.elements[7] - ma.elements[8]*ma.elements[1]) * detInv
	matrix.elements[2] = (ma.elements[5]*ma.elements[1] - ma.elements[2]*ma.elements[4]) * detInv
	matrix.elements[3] = t12 * detInv
	matrix.elements[4] = (ma.elements[8]*ma.elements[0] - ma.elements[2]*ma.elements[7]) * detInv
	matrix.elements[5] = (ma.elements[2]*ma.elements[3] - ma.elements[5]*ma.elements[0]) * detInv
	matrix.elements[6] = t13 * detInv
	matrix.elements[7] = (ma.elements[1]*ma.elements[6] - ma.elements[7]*ma.elements[0]) * detInv
	matrix.elements[8] = (ma.elements[4]*ma.elements[0] - ma.elements[1]*ma.elements[3]) * detInv

	return nil
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

// Set the transpose of [ma] in the current [matrix]
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

// Set the normal matrix of [ma] in the current matrix
func (matrix *Matrix3) SetNormalMatrix(ma *Matrix4) {
	matrix.SetFromMatrix4(ma)
	matrix.SetInverse(matrix, false)
	matrix.Transpose()
}

func (matrix *Matrix3) SetUvTransform(tx, ty, sx, sy, rotation, cx, cy float32) {
	cos := Cos(rotation)
	sin := Sin(rotation)

	matrix.Set(
		sx*cos, sx*sin, -sx*(cos*cx+sin*cy)+cx+tx,
		-sy*sin, sy*cos, -sy*(-sin*cx+cos*cy)+cy+ty,
		0, 0, 1,
	)
}

func (matrix *Matrix3) Scale(sx float32, sy float32) {
	matrix.elements[0] *= sx
	matrix.elements[1] *= sy

	matrix.elements[3] *= sx
	matrix.elements[4] *= sy

	matrix.elements[6] *= sx
	matrix.elements[7] *= sy
}

func (matrix *Matrix3) Rotate(rotation float32) {
	cos := Cos(rotation)
	sin := Sin(rotation)

	cpMatrix := matrix.Clone()

	matrix.elements[0] = cos*cpMatrix.elements[0] + sin*cpMatrix.elements[1]
	matrix.elements[3] = cos*cpMatrix.elements[3] + sin*cpMatrix.elements[4]
	matrix.elements[6] = cos*cpMatrix.elements[6] + sin*cpMatrix.elements[7]

	matrix.elements[1] = cos*cpMatrix.elements[0] + sin*cpMatrix.elements[1]
	matrix.elements[4] = cos*cpMatrix.elements[3] + sin*cpMatrix.elements[4]
	matrix.elements[7] = cos*cpMatrix.elements[6] + sin*cpMatrix.elements[7]
}

func (matrix *Matrix3) Translate(tx float32, ty float32) {
	matrix.elements[0] += tx * matrix.elements[2]
	matrix.elements[3] += tx * matrix.elements[5]
	matrix.elements[6] += tx * matrix.elements[8]
	matrix.elements[1] += ty * matrix.elements[2]
	matrix.elements[4] += ty * matrix.elements[5]
	matrix.elements[7] += ty * matrix.elements[8]
}

func (matrix *Matrix3) TranslateVector2(v *Vector2) {
	matrix.Translate(v.X, v.Y)
}

func (matrix *Matrix3) Equals(ma *Matrix3) bool {
	for ind := range matrix.elements {
		if matrix.elements[ind] != ma.elements[ind] {
			return false
		}
	}
	return true
}

func (matrix *Matrix3) EqualsRound(ma *Matrix3, decimals float32) bool {
	mul := Pow(10, decimals)
	for ind := range matrix.elements {
		if Round(mul * matrix.elements[ind]) / mul != Round(mul * ma.elements[ind]) / mul {
			return false
		}
	}
	return true
}

func (matrix *Matrix3) ToArray() [9]float32 {
	mc := matrix.Clone()
	return mc.elements
}

func (matrix *Matrix3) CopyToArray(array []float32, offset int) {
	va := matrix.ToArray()
	copy(array[offset:], va[0:])
}

func (matrix *Matrix3) ToString() string {
	return fmt.Sprintf("%9.2f %9.2f %9.2f\n%9.2f %9.2f %9.2f\n%9.2f %9.2f %9.2f",
		matrix.elements[0], matrix.elements[1], matrix.elements[2],
		matrix.elements[3], matrix.elements[4], matrix.elements[5],
		matrix.elements[6], matrix.elements[7], matrix.elements[8])
}