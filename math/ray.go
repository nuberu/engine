package math

type Ray struct {
	origin    Vector3
	direction Vector3
}

func NewDefaultRay() *Ray {
	return &Ray {
		origin: *NewDefaultVector3(),
		direction: *NewDefaultVector3(),
	}
}

func NewRay(origin *Vector3, direction *Vector3) *Ray {
	return &Ray {
		origin: *origin.Clone(),
		direction: *direction.Clone(),
	}
}

func (ray *Ray) Set(origin *Vector3, direction *Vector3) {
	ray.origin.Copy(origin)
	ray.direction.Copy(direction)
}

func (ray *Ray) Copy(other *Ray) {
	ray.origin.Copy(&other.origin)
	ray.direction.Copy(&other.direction)
}

func (ray *Ray) Clone() *Ray {
	nr := NewDefaultRay()
	nr.Copy(ray)
	return nr
}

func (ray *Ray) GetAt(t float32) *Vector3 {
	target := NewDefaultVector3()
	target.Copy(&ray.direction)
	target.MultiplyScalar(t)
	target.Add(&ray.origin)
	return target
}

func (ray *Ray) SetLookAt(v *Vector3) {
	ray.direction.Copy(v)
	ray.direction.Sub(&ray.origin)
	ray.direction.Normalize()
}

func (ray *Ray) SetRecast(t float32) {
	target := ray.GetAt(t)
	ray.origin.Copy(target)
}

func (ray *Ray) GetClosestPointToPoint(point *Vector3) *Vector3 {
	target := NewDefaultVector3()
	target.SetSubVectors(point, &ray.origin)

	directionDistance := target.Dot(&ray.direction)

	if directionDistance < 0 {
		target.Copy(&ray.origin)
	} else {
		target.Copy(&ray.direction)
		target.MultiplyScalar(directionDistance)
		target.Add(&ray.origin)
	}

	return target
}

func (ray *Ray) GetDistanceToPoint(point *Vector3) float32 {
	return Sqrt(ray.GetDistanceSqToPoint(point))
}

func (ray *Ray) GetDistanceSqToPoint(point *Vector3) float32 {
	v1 := NewDefaultVector3()
	v1.SetSubVectors(point, &ray.origin)
	directionDistance := v1.Dot(&ray.direction)

	if directionDistance < 0 {
		return ray.origin.GetDistanceToSquared(point)
	} else {
		v1.Copy(&ray.direction)
		v1.MultiplyScalar(directionDistance)
		v1.Add(&ray.origin)
		return v1.GetDistanceToSquared(point)
	}
}

func (ray *Ray) GetDistanceSqToSegment(v0 *Vector3, v1 *Vector3) float32 {
	return 0 // TODO
}