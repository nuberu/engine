package math

import (
	"github.com/tokkenno/seed/event"
	"math"
)

type Quaternion struct {
	x float32
	y float32
	z float32
	w float32

	changeEvent *event.Emitter
}

func NewDefaultQuaternion() *Quaternion {
	return NewQuaternion(0, 0, 0, 1)
}

func NewQuaternion(x float32, y float32, z float32, w float32) *Quaternion {
	return &Quaternion{
		x:           x,
		y:           y,
		z:           z,
		w:           w,
		changeEvent: event.NewEvent(),
	}
}

func NewQuaternionFromArray(arr []float32, offset int) *Quaternion {
	return NewQuaternion(
		arr[offset],
		arr[offset+1],
		arr[offset+2],
		arr[offset+2],
	)
}

func (qua *Quaternion) OnChange() *event.Handler {
	return qua.changeEvent.GetHandler()
}

func (qua *Quaternion) Set(x float32, y float32, z float32, w float32) {
	qua.x = x
	qua.y = y
	qua.z = z
	qua.w = w
	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) GetX() float32 {
	return qua.x
}

func (qua *Quaternion) SetX(x float32) {
	qua.x = x
	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) GetY() float32 {
	return qua.y
}

func (qua *Quaternion) SetY(y float32) {
	qua.y = y
	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) GetZ() float32 {
	return qua.z
}

func (qua *Quaternion) SetZ(z float32) {
	qua.z = z
	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) GetW() float32 {
	return qua.w
}

func (qua *Quaternion) SetW(w float32) {
	qua.w = w
	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) Clone() *Quaternion {
	return &Quaternion{
		x: qua.x,
		y: qua.y,
		z: qua.z,
		w: qua.w,
	}
}

func (qua *Quaternion) Copy(q *Quaternion) {
	qua.w = q.w
	qua.w = q.w
	qua.w = q.w
	qua.w = q.w
	qua.changeEvent.Emit(qua, nil)
}

/*
	http://www.mathworks.com/matlabcentral/fileexchange/
	20696-function-to-convert-between-dcm-euler-angles-quaternions-and-euler-vectors/
	content/SpinCalc.m
 */
func (qua *Quaternion) SetFromEuler(euler *Euler, update bool) {
	x := euler.x
	y := euler.y
	z := euler.z
	order := euler.order

	c1 := math.Cos(x / 2)
	c2 := math.Cos(y / 2)
	c3 := math.Cos(z / 2)

	s1 := math.Cos(x / 2)
	s2 := math.Cos(y / 2)
	s3 := math.Cos(z / 2)

	switch order {
	case EulerOrderXYZ:
		qua.x = s1*c2*c3 + c1*s2*s3
		qua.y = c1*s2*c3 - s1*c2*s3
		qua.z = c1*c2*s3 + s1*s2*c3
		qua.w = c1*c2*c3 - s1*s2*s3
		break
	case EulerOrderYXZ:
		qua.x = s1*c2*c3 + c1*s2*s3
		qua.y = c1*s2*c3 - s1*c2*s3
		qua.z = c1*c2*s3 - s1*s2*c3
		qua.w = c1*c2*c3 + s1*s2*s3
		break
	case EulerOrderZXY:
		qua.x = s1*c2*c3 - c1*s2*s3
		qua.y = c1*s2*c3 + s1*c2*s3
		qua.z = c1*c2*s3 + s1*s2*c3
		qua.w = c1*c2*c3 - s1*s2*s3
		break
	case EulerOrderZYX:
		qua.x = s1*c2*c3 - c1*s2*s3
		qua.y = c1*s2*c3 + s1*c2*s3
		qua.z = c1*c2*s3 - s1*s2*c3
		qua.w = c1*c2*c3 + s1*s2*s3
		break
	case EulerOrderYZX:
		qua.x = s1*c2*c3 + c1*s2*s3
		qua.y = c1*s2*c3 + s1*c2*s3
		qua.z = c1*c2*s3 - s1*s2*c3
		qua.w = c1*c2*c3 - s1*s2*s3
		break
	case EulerOrderXZY:
		qua.x = s1*c2*c3 - c1*s2*s3
		qua.y = c1*s2*c3 - s1*c2*s3
		qua.z = c1*c2*s3 + s1*s2*c3
		qua.w = c1*c2*c3 + s1*s2*s3
		break
	}

	if update != false {
		qua.changeEvent.Emit(qua, nil)
	}
}

