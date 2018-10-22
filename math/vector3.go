package math

import (
	"github.com/tokkenno/seed/core/types"
	nativeMath "math"
)

type Vector3 struct {
	Vector2
	Z float32
}

func NewDefaultVector3() *Vector3 {
	return NewVector3(0, 0, 0)
}

func NewVector3(x float32, y float32, z float32) *Vector3 {
	return &Vector3{
		Vector2: *NewVector2(x, y),
		Z:       z,
	}
}

func NewVector3Inf(sign int) *Vector3 {
	return &Vector3{
		Vector2: *NewVector2Inf(sign),
		Z: float32(nativeMath.Inf(sign)),
	}
}

func NewVector3FromArray(arr []float32, offset int) *Vector3 {
	return NewVector3(
		arr[offset],
		arr[offset+1],
		arr[offset+2],
	)
}

func (vec *Vector3) Set(x float32, y float32, z float32) {
	vec.X = x
	vec.Y = y
	vec.Z = z
}

func (vec *Vector3) SetZ(z float32) {
	vec.Z = z
}

func (vec *Vector3) SetScalar(num float32) {
	vec.X = num
	vec.Y = num
	vec.Z = num
}

func (vec *Vector3) SetFromSpherical(s *Spherical) {
	vec.SetFromSphericalCoordinates(s.radius, s.phi, s.theta)
}

func (vec *Vector3) SetFromSphericalCoordinates(radius float32, phi float32, theta float32) {
	sinPhiRadius := Sin(phi) * radius
	vec.X = sinPhiRadius * Sin(theta)
	vec.Y = Cos(phi) * radius
	vec.Z = sinPhiRadius * Cos(theta)
}

func (vec *Vector3) SetFromCylindrical(c *Cylindrical) {
	vec.SetFromCylindricalCoordinates(c.radius, c.theta, c.y)
}

func (vec *Vector3) SetFromCylindricalCoordinates(radius float32, theta float32, y float32) {
	vec.X = radius * Sin(theta)
	vec.Y = y
	vec.Z = radius * Cos(theta)
}

func (vec *Vector3) SetFromMatrixPosition(m *Matrix4) {
	vec.X = m.elements[12]
	vec.Y = m.elements[13]
	vec.Z = m.elements[14]
}

func (vec *Vector3) SetFromMatrixScale(m *Matrix4) {
	vec.SetFromMatrixColumn(m, 0)
	sx := vec.GetLength()
	vec.SetFromMatrixColumn(m, 1)
	sy := vec.GetLength()
	vec.SetFromMatrixColumn(m, 2)
	sz := vec.GetLength()

	vec.X = sx
	vec.Y = sy
	vec.Z = sz
}

func (vec *Vector3) SetFromMatrixColumn(m *Matrix4, col int) {
	offset := col * 4
	vec.X = m.elements[offset]
	vec.Y = m.elements[offset+1]
	vec.Z = m.elements[offset+2]
}

func (vec *Vector3) Copy(v *Vector3) {
	vec.X = v.X
	vec.Y = v.Y
	vec.Z = v.Z
}

func (vec *Vector3) Clone() *Vector3 {
	return &Vector3{
		Vector2: *vec.Vector2.Clone(),
		Z:       vec.Z,
	}
}

func (vec *Vector3) Add(a *Vector3) {
	vec.X += a.X
	vec.Y += a.Y
	vec.Z += a.Z
}

func (vec *Vector3) AddComponents(x float32, y float32, z float32) {
	vec.X += x
	vec.Y += y
	vec.Z += z
}

func (vec *Vector3) AddScalar(num float32) {
	vec.X += num
	vec.Y += num
	vec.Z += num
}

func (vec *Vector3) SetAddVectors(v1 *Vector3, v2 *Vector3) {
	vec.X = v1.X + v2.X
	vec.Y = v1.Y + v2.Y
	vec.Z = v1.Z + v2.Z
}

func (vec *Vector3) AddScaledVector(v1 *Vector3, scale float32) {
	vec.X += v1.X * scale
	vec.Y += v1.Y * scale
	vec.Z += v1.Z * scale
}

func (vec *Vector3) Sub(v *Vector3) {
	vec.X -= v.X
	vec.Y -= v.Y
	vec.Z -= v.Z
}

func (vec *Vector3) SubScalar(num float32) {
	vec.X -= num
	vec.Y -= num
	vec.Z -= num
}

func (vec *Vector3) SetSubVectors(v1 *Vector3, v2 *Vector3) {
	vec.X = v1.X - v2.X
	vec.Y = v1.Y - v2.Y
	vec.Z = v1.Z - v2.Z
}

