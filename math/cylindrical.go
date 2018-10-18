package math

import "math"

type Cylindrical struct {
	radius float64 // distance from the origin to a point in the x-z plane
	theta  float64 // counterclockwise angle in the x-z plane measured in radians from the positive z-axis
	y      float64 // height above the x-z plane
}

func NewDefaultCylindrical() *Cylindrical {
	return NewCylindrical(1, 0, 0)
}

func NewCylindrical(radius float64, theta float64, y float64) *Cylindrical {
	return &Cylindrical{
		radius: radius,
		theta:  theta,
		y:      y,
	}
}

func (cyl *Cylindrical) Set(radius float64, theta float64, y float64) {
	cyl.radius = radius
	cyl.theta = theta
	cyl.y = y
}

func (cyl *Cylindrical) Clone() *Cylindrical {
	return &Cylindrical{
		radius: cyl.radius,
		theta:  cyl.theta,
		y:      cyl.y,
	}
}

func (cyl *Cylindrical) Copy(other *Cylindrical) {
	cyl.radius = other.radius
	cyl.theta = other.theta
	cyl.y = other.y
}

func (cyl *Cylindrical) SetFromVector3(vec *Vector3) {
	cyl.SetFromCartesianCoordinates(vec.X, vec.Y, vec.Z)
}

func (cyl *Cylindrical) SetFromCartesianCoordinates(x float64, y float64, z float64) {
	cyl.radius = math.Sqrt(x*x + z*z)
	cyl.theta = math.Atan2(x, z)
	cyl.y = y
}
