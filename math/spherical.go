package math

type Spherical struct {
	radius float32
	phi    float32 // polar angle
	theta  float32 // azimuthal angle
}

func NewDefaultSpherical() *Spherical {
	return NewSpherical(1, 0, 0)
}

func NewSpherical(radius float32, phi float32, theta float32) *Spherical {
	return &Spherical{
		radius: radius,
		phi:    phi,
		theta:  theta,
	}
}

func (sphere *Spherical) Set(radius float32, phi float32, theta float32) {
	sphere.radius = radius
	sphere.phi = phi
	sphere.theta = theta
}

func (sphere *Spherical) Clone() *Spherical {
	return &Spherical{
		radius: sphere.radius,
		phi:    sphere.phi,
		theta:  sphere.theta,
	}
}

func (sphere *Spherical) Copy(source *Spherical) {
	sphere.radius = source.radius
	sphere.phi = source.phi
	sphere.theta = source.theta
}

func (sphere *Spherical) MakeSafe() {
	const EPS = float32(0.000001)
	sphere.phi = Max(EPS, Min(Pi-EPS, sphere.phi))
}

func (sphere *Spherical) SetFromVector3(vec *Vector3) {
	sphere.SetFromCartesianCoordinates(vec.X, vec.Y, vec.Z)
}

func (sphere *Spherical) SetFromCartesianCoordinates(x float32, y float32, z float32) {
	sphere.radius = Sqrt(x*x + y*y + z*z)

	if sphere.radius == 0 {
		sphere.theta = 0
		sphere.phi = 0
	} else {
		sphere.theta = Atan2(x, z)
		sphere.phi = Acos(Clamp(y/sphere.radius, - 1, 1))
	}
}