func (vec *Vector3) Multiply(v *Vector3) {
	vec.X *= v.X
	vec.Y *= v.Y
	vec.Z *= v.Z
}

func (vec *Vector3) MultiplyScalar(num float32) {
	vec.X *= num
	vec.Y *= num
	vec.Z *= num
}

func (vec *Vector3) Divide(v *Vector3) {
	vec.X /= v.X
	vec.Y /= v.Y
	vec.Z /= v.Z
}

func (vec *Vector3) DivideScalar(num float32) {
	vec.MultiplyScalar(1 / num)
}

func (vec *Vector3) ApplyEuler(euler *Euler) {
	quaternion := NewDefaultQuaternion()
	quaternion.SetFromEuler(euler, false)
	vec.ApplyQuaternion(quaternion)
}

func (vec *Vector3) ApplyAxisAngle(axis *Vector3, angle types.Angle) {
	quaternion := NewDefaultQuaternion()
	quaternion.SetFromAxisAngle(axis, angle)
	vec.ApplyQuaternion(quaternion)
}

func (vec *Vector3) ApplyMatrix3(m *Matrix3) {
	x := vec.X
	y := vec.Y
	z := vec.Z
	e := m.GetElements()

	vec.X = e[0]*x + e[3]*y + e[6]*z
	vec.Y = e[1]*x + e[4]*y + e[7]*z
	vec.Z = e[2]*x + e[5]*y + e[8]*z
}

func (vec *Vector3) ApplyMatrix4(m *Matrix4) {
	x := vec.X
	y := vec.Y
	z := vec.Z
	e := m.GetElements()
	w := 1 / (e[3]*x + e[7]*y + e[11]*z + e[15])

	vec.X = (e[0]*x + e[4]*y + e[8]*z + e[12]) * w
	vec.Y = (e[1]*x + e[5]*y + e[9]*z + e[13]) * w
	vec.Z = (e[2]*x + e[6]*y + e[10]*z + e[14]) * w
}

func (vec *Vector3) ApplyQuaternion(q *Quaternion) {
	x := vec.X
	y := vec.Y
	z := vec.Z

	qx := q.GetX()
	qy := q.GetY()
	qz := q.GetZ()
	qw := q.GetW()

	// calculate quaternion * vector
	ix := qw*x + qy*z - qz*y
	iy := qw*y + qz*x - qx*z
	iz := qw*z + qx*y - qy*x
	iw := - qx*x - qy*y - qz*z

	// calculate result * inverse quaternion
	vec.X = ix*qw + iw*- qx + iy*- qz - iz*- qy
	vec.Y = iy*qw + iw*- qy + iz*- qx - ix*- qz
	vec.Z = iz*qw + iw*- qz + ix*- qy - iy*- qx
}

// input: THREE.Matrix4 affine matrix
// vector interpreted as a direction
func (vec *Vector3) TransformDirection(matrix *Matrix4) {
	x := vec.X
	y := vec.Y
	z := vec.Z

	vec.X = matrix.elements[0]*x + matrix.elements[4]*y + matrix.elements[8]*z
	vec.Y = matrix.elements[1]*x + matrix.elements[5]*y + matrix.elements[9]*z
	vec.Z = matrix.elements[2]*x + matrix.elements[6]*y + matrix.elements[10]*z

	vec.Normalize()
}

func (vec *Vector3) Min(v *Vector3) {
	vec.X = Min(vec.X, v.X)
	vec.Y = Min(vec.Y, v.Y)
	vec.Z = Min(vec.Z, v.Z)
}

func (vec *Vector3) Max(v *Vector3) {
	vec.X = Max(vec.X, v.X)
	vec.Y = Max(vec.Y, v.Y)
	vec.Z = Max(vec.Z, v.Z)
}

// Clamps the value to be between min and max.
func (vec *Vector3) Clamp(min *Vector3, max *Vector3) {
	vec.X = Max(min.X, Min(max.X, vec.X))
	vec.Y = Max(min.Y, Min(max.Y, vec.Y))
	vec.Z = Max(min.Z, Min(max.Z, vec.Z))
}

func (vec *Vector3) ClampScalar(min float32, max float32) {
	minVec := NewVector3(min, min, min)
	maxVec := NewVector3(max, max, max)
	vec.Clamp(minVec, maxVec)
}

func (vec *Vector3) ClampLength(min float32, max float32) {
	length := vec.GetLength()

	div := length
	if length == 0 {
		div = 1
	}

	vec.DivideScalar(div)
	vec.MultiplyScalar(Max(min, Min(max, length)))
}

func (vec *Vector3) Floor() {
	vec.X = Floor(vec.X)
	vec.Y = Floor(vec.Y)
	vec.Z = Floor(vec.Z)
}

