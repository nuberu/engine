package types

import (
	"testing"
)

func TestAngle_Equivalent(t *testing.T) {
	a := Angle(0)
	a += AngleFromSexagesimal(360)

	if !a.Equivalent(Angle(0)) {
		t.Errorf("The angles must be equal, but a = %s and b = %s are different", a.Simplify().ToString(), Angle(0).ToString())
	}
}
