package math

import (
	"errors"
	"fmt"
	"github.com/google/logger"
	"github.com/tokkenno/seed/core/types"
)

type Matrix4 struct {
	elements [16]float32
}

func NewDefaultMatrix4() *Matrix4 {
	matrix := &Matrix4{
		elements: [16]float32{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		},
	}
	matrix.SetIdentity()
	return matrix
}

func NewMatrix4(n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44 float32) *Matrix4 {
	return &Matrix4{
		elements: [16]float32{
			n11, n12, n13, n14,
			n21, n22, n23, n24,
			n31, n32, n33, n34,
			n41, n42, n43, n44,
		},
	}
}

func NewMatrix4Translation(x float32, y float32, z float32) *Matrix4 {
	return &Matrix4{
		elements: [16]float32{
			1, 0, 0, x,
			0, 1, 0, y,
			0, 0, 1, z,
			0, 0, 0, 1,
		},
	}
}

func NewMatrix4RotationX(angle types.Angle) *Matrix4 {
	c := Cos(float32(angle))
	s := Sin(float32(angle))

	return &Matrix4{
		elements: [16]float32{
			1, 0, 0, 0,
			0, c, -s, 0,
			0, s, c, 0,
			0, 0, 0, 1,
		},
	}
}

func NewMatrix4RotationY(angle types.Angle) *Matrix4 {
	c := Cos(float32(angle))
	s := Sin(float32(angle))

	return &Matrix4{
		elements: [16]float32{
			c, 0, s, 0,
			0, 1, 0, 0,
			-s, 0, c, 0,
			0, 0, 0, 1,
		},
	}
}

func NewMatrix4RotationZ(angle types.Angle) *Matrix4 {
	c := Cos(float32(angle))
	s := Sin(float32(angle))

	return &Matrix4{
		elements: [16]float32{
			c, -s, 0, 0,
			s, c, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		},
	}
}

func NewMatrix4RotationAxis(axis *Vector3, angle types.Angle) *Matrix4 {
	c := Cos(float32(angle))
	s := Sin(float32(angle))
	t := 1 - c
	tx := t * axis.X
	ty := t * axis.Y

	return &Matrix4{
		elements: [16]float32{
			tx*axis.X + c, tx*axis.Y - s*axis.Z, tx*axis.Z + s*axis.Y, 0,
			tx*axis.Y + s*axis.Z, ty*axis.Y + c, ty*axis.Z - s*axis.X, 0,
			tx*axis.Z - s*axis.Y, ty*axis.Z + s*axis.X, t*axis.Z*axis.Z + c, 0,
			0, 0, 0, 1,
		},
	}
}

func NewMatrix4Scale(x float32, y float32, z float32) *Matrix4 {
	return &Matrix4{
		elements: [16]float32{
			x, 0, 0, 0,
			0, y, 0, 0,
			0, 0, z, 0,
			0, 0, 0, 1,
		},
	}
}

func NewMatrix4Shear(x float32, y float32, z float32) *Matrix4 {
	return &Matrix4{
		elements: [16]float32{
			1, y, z, 0,
			x, 1, z, 0,
			x, y, 1, 0,
			0, 0, 0, 1,
		},
	}
}

// TODO: Move to camera (check uses)
func NewMatrix4Perspective(left, right, top, bottom, near, far float32) *Matrix4 {
	x := 2 * near / (right - left)
	y := 2 * near / (top - bottom)

	a := (right + left) / (right - left)
	b := (top + bottom) / (top - bottom)
	c := -(far + near) / (far - near)
	d := -2 * far * near / (far - near)

	return &Matrix4{
		elements: [16]float32{
			x, 0, 0, 0,
			0, y, 0, 0,
			a, b, c, -1,
			0, 0, d, 0,
		},
	}
}

// TODO: Move to camera (check uses)
func NewMatrix4Orthographic(left, right, top, bottom, near, far float32) *Matrix4 {
	w := 1.0 / (right - left)
	h := 1.0 / (top - bottom)
	p := 1.0 / (far - near)

	x := (right + left) * w
	y := (top + bottom) * h
	z := (far + near) * p

	return &Matrix4{
		elements: [16]float32{
			2 * w, 0, 0, 0,
			0, 2 * h, 0, 0,
			0, 0, -2 * p, 0,
			-x, -y, -z, 1,
		},
	}
}

