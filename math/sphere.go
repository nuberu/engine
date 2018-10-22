package math

type Sphere struct {
	Center Vector3
	Radius float32
}

func NewDefaultSphere() *Sphere {
	return NewSphere(NewDefaultVector3(), 1)
}

func NewSphere(vector *Vector3, radius float32) *Sphere {
	return &Sphere{
		Center: Vector3{Vector2: Vector2{X: vector.X, Y: vector.Y}, Z: vector.Z},
		Radius: radius,
	}
}

func (sphere *Sphere) SetFromPoints(points []*Vector3) {
	box := NewBox3FromPoints(points)
	sphere.SetFromPointsAndCenter(points, box.GetCenter())
}

func (sphere *Sphere) SetFromPointsAndCenter(points []*Vector3, center *Vector3) {
	sphere.Center.Copy(center)

	var maxRadiusSq = float32(0)

	for i := 0; i < len(points); i++ {
		maxRadiusSq = Max(maxRadiusSq, center.GetDistanceToSquared(points[i]))
	}

	sphere.Radius = Sqrt(maxRadiusSq)
}
