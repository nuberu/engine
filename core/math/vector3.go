package math

type Vector3 struct {
	Vector2
	Z float64
}

func NewVector3(x float64, y float64, z float64) *Vector3 {
	return &Vector3{
		Vector2: *NewVector2(x, y),
		Z:       z,
	}
}

func (v *Vector3) Add(vec *Vector3) *Vector3 {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
	return v
}

func (v *Vector3) AddComponents(x float64, y float64, z float64) *Vector3 {
	v.X += x
	v.Y += y
	v.Z += z
	return v
}

func (v *Vector3) Set(x float64, y float64, z float64) *Vector3 {
	v.X = x
	v.Y = y
	v.Z = z
	return v
}



func (vec *Vector3) Clone() *Vector3 {
	return &Vector3{
		Vector2: *vec.Vector2.Clone(),
		Z: vec.Z,
	}
}