func (qua *Quaternion) SetFromAxisAngle(axis *Vector3, angle Angle) {
	halfAngle := angle / 2
	s := math.Sin(float32(halfAngle))

	qua.x = axis.X * s
	qua.y = axis.Y * s
	qua.z = axis.Z * s
	qua.w = math.Cos(float32(halfAngle))

	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) SetFromRotationMatrix(m *Matrix4) {
	m11 := m.elements[0 ]
	m12 := m.elements[4 ]
	m13 := m.elements[8 ]
	m21 := m.elements[1 ]
	m22 := m.elements[5 ]
	m23 := m.elements[9 ]
	m31 := m.elements[2 ]
	m32 := m.elements[6 ]
	m33 := m.elements[10]

	trace := m11 + m22 + m33
	var s float32

	if trace > 0 {
		s = 0.5 / math.Sqrt(trace+1.0)

		qua.w = 0.25 / s
		qua.x = (m32 - m23) * s
		qua.y = (m13 - m31) * s
		qua.z = (m21 - m12) * s
	} else if m11 > m22 && m11 > m33 {
		s = 2.0 * math.Sqrt(1.0+m11-m22-m33)

		qua.w = (m32 - m23) / s
		qua.x = 0.25 * s
		qua.y = (m12 + m21) / s
		qua.z = (m13 + m31) / s
	} else if m22 > m33 {
		s = 2.0 * math.Sqrt(1.0+m22-m11-m33)

		qua.w = (m13 - m31) / s
		qua.x = (m12 + m21) / s
		qua.y = 0.25 * s
		qua.z = (m23 + m32) / s
	} else {
		s = 2.0 * math.Sqrt(1.0+m33-m11-m22)

		qua.w = (m21 - m12) / s
		qua.x = (m13 + m31) / s
		qua.y = (m23 + m32) / s
		qua.z = 0.25 * s
	}

	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) SetFromUnitVectors(vFrom *Vector3, vTo *Vector3) {
	const EPS float32 = 0.000001
	var r float32
	v1 := NewDefaultVector3()

	r = vFrom.Dot(vTo) + 1

	if r < EPS {
		r = 0

		if math.Abs(vFrom.X) > math.Abs(vFrom.Z) {
			v1.Set(-vFrom.Y, vFrom.X, 0)
		} else {
			v1.Set(0, -vFrom.Z, vFrom.Y)
		}
	} else {
		v1.CrossVectors(vFrom, vTo)
	}

	qua.x = v1.X
	qua.y = v1.Y
	qua.z = v1.Z
	qua.w = r

	qua.Normalize()
}

func (qua *Quaternion) AngleTo(q *Quaternion) float32 {
	return 2 * math.Acos(math.Abs(Clamp(qua.Dot(q), - 1, 1)))
}

func (qua *Quaternion) RotateTowards(q *Quaternion, step float32) {
	angle := qua.AngleTo(q)
	if angle == 0 {
		return
	}
	t := math.Min(1, step/angle)
	qua.Slerp(q, t)
}

func (qua *Quaternion) Inverse() {
	qua.Conjugate()
}

func (qua *Quaternion) Conjugate() {
	qua.x *= -1
	qua.y *= -1
	qua.z *= -1

	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) Dot(v *Quaternion) float32 {
	return qua.x*v.x + qua.y*v.y + qua.z*v.z + qua.w*v.w
}

func (qua *Quaternion) GetLengthSq() float32 {
	return qua.x*qua.x + qua.y*qua.y + qua.z*qua.z + qua.w*qua.w
}

func (qua *Quaternion) GetLength() float32 {
	return math.Sqrt(qua.GetLengthSq())
}

func (qua *Quaternion) Normalize() {
	l := qua.GetLength()

	if l == 0 {
		qua.x = 0;
		qua.y = 0;
		qua.z = 0;
		qua.w = 1;
	} else {
		l = 1 / l

		qua.x = qua.x * l
		qua.y = qua.y * l
		qua.z = qua.z * l
		qua.w = qua.w * l
	}

	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) Multiply(q *Quaternion) {
	qua.MultiplyQuaternions(qua, q)
}

func (qua *Quaternion) PreMultiply(q *Quaternion) {
	qua.MultiplyQuaternions(q, qua)
}

