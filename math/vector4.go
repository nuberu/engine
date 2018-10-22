package math

type Vector4 struct {
	Vector3
	W float32
}

func NewDefaultVector4() *Vector4 {
	return NewVector4(0, 0, 0, 1)
}

func NewVector4(x float32, y float32, z float32, w float32) *Vector4 {
	return &Vector4{
		Vector3: *NewVector3(x, y, z),
		W:       w,
	}
}

func NewVector4FromArray(arr []float32, offset int) *Vector4 {
	return NewVector4(
		arr[offset],
		arr[offset+1],
		arr[offset+2],
		arr[offset+3],
	)
}

func (vec *Vector4) Set(x float32, y float32, z float32, w float32) {
	vec.X = x
	vec.Y = y
	vec.Z = z
	vec.W = w
}

func (vec *Vector4) SetW(w float32) {
	vec.W = w
}

func (vec *Vector4) SetScalar(scalar float32) {
	vec.X = scalar
	vec.Y = scalar
	vec.Z = scalar
	vec.W = scalar
}

func (vec *Vector4) Clone() *Vector4 {
	return &Vector4{
		Vector3: *vec.Vector3.Clone(),
		W:       vec.W,
	}
}

func (vec *Vector4) Copy(vector *Vector4) {
	vec.X = vector.X
	vec.Y = vector.Y
	vec.Z = vector.Z
	vec.W = vector.W
}

func (vec *Vector4) Add(vector *Vector4) {
	vec.X += vector.X
	vec.Y += vector.Y
	vec.Z += vector.Z
	vec.W += vector.W
}

func (vec *Vector4) AddComponents(x float32, y float32, z float32, w float32) {
	vec.X += x
	vec.Y += y
	vec.Z += z
	vec.W += w
}

func (vec *Vector4) AddScalar(num float32) {
	vec.X += num
	vec.Y += num
	vec.Z += num
	vec.W += num
}

func (vec *Vector4) SetAddVectors(v1 *Vector4, v2 *Vector4) {
	vec.X = v1.X + v2.X
	vec.Y = v1.Y + v2.Y
	vec.Z = v1.Z + v2.Z
	vec.W = v1.W + v2.W
}

func (vec *Vector4) AddScaledVector(v1 *Vector4, scale float32) {
	vec.X += v1.X * scale
	vec.Y += v1.Y * scale
	vec.Z += v1.Z * scale
	vec.W += v1.W * scale
}

func (vec *Vector4) Sub(v *Vector4) {
	vec.X -= v.X
	vec.Y -= v.Y
	vec.Z -= v.Z
	vec.W -= v.W
}

func (vec *Vector4) SubScalar(num float32) {
	vec.X -= num
	vec.Y -= num
	vec.Z -= num
	vec.W -= num
}

func (vec *Vector4) SetSubVectors(v1 *Vector4, v2 *Vector4) {
	vec.X = v1.X - v2.X
	vec.Y = v1.Y - v2.Y
	vec.Z = v1.Z - v2.Z
	vec.W = v1.W - v2.W
}

func (vec *Vector4) Multiply(v *Vector4) {
	vec.X *= v.X
	vec.Y *= v.Y
	vec.Z *= v.Z
	vec.W *= v.W
}

func (vec *Vector4) MultiplyScalar(num float32) {
	vec.X *= num
	vec.Y *= num
	vec.Z *= num
	vec.W *= num
}

func (vec *Vector4) Divide(v *Vector4) {
	vec.X /= v.X
	vec.Y /= v.Y
	vec.Z /= v.Z
	vec.W /= v.W
}

func (vec *Vector4) DivideScalar(num float32) {
	vec.MultiplyScalar(1 / num)
}

func (vec *Vector4) ApplyMatrix4(m *Matrix4) {
	x := vec.X
	y := vec.Y
	z := vec.Z
	e := m.GetElements()
	w := 1 / (e[3]*x + e[7]*y + e[11]*z + e[15])

	vec.X = (e[0]*x + e[4]*y + e[8]*z + e[12]) * w
	vec.Y = (e[1]*x + e[5]*y + e[9]*z + e[13]) * w
	vec.Z = (e[2]*x + e[6]*y + e[10]*z + e[14]) * w
	vec.Z = (e[3]*x + e[7]*y + e[11]*z + e[15]) * w
}

