package math

type Plane struct {
	normal *Vector3
	constant float64
}

func NewPlane(normal *Vector3, constant float64) *Plane {
	np := &Plane{}
	np.Set(normal, constant)
	return np
}

func (plane *Plane) GetNormal() *Vector3 {
	return plane.normal
}

func (plane *Plane) GetConstant() float64 {
	return plane.constant
}

func (plane *Plane) Set(normal *Vector3, constant float64) {
	plane.normal = normal.Clone()
	plane.constant = constant
}

func (plane *Plane) SetComponents(x, y, z float64, w float64) {
	plane.normal.Set(x, y, z)
	plane.constant = w
}

func (plane *Plane) SetFromNormalAndCoplanarPoint(normal *Vector3, point *Vector3) {
	plane.normal = normal.Clone()
	plane.constant = -point.Dot(point)
}

func (plane *Plane) SetFromCoplanarPoints(a *Vector3, b *Vector3, c *Vector3) {
	v1 := NewVector3(0, 0, 0)
	v2 := NewVector3(0, 0, 0)

	v1.SetSubVectors(c, b)
	v2.SetSubVectors(a, b)
	v1.Cross(v2)
	v1.Normalize()

	plane.SetFromNormalAndCoplanarPoint(v1, a)
}