func (vec *Vector3) Ceil() {
	vec.X = Ceil(vec.X)
	vec.Y = Ceil(vec.Y)
	vec.Z = Ceil(vec.Z)
}

func (vec *Vector3) Round() {
	vec.X = Round(vec.X)
	vec.Y = Round(vec.Y)
	vec.Z = Round(vec.Z)
}

func (vec *Vector3) RoundToZero() {
	if vec.X < 0 {
		vec.X = Ceil(vec.X)
	} else {
		vec.X = Floor(vec.X)
	}

	if vec.Y < 0 {
		vec.Y = Ceil(vec.Y)
	} else {
		vec.Y = Floor(vec.Y)
	}

	if vec.Z < 0 {
		vec.Z = Ceil(vec.Z)
	} else {
		vec.Z = Floor(vec.Z)
	}
}

func (vec *Vector3) Negate() {
	vec.X = -vec.X
	vec.Y = -vec.Y
	vec.Z = -vec.Z
}

func (vec *Vector3) Dot(v *Vector3) float32 {
	return vec.X*v.X + vec.Y*v.Y + vec.Z*v.Z
}

func (vec *Vector3) Cross(v *Vector3) {
	ax := vec.X
	ay := vec.Y
	az := vec.Z

	vec.X = ay*v.Z - az*v.Y
	vec.Y = az*v.X - ax*v.Z
	vec.Z = ax*v.Y - ay*v.X
}

func (vec *Vector3) CrossVectors(a *Vector3, b *Vector3) {
	vec.X = a.Y*b.Z - a.Z*b.Y
	vec.Y = a.Z*b.X - a.X*b.Z
	vec.Z = a.X*b.Y - a.Y*b.X
}

func (vec *Vector3) GetLengthSq() float32 {
	return vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z
}

func (vec *Vector3) GetLength() float32 {
	return Sqrt(vec.GetLengthSq())
}

func (vec *Vector3) SetLength(length float32) {
	vec.Normalize()
	vec.MultiplyScalar(length)
}

func (vec *Vector3) GetManhattanLength() float32 {
	return Abs(vec.X) + Abs(vec.Y) + Abs(vec.Z)
}

func (vec *Vector3) Normalize() {
	div := vec.GetLength()
	if div == 0 {
		div = 1
	}
	vec.DivideScalar(div)
}

func (vec *Vector3) ProjectOnVector(v *Vector3) {
	scalar := v.Dot(vec) * v.GetLengthSq()
	vec.Copy(v)
	vec.MultiplyScalar(scalar)
}

func (vec *Vector3) ProjectOnPlane(p *Plane) {
	v1 := vec.Clone()
	v1.ProjectOnVector(p.GetNormal())
	vec.Sub(v1)
}

func (vec *Vector3) Reflect(normal *Vector3) {
	v1 := normal.Clone()
	v1.MultiplyScalar(2 * vec.Dot(normal))
	vec.Sub(v1)
}

func (vec *Vector3) AngleTo(v *Vector3) float32 {
	theta := vec.Dot(v) / (Sqrt(vec.GetLengthSq() * v.GetLengthSq()))
	return Acos(Clamp(theta, -1, 1))
}

func (vec *Vector3) GetDistanceTo(v *Vector3) float32 {
	return Sqrt(vec.GetDistanceToSquared(v))
}

func (vec *Vector3) GetDistanceToSquared(v *Vector3) float32 {
	dx := vec.X - v.X
	dy := vec.Y - v.Y
	dz := vec.Z - v.Z
	return dx*dx + dy*dy + dz*dz
}

func (vec *Vector3) GetManhattanDistanceTo(v *Vector3) float32 {
	return Abs(vec.X-v.X) + Abs(vec.Y-v.Y) + Abs(vec.Z-v.Z)
}

func (vec *Vector3) Lerp(v *Vector3, alpha float32) {
	vec.X += (v.X - vec.X) * alpha
	vec.Y += (v.Y - vec.Y) * alpha
	vec.Z += (v.Z - vec.Z) * alpha
}

func (vec *Vector3) LerpVectors(v1 *Vector3, v2 *Vector3, alpha float32) {
	vec.SetSubVectors(v2, v1)
	vec.MultiplyScalar(alpha)
	vec.Add(v1)
}

func (vec *Vector3) Equals(v *Vector3) bool {
	return vec.X == v.X && vec.Y == v.Y && vec.Z == v.Z
}

func (vec *Vector3) ToArray() [3]float32 {
	return [3]float32{vec.X, vec.Y, vec.Z}
}

func (vec *Vector3) CopyToArray(array []float32, offset int) {
	va := vec.ToArray()
	copy(array[offset:], va[0:])
}
