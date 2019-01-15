package math

type Box3 struct {
	min Vector3
	max Vector3
}

// Equivalent to makeEmpty
func NewDefaultBox3() *Box3 {
	return NewBox3(
		NewVector3Inf(1),
		NewVector3Inf(-1),
	)
}

func NewBox3(min *Vector3, max *Vector3) *Box3 {
	return &Box3{
		min: Vector3{Vector2: Vector2{X: min.X, Y: min.Y}, Z: min.Z},
		max: Vector3{Vector2: Vector2{X: max.X, Y: max.Y}, Z: max.Z},
	}
}

func NewBox3FromComponents(minX, minY, minZ, maxX, maxY, maxZ float32) *Box3 {
	return &Box3{
		min: Vector3{Vector2: Vector2{X: minX, Y: minY}, Z: minZ},
		max: Vector3{Vector2: Vector2{X: maxX, Y: maxY}, Z: maxZ},
	}
}

func NewBox3FromPoints(points []*Vector3) *Box3 {
	box := NewDefaultBox3()
	box.SetFromPoints(points)
	return box
}

func NewBox3FromCenterAndSize(center *Vector3, size *Vector3) *Box3 {
	halfSize := size.Clone()
	halfSize.MultiplyScalar(0.5)

	nb := &Box3{
		min: *center.Clone(),
		max: *center.Clone(),
	}

	nb.min.Sub(halfSize)
	nb.max.Add(halfSize)

	return nb
}

func (box *Box3) Clone() *Box3 {
	return &Box3{
		min: *box.min.Clone(),
		max: *box.max.Clone(),
	}
}

func (box *Box3) Copy(source *Box3) {
	box.min.Copy(&source.min)
	box.max.Copy(&source.max)
}

func (box *Box3) SetFromPoints(points []*Vector3) {
	for i := 0; i < len(points); i++ {
		box.ExpandByPoint(points[i])
	}
}

func (box *Box3) IsEmpty() bool {
	return box.max.X < box.min.X || box.max.Y < box.min.Y || box.max.Z < box.min.Z
}

func (box *Box3) GetCenter() *Vector3 {
	if box.IsEmpty() {
		return NewVector3(0, 0, 0)
	} else {
		c := NewDefaultVector3()
		c.SetAddVectors(&box.min, &box.max)
		c.MultiplyScalar(0.5)
		return c
	}
}

func (box *Box3) GetSize() *Vector3 {
	if box.IsEmpty() {
		return NewVector3(0, 0, 0)
	} else {
		c := NewDefaultVector3()
		c.SetSubVectors(&box.min, &box.max)
		return c
	}
}

func (box *Box3) ExpandByPoint(point *Vector3) {
	box.min.Min(point)
	box.max.Max(point)
}

func (box *Box3) ExpandByVector(vector *Vector3) {
	box.min.Sub(vector)
	box.max.Add(vector)
}

func (box *Box3) ExpandByScalar(scalar float32) {
	box.min.SubScalar(scalar)
	box.max.AddScalar(scalar)
}

func (box *Box3) ContainsPoint(point *Vector3) bool {
	return !(point.X < box.min.X || point.X > box.max.X ||
		point.Y < box.min.Y || point.Y > box.max.Y ||
		point.Z < box.min.Z || point.Z > box.max.Z)
}

func (box *Box3) ContainsBox(b *Box3) bool {
	return box.min.X <= b.min.X && b.max.X <= box.max.X &&
		box.min.Y <= b.min.Y && b.max.Y <= box.max.Y &&
		box.min.Z <= b.min.Z && b.max.Z <= box.max.Z
}

func (box *Box3) IntersectsBox(b *Box3) bool {
	return !(b.max.X < box.min.X || b.min.X > box.max.X ||
		b.max.Y < box.min.Y || b.min.Y > box.max.Y ||
		b.max.Z < box.min.Z || b.min.Z > box.max.Z)
}

func (box *Box3) ClampPoint(point *Vector3) *Vector3 {
	target := NewDefaultVector3()
	target.Copy(point)
	target.Clamp(&box.min, &box.max)
	return target
}

func (box *Box3) DistanceToPoint(point *Vector3) float32 {
	v1 := point.Clone()
	v1.Clamp(&box.min, &box.max)
	v1.Sub(point)
	return v1.GetLength()
}

func (box *Box3) Intersect(b *Box3) {
	box.min.Max(&b.min)
	box.max.Min(&b.max)
}

func (box *Box3) Union(b *Box3) {
	box.min.Min(&b.min)
	box.max.Max(&b.max)
}

func (box *Box3) Translate(offset *Vector3) {
	box.min.Add(offset)
	box.max.Add(offset)
}

func (box *Box3) Equals(b *Box3) bool {
	return box.min.Equals(&b.min) && box.max.Equals(&b.max)
}