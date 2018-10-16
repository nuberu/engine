package math

type Vector4 struct {
	Vector3
	W float64
}

func NewVector4(x float64, y float64, z float64, w float64) *Vector4 {
	return &Vector4{
		Vector3: *NewVector3(x, y, z),
		W:       w,
	}
}

func (v *Vector4) Add(vec *Vector4) *Vector4 {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
	v.W += vec.W
	return v
}

func (v *Vector4) AddComponents(x float64, y float64, z float64, w float64) *Vector4 {
	v.X += x
	v.Y += y
	v.Z += z
	v.W += w
	return v
}

func (v *Vector4) Set(x float64, y float64, z float64, w float64) *Vector4 {
	v.X = x
	v.Y = y
	v.Z = z
	v.W = w
	return v
}



func (vec *Vector4) Clone() *Vector4 {
	return &Vector4{
		Vector3: *vec.Vector3.Clone(),
		W: vec.W,
	}
}