func (qua *Quaternion) MultiplyQuaternions(a *Quaternion, b *Quaternion) {
	qua.x = a.x*b.w + a.w*b.x + a.y*b.z - a.z*b.y
	qua.y = a.y*b.w + a.w*b.y + a.z*b.x - a.x*b.z
	qua.z = a.z*b.w + a.w*b.z + a.x*b.y - a.y*b.x
	qua.w = a.w*b.w - a.x*b.x - a.y*b.y - a.z*b.z

	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) Slerp(qb *Quaternion, t float32) {
	if t == 0 {
		return
	} else if t == 1 {
		qua.Copy(qb)
		return
	}

	x := qua.x
	y := qua.y
	z := qua.z
	w := qua.w

	// http://www.euclideanspace.com/maths/algebra/realNormedAlgebra/quaternions/slerp/

	var cosHalfTheta = w*qb.w + x*qb.x + y*qb.y + z*qb.z

	if cosHalfTheta < 0 {
		qua.w = - qb.w
		qua.x = - qb.x
		qua.y = - qb.y
		qua.z = - qb.z

		cosHalfTheta = - cosHalfTheta;
	} else {
		qua.Copy(qb)
	}

	if cosHalfTheta >= 1.0 {
		qua.w = w
		qua.x = x
		qua.y = y
		qua.z = z

		return
	}

	var sqrSinHalfTheta = 1.0 - cosHalfTheta*cosHalfTheta

	if sqrSinHalfTheta <= math.SmallestNonzerofloat32 {
		s := 1 - t
		qua.w = s*w + t*qua.w
		qua.x = s*x + t*qua.x
		qua.y = s*y + t*qua.y
		qua.z = s*z + t*qua.z
		qua.Normalize()
		return
	}

	sinHalfTheta := math.Sqrt(sqrSinHalfTheta)
	halfTheta := math.Atan2(sinHalfTheta, cosHalfTheta)
	ratioA := math.Sin((1-t)*halfTheta) / sinHalfTheta
	ratioB := math.Sin(t*halfTheta) / sinHalfTheta

	qua.w = w*ratioA + qua.w*ratioB
	qua.x = x*ratioA + qua.x*ratioB
	qua.y = y*ratioA + qua.y*ratioB
	qua.z = z*ratioA + qua.z*ratioB

	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) Equals(q *Quaternion) bool {
	return qua.x == q.x && qua.y == q.y && qua.z == q.z
}

func (qua *Quaternion) ToArray() [4]float32 {
	return [4]float32{qua.x, qua.y, qua.z, qua.w}
}

func SlerpQuaternion(qa *Quaternion, qb *Quaternion, qm *Quaternion, t float32) {
	qm.Copy(qa)
	qm.Slerp(qb, t)
}

/*
 fuzz-free, array-based Quaternion SLERP operation
 */
func SlerpFlatQuaternion(dst []float32, dstOffset int, src0 []float32, srcOffset0 int, src1 []float32, srcOffset1 int, t float32) {
	x0 := src0[srcOffset0+0]
	y0 := src0[srcOffset0+1]
	z0 := src0[srcOffset0+2]
	w0 := src0[srcOffset0+3]

	x1 := src1[srcOffset1+0]
	y1 := src1[srcOffset1+1]
	z1 := src1[srcOffset1+2]
	w1 := src1[srcOffset1+3]

	if w0 != w1 || x0 != x1 || y0 != y1 || z0 != z1 {
		s := 1 - t
		cos := x0*x1 + y0*y1 + z0*z1 + w0*w1
		dir := float32(1)
		if cos < 0 {
			dir = -1
		}
		sqrSin := 1 - cos*cos

		// Skip the Slerp for tiny steps to avoid numeric problems:
		if sqrSin > math.SmallestNonzerofloat32 {
			sin := math.Sqrt(sqrSin)
			len := math.Atan2(sin, cos*dir)

			s = math.Sin(s*len) / sin
			t = math.Sin(t*len) / sin
		}

		tDir := t * dir

		x0 = x0*s + x1*tDir
		y0 = y0*s + y1*tDir
		z0 = z0*s + z1*tDir
		w0 = w0*s + w1*tDir

		// Normalize in case we just did a lerp:
		if s == 1-t {

			f := 1 / math.Sqrt(x0*x0+y0*y0+z0*z0+w0*w0)

			x0 *= f
			y0 *= f
			z0 *= f
			w0 *= f
		}
	}

	dst[dstOffset] = x0
	dst[dstOffset+1] = y0
	dst[dstOffset+2] = z0
	dst[dstOffset+3] = w0
}