func (vec *Vector4) SetAxisAngleFromQuaternion(q *Quaternion) {
	vec.W = 2 * Acos(q.w)
	s := Sqrt(1 - q.w*q.w)

	if s < 0.0001 {
		vec.X = 1
		vec.Y = 0
		vec.Z = 0
	} else {
		vec.X = q.x / s
		vec.Y = q.y / s
		vec.Z = q.z / s
	}
}

// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToAngle/index.htm
// assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled)
func (vec *Vector4) SetAxisAngleFromRotationMatrix(m *Matrix4) {
	angle := float32(0)

	// variables for result
	x := float32(0)
	y := float32(0)
	z := float32(0)

	// margin to allow for rounding errors
	epsilon := float32(0.01)

	// margin to distinguish between 0 and 180 degrees
	epsilon2 := float32(0.1)

	m11 := m.elements[0]
	m12 := m.elements[4]
	m13 := m.elements[8]
	m21 := m.elements[1]
	m22 := m.elements[5]
	m23 := m.elements[9]
	m31 := m.elements[2]
	m32 := m.elements[6]
	m33 := m.elements[10]

	if (Abs(m12-m21) < epsilon) &&
		(Abs(m13-m31) < epsilon) &&
		(Abs(m23-m32) < epsilon) {

		// singularity found
		// first check for identity matrix which must have +1 for all terms
		// in leading diagonal and zero in other terms

		if (Abs(m12+m21) < epsilon2) &&
			(Abs(m13+m31) < epsilon2) &&
			(Abs(m23+m32) < epsilon2) &&
			(Abs(m11+m22+m33-3) < epsilon2) {

			// this singularity is identity matrix so angle = 0

			vec.Set(1, 0, 0, 0)

			// zero angle, arbitrary axis
			return
		}

		// otherwise this singularity is angle = 180

		angle = Pi

		xx := (m11 + 1) / 2
		yy := (m22 + 1) / 2
		zz := (m33 + 1) / 2
		xy := (m12 + m21) / 4
		xz := (m13 + m31) / 4
		yz := (m23 + m32) / 4

		if (xx > yy) && (xx > zz) {
			// m11 is the largest diagonal term
			if xx < epsilon {
				x = 0
				y = 0.707106781
				z = 0.707106781
			} else {
				x = Sqrt(xx)
				y = xy / x
				z = xz / x
			}
		} else if yy > zz {
			// m22 is the largest diagonal term
			if yy < epsilon {
				x = 0.707106781
				y = 0
				z = 0.707106781
			} else {
				y = Sqrt(yy)
				x = xy / y
				z = yz / y
			}
		} else {
			// m33 is the largest diagonal term so base result on this
			if zz < epsilon {
				x = 0.707106781
				y = 0.707106781
				z = 0
			} else {
				z = Sqrt(zz)
				x = xz / z
				y = yz / z
			}
		}

		vec.Set(x, y, z, angle)
		return
	}

	// as we have reached here there are no singularities so we can handle normally
	var s = Sqrt((m32-m23)*(m32-m23) + (m13-m31)*(m13-m31) + (m21-m12)*(m21-m12)) // used to normalize

	if Abs(s) < 0.001 {
		s = 1
	}

	// prevent divide by zero, should not happen if matrix is orthogonal and should be
	// caught by singularity test above, but I've left it in just in case
	vec.X = (m32 - m23) / s
	vec.Y = (m13 - m31) / s
	vec.Z = (m21 - m12) / s
	vec.W = Acos((m11 + m22 + m33 - 1) / 2)
}

func (vec *Vector4) Min(v *Vector4) {
	vec.X = Min(vec.X, v.X)
	vec.Y = Min(vec.Y, v.Y)
	vec.Z = Min(vec.Z, v.Z)
	vec.W = Min(vec.W, v.W)
}