func NewMatrix4FromArray(arr []float32, offset int) *Matrix4 {
	return NewMatrix4(
		arr[offset], arr[offset+1], arr[offset+2], arr[offset+3],
		arr[offset+4], arr[offset+5], arr[offset+6], arr[offset+7],
		arr[offset+8], arr[offset+9], arr[offset+10], arr[offset+11],
		arr[offset+12], arr[offset+13], arr[offset+14], arr[offset+15],
	)
}

func (matrix *Matrix4) GetElements() [16]float32 {
	return matrix.elements
}

func (matrix *Matrix4) Set(n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44 float32) {
	matrix.elements[0] = n11
	matrix.elements[1] = n12
	matrix.elements[2] = n13
	matrix.elements[3] = n14
	matrix.elements[4] = n21
	matrix.elements[5] = n22
	matrix.elements[6] = n23
	matrix.elements[7] = n24
	matrix.elements[8] = n31
	matrix.elements[9] = n32
	matrix.elements[10] = n33
	matrix.elements[11] = n34
	matrix.elements[12] = n41
	matrix.elements[13] = n42
	matrix.elements[14] = n43
	matrix.elements[15] = n44
}

func (matrix *Matrix4) SetIdentity() {
	matrix.elements = [16]float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (matrix *Matrix4) Clone() *Matrix4 {
	m := Matrix4{
		elements: [16]float32{},
	}
	copy(m.elements[0:], matrix.elements[0:])
	return &m
}

func (matrix *Matrix4) Copy(other *Matrix4) {
	copy(matrix.elements[0:], other.elements[0:])
}

func (matrix *Matrix4) CopyPosition(other *Matrix4) {
	copy(matrix.elements[12:], other.elements[12:15])
}

func (matrix *Matrix4) ExtractBasis(xAxis *Vector4, yAxis *Vector4, zAxis *Vector4) {
	xAxis.SetFromMatrixColumn(matrix, 0)
	yAxis.SetFromMatrixColumn(matrix, 1)
	zAxis.SetFromMatrixColumn(matrix, 2)
}

func (matrix *Matrix4) MakeBasis(xAxis *Vector4, yAxis *Vector4, zAxis *Vector4) {
	matrix.Set(
		xAxis.X, yAxis.X, zAxis.X, 0,
		xAxis.Y, yAxis.Y, zAxis.Y, 0,
		xAxis.Z, yAxis.Z, zAxis.Z, 0,
		0, 0, 0, 0,
	)
}

func (matrix *Matrix4) ExtractRotation(m *Matrix4) {
	v1 := NewDefaultVector3()

	v1.SetFromMatrixColumn(m, 0)
	scaleX := 1 / v1.GetLength()

	v1.SetFromMatrixColumn(m, 1)
	scaleY := 1 / v1.GetLength()

	v1.SetFromMatrixColumn(m, 2)
	scaleZ := 1 / v1.GetLength()

	matrix.elements[0 ] = m.elements[0 ] * scaleX
	matrix.elements[1 ] = m.elements[1 ] * scaleX
	matrix.elements[2 ] = m.elements[2 ] * scaleX
	matrix.elements[3 ] = 0

	matrix.elements[4 ] = m.elements[4 ] * scaleY
	matrix.elements[5 ] = m.elements[5 ] * scaleY
	matrix.elements[6 ] = m.elements[6 ] * scaleY
	matrix.elements[7 ] = 0

	matrix.elements[8 ] = m.elements[8 ] * scaleZ
	matrix.elements[9 ] = m.elements[9 ] * scaleZ
	matrix.elements[10] = m.elements[10] * scaleZ
	matrix.elements[11] = 0

	matrix.elements[12] = 0
	matrix.elements[13] = 0
	matrix.elements[14] = 0
	matrix.elements[15] = 1
}

func (matrix *Matrix4) MakeRotationFromEuler(euler *Euler) {
	x := euler.x
	y := euler.y
	z := euler.z

	a := Cos(x)
	b := Sin(x)
	c := Cos(y)

	d := Sin(y)
	e := Cos(z)
	f := Sin(z)

	if euler.order == EulerOrderXYZ {
		ae := a * e
		af := a * f
		be := b * e
		bf := b * f

		matrix.elements[ 0 ] = c * e
		matrix.elements[ 4 ] = -c * f
		matrix.elements[ 8 ] = d

		matrix.elements[ 1 ] = af + be*d
		matrix.elements[ 5 ] = ae - bf*d
		matrix.elements[ 9 ] = -b * c

		matrix.elements[ 2 ] = bf - ae*d
		matrix.elements[ 6 ] = be + af*d
		matrix.elements[ 10 ] = a * c
	} else if euler.order == EulerOrderYXZ {
		ce := c * e
		cf := c * f
		de := d * e
		df := d * f

		matrix.elements[ 0 ] = ce + df*b
		matrix.elements[ 4 ] = de*b - cf
		matrix.elements[ 8 ] = a * d

		matrix.elements[ 1 ] = a * f
		matrix.elements[ 5 ] = a * e
		matrix.elements[ 9 ] = - b

		matrix.elements[ 2 ] = cf*b - de
		matrix.elements[ 6 ] = df + ce*b
		matrix.elements[ 10 ] = a * c
	} else if euler.order == EulerOrderZXY {
		ce := c * e
		cf := c * f
		de := d * e
		df := d * f

		matrix.elements[ 0 ] = ce - df*b
		matrix.elements[ 4 ] = -a * f
		matrix.elements[ 8 ] = de + cf*b

		matrix.elements[ 1 ] = cf + de*b
		matrix.elements[ 5 ] = a * e
		matrix.elements[ 9 ] = df - ce*b

		matrix.elements[ 2 ] = -a * d
		matrix.elements[ 6 ] = b
		matrix.elements[ 10 ] = a * c
	} else if euler.order == EulerOrderZYX {
		ae := a * e
		af := a * f
		be := b * e
		bf := b * f

		matrix.elements[ 0 ] = c * e
		matrix.elements[ 4 ] = be*d - af
		matrix.elements[ 8 ] = ae*d + bf

		matrix.elements[ 1 ] = c * f
		matrix.elements[ 5 ] = bf*d + ae
		matrix.elements[ 9 ] = af*d - be

		matrix.elements[ 2 ] = -d
		matrix.elements[ 6 ] = b * c
		matrix.elements[ 10 ] = a * c
	} else if euler.order == EulerOrderYZX {
		ac := a * c
		ad := a * d
		bc := b * c
		bd := b * d

		matrix.elements[ 0 ] = c * e
		matrix.elements[ 4 ] = bd - ac*f
		matrix.elements[ 8 ] = bc*f + ad

		matrix.elements[ 1 ] = f
		matrix.elements[ 5 ] = a * e
		matrix.elements[ 9 ] = -b * e

		matrix.elements[ 2 ] = -d * e
		matrix.elements[ 6 ] = ad*f + bc
		matrix.elements[ 10 ] = ac - bd*f
	} else if euler.order == EulerOrderXZY {
		ac := a * c
		ad := a * d
		bc := b * c
		bd := b * d

		matrix.elements[ 0 ] = c * e
		matrix.elements[ 4 ] = -f
		matrix.elements[ 8 ] = d * e

		matrix.elements[ 1 ] = ac*f + bd
		matrix.elements[ 5 ] = a * e
		matrix.elements[ 9 ] = ad*f - bc

		matrix.elements[ 2 ] = bc*f - ad
		matrix.elements[ 6 ] = b * e
		matrix.elements[ 10 ] = bd*f + ac
	}

	// bottom row
	matrix.elements[3 ] = 0
	matrix.elements[7 ] = 0
	matrix.elements[11] = 0

	// last column
	matrix.elements[12] = 0
	matrix.elements[13] = 0
	matrix.elements[14] = 0
	matrix.elements[15] = 1
}

func (matrix *Matrix4) MakeRotationFromQuaternion(q *Quaternion) {
	zero := NewVector3(0, 0, 0)
	one := NewVector3(1, 1, 1)
	matrix.Compose(zero, q, one)
}

func (matrix *Matrix4) LookAt(eye *Vector3, target *Vector3, up *Vector3) {
	x := NewDefaultVector3()
	y := NewDefaultVector3()
	z := NewDefaultVector3()

	z.SetSubVectors(eye, target)

	if z.GetLengthSq() == 0 {
		// eye and target are in the same position
		z.Z = 1
	}

	z.Normalize()
	x.CrossVectors(up, z)

	if x.GetLengthSq() == 0 {
		// up and z are parallel
		if Abs(up.Z) == 1 {
			z.X += 0.0001
		} else {
			z.Z += 0.0001
		}

		z.Normalize()
		x.CrossVectors(up, z)
	}

	x.Normalize()
	y.CrossVectors(z, x)
	matrix.elements[0 ] = x.X
	matrix.elements[4 ] = y.X
	matrix.elements[8 ] = z.X
	matrix.elements[1 ] = x.Y
	matrix.elements[5 ] = y.Y
	matrix.elements[9 ] = z.Y
	matrix.elements[2 ] = x.Z
	matrix.elements[6 ] = y.Z
	matrix.elements[10] = z.Z
}

func (matrix *Matrix4) Multiply(m *Matrix4) {
	matrix.MultiplyMatrices(matrix, m)
}

func (matrix *Matrix4) PreMultiply(m *Matrix4) {
	matrix.MultiplyMatrices(m, matrix)
}

func (matrix *Matrix4) MultiplyMatrices(a *Matrix4, b *Matrix4) {
	matrix.elements[0 ] = a.elements[0 ]*b.elements[0 ] + a.elements[4 ]*b.elements[1 ] + a.elements[8 ]*b.elements[2 ] + a.elements[12]*b.elements[3 ]
	matrix.elements[4 ] = a.elements[0 ]*b.elements[4 ] + a.elements[4 ]*b.elements[5 ] + a.elements[8 ]*b.elements[6 ] + a.elements[12]*b.elements[7 ]
	matrix.elements[8 ] = a.elements[0 ]*b.elements[8 ] + a.elements[4 ]*b.elements[9 ] + a.elements[8 ]*b.elements[10] + a.elements[12]*b.elements[11]
	matrix.elements[12] = a.elements[0 ]*b.elements[12] + a.elements[4 ]*b.elements[13] + a.elements[8 ]*b.elements[14] + a.elements[12]*b.elements[15]

	matrix.elements[1 ] = a.elements[1 ]*b.elements[0 ] + a.elements[5 ]*b.elements[1 ] + a.elements[9 ]*b.elements[2 ] + a.elements[13]*b.elements[3 ]
	matrix.elements[5 ] = a.elements[1 ]*b.elements[4 ] + a.elements[5 ]*b.elements[5 ] + a.elements[9 ]*b.elements[6 ] + a.elements[13]*b.elements[7 ]
	matrix.elements[9 ] = a.elements[1 ]*b.elements[8 ] + a.elements[5 ]*b.elements[9 ] + a.elements[9 ]*b.elements[10] + a.elements[13]*b.elements[11]
	matrix.elements[13] = a.elements[1 ]*b.elements[12] + a.elements[5 ]*b.elements[13] + a.elements[9 ]*b.elements[14] + a.elements[13]*b.elements[15]

	matrix.elements[2 ] = a.elements[2 ]*b.elements[0 ] + a.elements[6 ]*b.elements[1 ] + a.elements[10]*b.elements[2 ] + a.elements[14]*b.elements[3 ]
	matrix.elements[6 ] = a.elements[2 ]*b.elements[4 ] + a.elements[6 ]*b.elements[5 ] + a.elements[10]*b.elements[6 ] + a.elements[14]*b.elements[7 ]
	matrix.elements[10] = a.elements[2 ]*b.elements[8 ] + a.elements[6 ]*b.elements[9 ] + a.elements[10]*b.elements[10] + a.elements[14]*b.elements[11]
	matrix.elements[14] = a.elements[2 ]*b.elements[12] + a.elements[6 ]*b.elements[13] + a.elements[10]*b.elements[14] + a.elements[14]*b.elements[15]

	matrix.elements[3 ] = a.elements[3 ]*b.elements[0 ] + a.elements[7 ]*b.elements[1 ] + a.elements[11]*b.elements[2 ] + a.elements[15]*b.elements[3 ]
	matrix.elements[7 ] = a.elements[3 ]*b.elements[4 ] + a.elements[7 ]*b.elements[5 ] + a.elements[11]*b.elements[6 ] + a.elements[15]*b.elements[7 ]
	matrix.elements[11] = a.elements[3 ]*b.elements[8 ] + a.elements[7 ]*b.elements[9 ] + a.elements[11]*b.elements[10] + a.elements[15]*b.elements[11]
	matrix.elements[15] = a.elements[3 ]*b.elements[12] + a.elements[7 ]*b.elements[13] + a.elements[11]*b.elements[14] + a.elements[15]*b.elements[15]
}

func (matrix *Matrix4) MultiplyScalar(s float32) {
	matrix.elements[0 ] *= s
	matrix.elements[4 ] *= s
	matrix.elements[8 ] *= s
	matrix.elements[12] *= s

	matrix.elements[1 ] *= s
	matrix.elements[5 ] *= s
	matrix.elements[9 ] *= s
	matrix.elements[13] *= s

	matrix.elements[2 ] *= s
	matrix.elements[6 ] *= s
	matrix.elements[10] *= s
	matrix.elements[14] *= s

	matrix.elements[3 ] *= s
	matrix.elements[7 ] *= s
	matrix.elements[11] *= s
	matrix.elements[15] *= s
}

func (matrix *Matrix4) GetDeterminant() float32 {
	return matrix.elements[3 ]*(
		matrix.elements[12]*matrix.elements[9 ]*matrix.elements[6 ] -
			matrix.elements[8 ]*matrix.elements[13]*matrix.elements[6 ] -
			matrix.elements[12]*matrix.elements[5 ]*matrix.elements[10] +
			matrix.elements[4 ]*matrix.elements[13]*matrix.elements[10] +
			matrix.elements[8 ]*matrix.elements[5 ]*matrix.elements[14] -
			matrix.elements[4 ]*matrix.elements[9 ]*matrix.elements[14]) +
		matrix.elements[7 ]*(
			matrix.elements[0 ]*matrix.elements[9 ]*matrix.elements[14] -
				matrix.elements[0 ]*matrix.elements[13]*matrix.elements[10] +
				matrix.elements[12]*matrix.elements[1 ]*matrix.elements[10] -
				matrix.elements[8 ]*matrix.elements[1 ]*matrix.elements[14] +
				matrix.elements[8 ]*matrix.elements[13]*matrix.elements[2 ] -
				matrix.elements[12]*matrix.elements[9 ]*matrix.elements[2 ]) +
		matrix.elements[11]*(
			matrix.elements[0 ]*matrix.elements[13]*matrix.elements[6 ] -
				matrix.elements[0 ]*matrix.elements[5 ]*matrix.elements[14] -
				matrix.elements[12]*matrix.elements[1 ]*matrix.elements[6 ] +
				matrix.elements[4 ]*matrix.elements[1 ]*matrix.elements[14] +
				matrix.elements[12]*matrix.elements[5 ]*matrix.elements[2 ] -
				matrix.elements[4 ]*matrix.elements[13]*matrix.elements[2 ]) +
		matrix.elements[15]*(
			-matrix.elements[8 ]*matrix.elements[5 ]*matrix.elements[2 ] -
				matrix.elements[0 ]*matrix.elements[9 ]*matrix.elements[6 ] +
				matrix.elements[0 ]*matrix.elements[5 ]*matrix.elements[10] +
				matrix.elements[8 ]*matrix.elements[1 ]*matrix.elements[6 ] -
				matrix.elements[4 ]*matrix.elements[1 ]*matrix.elements[10] +
				matrix.elements[4 ]*matrix.elements[9 ]*matrix.elements[2 ])
}

func (matrix *Matrix4) Transpose() {
	tmp := matrix.elements[1 ]
	matrix.elements[1 ] = matrix.elements[4 ]
	matrix.elements[4 ] = tmp
	tmp = matrix.elements[2 ]
	matrix.elements[2 ] = matrix.elements[8 ]
	matrix.elements[8 ] = tmp
	tmp = matrix.elements[6 ]
	matrix.elements[6 ] = matrix.elements[9 ]
	matrix.elements[9 ] = tmp

	tmp = matrix.elements[3 ]
	matrix.elements[3 ] = matrix.elements[12]
	matrix.elements[12] = tmp
	tmp = matrix.elements[7 ]
	matrix.elements[7 ] = matrix.elements[13]
	matrix.elements[13] = tmp
	tmp = matrix.elements[11]
	matrix.elements[11] = matrix.elements[14]
	matrix.elements[14] = tmp
}

func (matrix *Matrix4) SetPosition(v *Vector3) {
	matrix.elements[12] = v.X
	matrix.elements[13] = v.Y
	matrix.elements[14] = v.Z
}

func (matrix *Matrix4) Inverse() error {
	return matrix.SetInverseOf(matrix.Clone(), true)
}

// Get the inverse of the matrix [m] and set its value to the current matrix
func (matrix *Matrix4) SetInverseOf(m *Matrix4, errorOnDegenerate bool) error {
	t11 := m.elements[9 ]*m.elements[14]*m.elements[7 ] - m.elements[13]*m.elements[10]*m.elements[7 ] + m.elements[13]*m.elements[6 ]*m.elements[11] - m.elements[5 ]*m.elements[14]*m.elements[11] - m.elements[9 ]*m.elements[6 ]*m.elements[15] + m.elements[5 ]*m.elements[10]*m.elements[15]
	t12 := m.elements[12]*m.elements[10]*m.elements[7 ] - m.elements[8 ]*m.elements[14]*m.elements[7 ] - m.elements[12]*m.elements[6 ]*m.elements[11] + m.elements[4 ]*m.elements[14]*m.elements[11] + m.elements[8 ]*m.elements[6 ]*m.elements[15] - m.elements[4 ]*m.elements[10]*m.elements[15]
	t13 := m.elements[8 ]*m.elements[13]*m.elements[7 ] - m.elements[12]*m.elements[9 ]*m.elements[7 ] + m.elements[12]*m.elements[5 ]*m.elements[11] - m.elements[4 ]*m.elements[13]*m.elements[11] - m.elements[8 ]*m.elements[5 ]*m.elements[15] + m.elements[4 ]*m.elements[9 ]*m.elements[15]
	t14 := m.elements[12]*m.elements[9 ]*m.elements[6 ] - m.elements[8 ]*m.elements[13]*m.elements[6 ] - m.elements[12]*m.elements[5 ]*m.elements[10] + m.elements[4 ]*m.elements[13]*m.elements[10] + m.elements[8 ]*m.elements[5 ]*m.elements[14] - m.elements[4 ]*m.elements[9 ]*m.elements[14]

	det := m.elements[0 ]*t11 + m.elements[1 ]*t12 + m.elements[2 ]*t13 + m.elements[3 ]*t14

	if det == 0 {
		if errorOnDegenerate == true {
			return errors.New(".getInverse() can't invert matrix, determinant is 0")
		} else {
			logger.Warning(".getInverse() can't invert matrix, determinant is 0")
			matrix.SetIdentity()
			return nil
		}
	}

	detInv := 1 / det

	matrix.elements[0 ] = t11 * detInv
	matrix.elements[1 ] = (m.elements[13]*m.elements[10]*m.elements[3 ] - m.elements[9 ]*m.elements[14]*m.elements[3 ] - m.elements[13]*m.elements[2 ]*m.elements[11] + m.elements[1 ]*m.elements[14]*m.elements[11] + m.elements[9 ]*m.elements[2 ]*m.elements[15] - m.elements[1 ]*m.elements[10]*m.elements[15]) * detInv
	matrix.elements[2 ] = (m.elements[5 ]*m.elements[14]*m.elements[3 ] - m.elements[13]*m.elements[6 ]*m.elements[3 ] + m.elements[13]*m.elements[2 ]*m.elements[7 ] - m.elements[1 ]*m.elements[14]*m.elements[7 ] - m.elements[5 ]*m.elements[2 ]*m.elements[15] + m.elements[1 ]*m.elements[6 ]*m.elements[15]) * detInv
	matrix.elements[3 ] = (m.elements[9 ]*m.elements[6 ]*m.elements[3 ] - m.elements[5 ]*m.elements[10]*m.elements[3 ] - m.elements[9 ]*m.elements[2 ]*m.elements[7 ] + m.elements[1 ]*m.elements[10]*m.elements[7 ] + m.elements[5 ]*m.elements[2 ]*m.elements[11] - m.elements[1 ]*m.elements[6 ]*m.elements[11]) * detInv

	matrix.elements[4 ] = t12 * detInv
	matrix.elements[5 ] = (m.elements[8 ]*m.elements[14]*m.elements[3 ] - m.elements[12]*m.elements[10]*m.elements[3 ] + m.elements[12]*m.elements[2 ]*m.elements[11] - m.elements[0 ]*m.elements[14]*m.elements[11] - m.elements[8 ]*m.elements[2 ]*m.elements[15] + m.elements[0 ]*m.elements[10]*m.elements[15]) * detInv
	matrix.elements[6 ] = (m.elements[12]*m.elements[6 ]*m.elements[3 ] - m.elements[4 ]*m.elements[14]*m.elements[3 ] - m.elements[12]*m.elements[2 ]*m.elements[7 ] + m.elements[0 ]*m.elements[14]*m.elements[7 ] + m.elements[4 ]*m.elements[2 ]*m.elements[15] - m.elements[0 ]*m.elements[6 ]*m.elements[15]) * detInv
	matrix.elements[7 ] = (m.elements[4 ]*m.elements[10]*m.elements[3 ] - m.elements[8 ]*m.elements[6 ]*m.elements[3 ] + m.elements[8 ]*m.elements[2 ]*m.elements[7 ] - m.elements[0 ]*m.elements[10]*m.elements[7 ] - m.elements[4 ]*m.elements[2 ]*m.elements[11] + m.elements[0 ]*m.elements[6 ]*m.elements[11]) * detInv

	matrix.elements[8 ] = t13 * detInv
	matrix.elements[9 ] = (m.elements[12]*m.elements[9 ]*m.elements[3 ] - m.elements[8 ]*m.elements[13]*m.elements[3 ] - m.elements[12]*m.elements[1 ]*m.elements[11] + m.elements[0 ]*m.elements[13]*m.elements[11] + m.elements[8 ]*m.elements[1 ]*m.elements[15] - m.elements[0 ]*m.elements[9 ]*m.elements[15]) * detInv
	matrix.elements[10] = (m.elements[4 ]*m.elements[13]*m.elements[3 ] - m.elements[12]*m.elements[5 ]*m.elements[3 ] + m.elements[12]*m.elements[1 ]*m.elements[7 ] - m.elements[0 ]*m.elements[13]*m.elements[7 ] - m.elements[4 ]*m.elements[1 ]*m.elements[15] + m.elements[0 ]*m.elements[5 ]*m.elements[15]) * detInv
	matrix.elements[11] = (m.elements[8 ]*m.elements[5 ]*m.elements[3 ] - m.elements[4 ]*m.elements[9 ]*m.elements[3 ] - m.elements[8 ]*m.elements[1 ]*m.elements[7 ] + m.elements[0 ]*m.elements[9 ]*m.elements[7 ] + m.elements[4 ]*m.elements[1 ]*m.elements[11] - m.elements[0 ]*m.elements[5 ]*m.elements[11]) * detInv

	matrix.elements[12] = t14 * detInv
	matrix.elements[13] = (m.elements[8 ]*m.elements[13]*m.elements[2 ] - m.elements[12]*m.elements[9 ]*m.elements[2 ] + m.elements[12]*m.elements[1 ]*m.elements[10] - m.elements[0 ]*m.elements[13]*m.elements[10] - m.elements[8 ]*m.elements[1 ]*m.elements[14] + m.elements[0 ]*m.elements[9 ]*m.elements[14]) * detInv
	matrix.elements[14] = (m.elements[12]*m.elements[5 ]*m.elements[2 ] - m.elements[4 ]*m.elements[13]*m.elements[2 ] - m.elements[12]*m.elements[1 ]*m.elements[6 ] + m.elements[0 ]*m.elements[13]*m.elements[6 ] + m.elements[4 ]*m.elements[1 ]*m.elements[14] - m.elements[0 ]*m.elements[5 ]*m.elements[14]) * detInv
	matrix.elements[15] = (m.elements[4 ]*m.elements[9 ]*m.elements[2 ] - m.elements[8 ]*m.elements[5 ]*m.elements[2 ] + m.elements[8 ]*m.elements[1 ]*m.elements[6 ] - m.elements[0 ]*m.elements[9 ]*m.elements[6 ] - m.elements[4 ]*m.elements[1 ]*m.elements[10] + m.elements[0 ]*m.elements[5 ]*m.elements[10]) * detInv

	return nil
}

func (matrix *Matrix4) Scale(scale *Vector3) {
	matrix.elements[0 ] *= scale.X
	matrix.elements[1 ] *= scale.X
	matrix.elements[2 ] *= scale.X
	matrix.elements[3 ] *= scale.X
	matrix.elements[4 ] *= scale.Y
	matrix.elements[5 ] *= scale.Y
	matrix.elements[6 ] *= scale.Y
	matrix.elements[7 ] *= scale.Y
	matrix.elements[8 ] *= scale.Z
	matrix.elements[9 ] *= scale.Z
	matrix.elements[10] *= scale.Z
	matrix.elements[11] *= scale.Z
}

func (matrix *Matrix4) GetMaxScaleOnAxis() float32 {
	scaleXSq := matrix.elements[0 ]*matrix.elements[0 ] + matrix.elements[1 ]*matrix.elements[1 ] + matrix.elements[2 ]*matrix.elements[2 ]
	scaleYSq := matrix.elements[4 ]*matrix.elements[4 ] + matrix.elements[5 ]*matrix.elements[5 ] + matrix.elements[6 ]*matrix.elements[6 ]
	scaleZSq := matrix.elements[8 ]*matrix.elements[8 ] + matrix.elements[9 ]*matrix.elements[9 ] + matrix.elements[10]*matrix.elements[10]
	return Sqrt(Max(scaleXSq, Max(scaleYSq, scaleZSq)))
}

func (matrix *Matrix4) Compose(position *Vector3, q *Quaternion, scale *Vector3) {
	x := q.GetX()
	y := q.GetY()
	z := q.GetZ()
	w := q.GetW()

	x2 := x + x
	y2 := y + y
	z2 := z + z
	xx := x * x2
	xy := x * y2
	xz := x * z2
	yy := y * y2
	yz := y * z2
	zz := z * z2
	wx := w * x2
	wy := w * y2
	wz := w * z2

	matrix.elements[0 ] = (1 - (yy + zz)) * scale.X
	matrix.elements[1 ] = (xy + wz) * scale.X
	matrix.elements[2 ] = (xz - wy) * scale.X
	matrix.elements[3 ] = 0

	matrix.elements[4 ] = (xy - wz) * scale.Y
	matrix.elements[5 ] = (1 - (xx + zz)) * scale.Y
	matrix.elements[6 ] = (yz + wx) * scale.Y
	matrix.elements[7 ] = 0

	matrix.elements[8 ] = (xz + wy) * scale.Z
	matrix.elements[9 ] = (yz - wx) * scale.Z
	matrix.elements[10] = (1 - (xx + yy)) * scale.Z
	matrix.elements[11] = 0

	matrix.elements[12] = position.X
	matrix.elements[13] = position.Y
	matrix.elements[14] = position.Z
	matrix.elements[15] = 1
}

func (matrix *Matrix4) DeCompose(position *Vector3, q *Quaternion, scale *Vector3) {
	tmpVector := NewDefaultVector3()
	tmpMatrix := matrix.Clone()

	tmpVector.Set(matrix.elements[0 ], matrix.elements[1 ], matrix.elements[2 ])
	scale.X = tmpVector.GetLength()
	tmpVector.Set(matrix.elements[4 ], matrix.elements[5 ], matrix.elements[6 ])
	scale.Y = tmpVector.GetLength()
	tmpVector.Set(matrix.elements[8 ], matrix.elements[9 ], matrix.elements[10])
	scale.Z = tmpVector.GetLength()

	// if determine is negative, we need to invert one scale
	if matrix.GetDeterminant() < 0 {
		scale.X = -scale.X
	}

	position.X = matrix.elements[12]
	position.Y = matrix.elements[13]
	position.Z = matrix.elements[14]

	// scale the rotation part
	var invSX = 1 / scale.X
	var invSY = 1 / scale.Y
	var invSZ = 1 / scale.Z

	tmpMatrix.elements[0 ] *= invSX
	tmpMatrix.elements[1 ] *= invSX
	tmpMatrix.elements[2 ] *= invSX

	tmpMatrix.elements[4 ] *= invSY
	tmpMatrix.elements[5 ] *= invSY
	tmpMatrix.elements[6 ] *= invSY

	tmpMatrix.elements[8 ] *= invSZ
	tmpMatrix.elements[9 ] *= invSZ
	tmpMatrix.elements[10] *= invSZ

	q.SetFromRotationMatrix(tmpMatrix)
}

func (matrix *Matrix4) Equals(ma *Matrix4) bool {
	for ind := range matrix.elements {
		if matrix.elements[ind] != ma.elements[ind] {
			return false
		}
	}
	return true
}

func (matrix *Matrix4) EqualsRound(ma *Matrix4, decimals float32) bool {
	mul := Pow(10, decimals)
	for ind := range matrix.elements {
		if Round(mul*matrix.elements[ind])/mul != Round(mul*ma.elements[ind])/mul {
			return false
		}
	}
	return true
}

func (matrix *Matrix4) ToString() string {
	return fmt.Sprintf("%9.2f %9.2f %9.2f %9.2f\n%9.2f %9.2f %9.2f %9.2f\n%9.2f %9.2f %9.2f %9.2f\n%9.2f %9.2f %9.2f %9.2f",
		matrix.elements[0 ], matrix.elements[1 ], matrix.elements[2 ], matrix.elements[3 ],
		matrix.elements[4 ], matrix.elements[5 ], matrix.elements[6 ], matrix.elements[7 ],
		matrix.elements[8 ], matrix.elements[9 ], matrix.elements[10], matrix.elements[11],
		matrix.elements[12], matrix.elements[13], matrix.elements[14], matrix.elements[15])
}

func (matrix *Matrix4) ToArray() [16]float32 {
	return matrix.Clone().elements
}
