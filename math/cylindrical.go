package math

type Cylindrical struct {
	radius float32 // distance from the origin to a point in the x-z plane
	theta  float32 // counterclockwise angle in the x-z plane measured in radians from the positive z-axis
	y      float32 // height above the x-z plane
}

func NewDefaultCylindrical() *Cylindrical {
	return NewCylindrical(1, 0, 0)
}

func NewCylindrical(radius float32, theta float32, y float32) *Cylindrical {
	return &Cylindrical{
		radius: radius,
		theta:  theta,
		y:      y,
	}
}

func (cyl *Cylindrical) Set(radius float32, theta float32, y float32) {
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

func (cyl *Cylindrical) Copy(source *Cylindrical) {
	cyl.radius = source.radius
	cyl.theta = source.theta
	cyl.y = source.y
}

func (cyl *Cylindrical) SetFromVector3(vec *Vector3) {
	cyl.SetFromCartesianCoordinates(vec.X, vec.Y, vec.Z)
}

func (cyl *Cylindrical) SetFromCartesianCoordinates(x float32, y float32, z float32) {
	cyl.radius = Sqrt(x*x + z*z)
	cyl.theta = Atan2(x, z)
	cyl.y = y
}
