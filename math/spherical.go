package math

import "math"

type Spherical struct {
	radius float64
	phi    float64 // polar angle
	theta  float64 // azimuthal angle
}

func NewDefaultSpherical() *Spherical {
	return NewSpherical(1, 0, 0)
}

func NewSpherical(radius float64, phi float64, theta float64) *Spherical {
	return &Spherical{
		radius: radius,
		phi:    phi,
		theta:  theta,
	}
}

func (sphere *Spherical) Set(radius float64, phi float64, theta float64) {
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

func (sphere *Spherical) Copy(other *Spherical) {
	sphere.radius = other.radius
	sphere.phi = other.phi
	sphere.theta = other.theta
}

func (sphere *Spherical) MakeSafe() {
	const EPS = float64(0.000001)
	sphere.phi = math.Max(EPS, math.Min(math.Pi-EPS, sphere.phi))
}

func (sphere *Spherical) SetFromVector3(vec *Vector3) {
	sphere.SetFromCartesianCoordinates(vec.X, vec.Y, vec.Z)
}

func (sphere *Spherical) SetFromCartesianCoordinates(x float64, y float64, z float64) {
	sphere.radius = math.Sqrt(x*x + y*y + z*z)

	if sphere.radius == 0 {
		sphere.theta = 0
		sphere.phi = 0
	} else {
		sphere.theta = math.Atan2(x, z)
		sphere.phi = math.Acos(Clamp(y/sphere.radius, - 1, 1))
	}
}