func (vec *Vector4) Max(v *Vector4) {
	vec.X = Max(vec.X, v.X)
	vec.Y = Max(vec.Y, v.Y)
	vec.Z = Max(vec.Z, v.Z)
	vec.W = Max(vec.W, v.W)
}

// Clamps the value to be between min and max.
func (vec *Vector4) Clamp(min *Vector4, max *Vector4) {
	vec.X = Max(min.X, Min(max.X, vec.X))
	vec.Y = Max(min.Y, Min(max.Y, vec.Y))
	vec.Z = Max(min.Z, Min(max.Z, vec.Z))
	vec.W = Max(min.W, Min(max.W, vec.W))
}

func (vec *Vector4) ClampScalar(min float32, max float32) {
	minVec := NewVector4(min, min, min, min)
	maxVec := NewVector4(max, max, max, max)
	vec.Clamp(minVec, maxVec)
}

func (vec *Vector4) ClampLength(min float32, max float32) {
	length := vec.GetLength()

	div := length
	if length == 0 {
		div = 1
	}

	vec.DivideScalar(div)
	vec.MultiplyScalar(Max(min, Min(max, length)))
}

func (vec *Vector4) Floor() {
	vec.X = Floor(vec.X)
	vec.Y = Floor(vec.Y)
	vec.Z = Floor(vec.Z)
	vec.W = Floor(vec.W)
}

func (vec *Vector4) Ceil() {
	vec.X = Ceil(vec.X)
	vec.Y = Ceil(vec.Y)
	vec.Z = Ceil(vec.Z)
	vec.W = Ceil(vec.W)
}

func (vec *Vector4) Round() {
	vec.X = Round(vec.X)
	vec.Y = Round(vec.Y)
	vec.Z = Round(vec.Z)
	vec.W = Round(vec.W)
}

func (vec *Vector4) RoundToZero() {
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

	if vec.W < 0 {
		vec.W = Ceil(vec.W)
	} else {
		vec.W = Floor(vec.W)
	}
}

func (vec *Vector4) Negate() {
	vec.X = -vec.X
	vec.Y = -vec.Y
	vec.Z = -vec.Z
	vec.W = -vec.W
}

func (vec *Vector4) Dot(v *Vector4) float32 {
	return vec.X*v.X + vec.Y*v.Y + vec.Z*v.Z + vec.W*v.W
}

func (vec *Vector4) GetLengthSq() float32 {
	return vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z + vec.W*vec.W
}

func (vec *Vector4) GetLength() float32 {
	return Sqrt(vec.GetLengthSq())
}

func (vec *Vector4) SetLength(length float32) {
	vec.Normalize()
	vec.MultiplyScalar(length)
}

func (vec *Vector4) GetManhattanLength() float32 {
	return Abs(vec.X) + Abs(vec.Y) + Abs(vec.Z) + Abs(vec.W)
}

func (vec *Vector4) Normalize() {
	div := vec.GetLength()
	if div == 0 {
		div = 1
	}
	vec.DivideScalar(div)
}

func (vec *Vector4) Lerp(v *Vector4, alpha float32) {
	vec.X += (v.X - vec.X) * alpha
	vec.Y += (v.Y - vec.Y) * alpha
	vec.Z += (v.Z - vec.Z) * alpha
	vec.W += (v.W - vec.W) * alpha
}

func (vec *Vector4) LerpVectors(v1 *Vector4, v2 *Vector4, alpha float32) {
	vec.SetSubVectors(v2, v1)
	vec.MultiplyScalar(alpha)
	vec.Add(v1)
}

func (vec *Vector4) Equals(v *Vector4) bool {
	return vec.X == v.X && vec.Y == v.Y && vec.Z == v.Z && vec.W == v.W
}

func (vec *Vector4) ToArray() [4]float32 {
	return [4]float32{vec.X, vec.Y, vec.Z, vec.W}
}

func (vec *Vector4) CopyToArray(array []float32, offset int) {
	va := vec.ToArray()
	copy(array[offset:], va[0:])
}