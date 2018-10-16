package math

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(x float64, y float64) *Vector2 {
	return &Vector2{
		X: x,
		Y: y,
	}
}

func NewVector2FromArray(arr []float64, offset int) *Vector2 {
	return &Vector2{
		X: arr[offset],
		Y: arr[offset+1],
	}
}

func (vec *Vector2) GetWidth() float64 {
	return vec.X
}

func (vec *Vector2) SetWidth(width float64) {
	vec.X = width
}

func (vec *Vector2) SetX(x float64) {
	vec.X = x
}

func (vec *Vector2) GetHeight() float64 {
	return vec.Y
}

func (vec *Vector2) SetHeight(height float64) {
	vec.Y = height
}

func (vec *Vector2) SetY(y float64) {
	vec.Y = y
}

func (vec *Vector2) Set(x float64, y float64) {
	vec.X = x
	vec.Y = y
}

func (vec *Vector2) SetScalar(num float64) {
	vec.X = num
	vec.Y = num
}

func (vec *Vector2) Clone() *Vector2 {
	return &Vector2{
		X: vec.X,
		Y: vec.Y,
	}
}

func (vec *Vector2) Add(v *Vector2) {
	vec.X += v.X
	vec.Y += v.Y
}

func (vec *Vector2) AddScalar(num float64) {
	vec.X += num
	vec.Y += num
}

func (vec *Vector2) SetAddVectors(v1 *Vector2, v2 *Vector2) {
	vec.X = v1.X + v2.X
	vec.Y = v1.Y + v2.Y
}

func (vec *Vector2) AddScaledVector(v1 *Vector2, scale float64) {
	vec.X += v1.X * scale
	vec.Y += v1.Y * scale
}

func (vec *Vector2) Sub(v *Vector2) {
	vec.X -= v.X
	vec.Y -= v.Y
}

func (vec *Vector2) SubScalar(num float64) {
	vec.X -= num
	vec.Y -= num
}

func (vec *Vector2) SetSubVectors(v1 *Vector2, v2 *Vector2) {
	vec.X = v1.X - v2.X
	vec.Y = v1.Y - v2.Y
}

func (vec *Vector2) Multiply(v *Vector2) {
	vec.X *= v.X
	vec.Y *= v.Y
}

func (vec *Vector2) MultiplyScalar(num float64) {
	vec.X *= num
	vec.Y *= num
}

func (vec *Vector2) Divide(v *Vector2) {
	vec.X /= v.X
	vec.Y /= v.Y
}

func (vec *Vector2) DivideScalar(num float64) {
	vec.X /= num
	vec.Y /= num
}

func (vec *Vector2) ApplyMatrix3(m *Matrix3) {
	x := vec.X
	y := vec.Y
	e := m.GetElements()

	vec.X = e[0]*x + e[3]*y + e[6]
	vec.Y = e[1]*x + e[4]*y + e[7]
}

func (vec *Vector2) Min(v *Vector2) {
	vec.X = math.Min(vec.X, v.X)
	vec.Y = math.Min(vec.Y, v.Y)
}

func (vec *Vector2) Max(v *Vector2) {
	vec.X = math.Max(vec.X, v.X)
	vec.Y = math.Max(vec.Y, v.Y)
}

/*
 Clamps the value to be between min and max.
 */
func (vec *Vector2) Clamp(min *Vector2, max *Vector2) {
	vec.X = math.Max(min.X, math.Min(max.X, vec.X))
	vec.Y = math.Max(min.Y, math.Min(max.Y, vec.Y))
}

func (vec *Vector2) ClampScalar(min float64, max float64) {
	minVec := NewVector2(min, min)
	maxVec := NewVector2(max, max)
	vec.Clamp(minVec, maxVec)
}

func (vec *Vector2) ClampLength(min float64, max float64) {
	length := vec.GetLength()

	div := length
	if length == 0 {
		div = 1
	}

	vec.DivideScalar(div)
	vec.MultiplyScalar(math.Max(min, math.Min(max, length)))
}

func (vec *Vector2) Floor() {
	vec.X = math.Floor(vec.X)
	vec.Y = math.Floor(vec.Y)
}

func (vec *Vector2) Ceil() {
	vec.X = math.Ceil(vec.X)
	vec.Y = math.Ceil(vec.Y)
}

func (vec *Vector2) Round() {
	vec.X = math.Round(vec.X)
	vec.Y = math.Round(vec.Y)
}

func (vec *Vector2) RoundToZero() {
	if vec.X < 0 {
		vec.X = math.Ceil(vec.X)
	} else {
		vec.X = math.Floor(vec.X)
	}

	if vec.Y < 0 {
		vec.Y = math.Ceil(vec.Y)
	} else {
		vec.Y = math.Floor(vec.Y)
	}
}

func (vec *Vector2) Negate() {
	vec.X = -vec.X
	vec.Y = -vec.Y
}

func (vec *Vector2) Dot(v *Vector2) float64 {
	return vec.X*v.X + vec.Y + v.Y
}

func (vec *Vector2) Cross(v *Vector2) float64 {
	return vec.X*v.X - vec.Y + v.Y
}

func (vec *Vector2) GetLengthSq() float64 {
	return vec.X*vec.X + vec.Y*vec.Y
}

func (vec *Vector2) GetLength() float64 {
	return math.Sqrt(vec.X*vec.X + vec.Y*vec.Y)
}

func (vec *Vector2) SetLength(length float64) {
	vec.Normalize()
	vec.MultiplyScalar(length)
}

func (vec *Vector2) GetManhattanLength() float64 {
	return math.Abs(vec.X) + math.Abs(vec.Y)
}

func (vec *Vector2) Normalize() {
	div := vec.GetLength()
	if div == 0 {
		div = 1
	}
	vec.DivideScalar(div)
}

func (vec *Vector2) GetAngle() float64 {
	angle := math.Atan2(vec.Y, vec.X)

	if angle < 0 {
		angle += 2 * math.Pi
	}

	return angle
}

func (vec *Vector2) GetDistanceTo(v *Vector2) float64 {
	return math.Sqrt(vec.GetDistanceToSquared(v))
}

func (vec *Vector2) GetDistanceToSquared(v *Vector2) float64 {
	dx := vec.X - v.X
	dy := vec.Y - v.Y
	return dx*dx + dy*dy
}

func (vec *Vector2) GetManhattanDistanceTo(v *Vector2) float64 {
	return math.Abs(vec.X-v.X) + math.Abs(vec.Y-v.Y)
}

func (vec *Vector2) Lerp(v *Vector2, alpha float64) {
	vec.X += (v.X - vec.X) * alpha
	vec.Y += (v.Y - vec.Y) * alpha
}

func (vec *Vector2) LerpVectors(v1 *Vector2, v2 *Vector2, alpha float64) {
	vec.SetSubVectors(v2, v1)
	vec.MultiplyScalar(alpha)
	vec.Add(v1)
}

func (vec *Vector2) Equals(v *Vector2) bool {
	return vec.X == v.X && vec.Y == v.Y
}

func (vec *Vector2) RotateAround(center *Vector2, angle float64) {
	c := math.Cos(angle)
	s := math.Sin(angle)

	x := vec.X - center.X
	y := vec.Y - center.Y

	vec.X = x*c - y*s + center.X
	vec.Y = x*s + y*c + center.Y
}

func (vec *Vector2) ToArray() []float64 {
	return []float64{vec.X, vec.Y}
}

func (vec *Vector2) CopyToArray(array []float64, offset int) {
	copy(array[offset:], vec.ToArray())
}