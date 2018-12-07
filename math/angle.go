package math

import (
	"fmt"
	"math"
)

const radiansCircle = float32(2) * float32(math.Pi)
const sexagesimal = float32(360)

// An angle expressed in radians
type Angle float32

// If the angle is bigger than a complete rotation, simplify it removing the extra rotations
func (angle Angle) Simplify() Angle {
	if angle < 0 {
		return Angle(float32(angle) + radiansCircle).Simplify()
	} else if float32(angle) < radiansCircle {
		return angle
	} else {
		return Angle(float32(angle) - radiansCircle).Simplify()
	}
}

func (angle Angle) Equivalent(a Angle) bool {
	return angle.Simplify() == a.Simplify()
}

func (angle Angle) ToString() string {
	if float32(angle) >= 0 || float32(angle) < radiansCircle {
		return fmt.Sprintf("%.2fº (%.5f rad)", angle.ToSexagesimal(), angle)
	} else {
		return fmt.Sprintf("%f rad", angle)
	}
}

func AngleFromSexagesimal(angle float32) Angle {
	return Angle(angle / sexagesimal * radiansCircle)
}

func (angle Angle) ToSexagesimal() float32 {
	return float32(angle) * sexagesimal / radiansCircle
}
