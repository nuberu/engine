package math

type Box2 struct {
	min Vector2
	max Vector2
}

// Equivalent to makeEmpty
func NewDefaultBox2() *Box2 {
	return NewBox2(
		NewVector2Inf(1),
		NewVector2Inf(-1),
	)
}

func NewBox2(min *Vector2, max *Vector2) *Box2 {
	return &Box2{
		min: Vector2{X: min.X, Y: min.Y},
		max: Vector2{X: max.X, Y: max.Y},
	}
}

func NewBox2FromComponents(minX, minY, maxX, maxY float32) *Box2 {
	return &Box2{
		min: Vector2{X: minX, Y: minY},
		max: Vector2{X: maxX, Y: maxY},
	}
}

func NewBox2FromCenterAndSize(center *Vector2, size *Vector2) *Box2 {
	halfSize := size.Clone()
	halfSize.MultiplyScalar(0.5)

	nb := &Box2{
		min: *center.Clone(),
		max: *center.Clone(),
	}

	nb.min.Sub(halfSize)
	nb.max.Add(halfSize)

	return nb
}

func (box *Box2) Clone() *Box2 {
	return &Box2{
		min: *box.min.Clone(),
		max: *box.max.Clone(),
	}
}

func (box *Box2) Copy(source *Box2) {
	box.min.Copy(&source.min)
	box.max.Copy(&source.max)
}

func (box *Box2) IsEmpty() bool {
	return box.max.X < box.min.X || box.max.Y < box.min.Y
}

func (box *Box2) GetCenter() *Vector2 {
	if box.IsEmpty() {
		return NewVector2(0, 0)
	} else {
		c := NewVector2(0, 0)
		c.SetAddVectors(&box.min, &box.max)
		c.MultiplyScalar(0.5)
		return c
	}
}

func (box *Box2) GetSize() *Vector2 {
	if box.IsEmpty() {
		return NewVector2(0, 0)
	} else {
		c := NewVector2(0, 0)
		c.SetSubVectors(&box.min, &box.max)
		return c
	}
}

func (box *Box2) ExpandByPoint(point *Vector2) {
	box.min.Min(point)
	box.max.Max(point)
}

func (box *Box2) ExpandByVector(vector *Vector2) {
	box.min.Sub(vector)
	box.max.Add(vector)
}

func (box *Box2) ExpandByScalar(scalar float32) {
	box.min.SubScalar(scalar)
	box.max.AddScalar(scalar)
}

func (box *Box2) ContainsPoint(point *Vector2) bool {
	return !(point.X < box.min.X || point.X > box.max.X ||
		point.Y < box.min.Y || point.Y > box.max.Y)
}

func (box *Box2) ContainsBox(b *Box2) bool {
	return box.min.X <= b.min.X && b.max.X <= box.max.X &&
		box.min.Y <= b.min.Y && b.max.Y <= box.max.Y
}

func (box *Box2) IntersectsBox(b *Box2) bool {
	return !(b.max.X < box.min.X || b.min.X > box.max.X ||
		b.max.Y < box.min.Y || b.min.Y > box.max.Y)
}

func (box *Box2) ClampPoint(point *Vector2) *Vector2 {
	target := NewDefaultVector2()
	target.Copy(point)
	target.Clamp(&box.min, &box.max)
	return target
}

func (box *Box2) DistanceToPoint(point *Vector2) float32 {
	v1 := point.Clone()
	v1.Clamp(&box.min, &box.max)
	v1.Sub(point)
	return v1.GetLength()
}

func (box *Box2) Intersect(b *Box2) {
	box.min.Max(&b.min)
	box.max.Min(&b.max)
}

func (box *Box2) Union(b *Box2) {
	box.min.Min(&b.min)
	box.max.Max(&b.max)
}

func (box *Box2) Translate(offset *Vector2) {
	box.min.Add(offset)
	box.max.Add(offset)
}

func (box *Box2) Equals(b *Box2) bool {
	return box.min.Equals(&b.min) && box.max.Equals(&b.max)
}
