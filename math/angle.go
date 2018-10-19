package math

import (
	"fmt"
	"math"
)

const radiansCircle = float64(2) * float64(math.Pi)
const sexagesimal = float64(360)

// An angle expressed in radians
type Angle float64

// If the angle is bigger than a complete rotation, simplify it removing the extra rotations
func (angle Angle) Simplify() Angle {
	if angle < 0 {
		return Angle(float64(angle) + radiansCircle).Simplify()
	} else if float64(angle) < radiansCircle {
		return angle
	} else {
		return Angle(float64(angle) - radiansCircle).Simplify()
	}
}

func (angle Angle) Equivalent(a Angle) bool {
	return angle.Simplify() == a.Simplify()
}

func (angle Angle) ToString() string {
	if float64(angle) >= 0 || float64(angle) < radiansCircle {
		return fmt.Sprintf("%.2fº (%.5f rad)", angle.ToSexagesimal(), angle)
	} else {
		return fmt.Sprintf("%f rad", angle)
	}
}

func AngleFromSexagesimal(angle float64) Angle {
	return Angle(angle / sexagesimal * radiansCircle)
}

func (angle Angle) ToSexagesimal() float64 {
	return float64(angle) * sexagesimal / radiansCircle
}
