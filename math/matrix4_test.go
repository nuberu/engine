package math

import (
	"testing"
)

var (
	namedExample    *Matrix4
	mixedExample    *Matrix4
	mixedExampleInv *Matrix4
	mixedExampleDet float64
)

func init() {
	namedExample = NewMatrix4(11, 12, 13, 14, 21, 22, 23, 24, 31, 32, 33, 34, 41, 42, 43, 44)
	mixedExample = NewMatrix4(1, 2, 3, 4, 5, -6, -7, 8, -9, 10, 11, -12, -13, -14, -15, -16)
	mixedExampleInv = NewMatrix4(float64(-2)/float64(9), float64(-1)/float64(2), float64(-1)/float64(3), float64(-1)/float64(18), float64(-5)/float64(8), float64(11)/float64(8), float64(7)/float64(8), float64(-1)/float64(8), float64(7)/float64(12), float64(-5)/float64(4), float64(-3)/float64(4), float64(1)/float64(12), float64(13)/float64(72), float64(3)/float64(8), float64(5)/float64(24), float64(1)/float64(72))
	mixedExampleDet = float64(576)
}

func TestMatrix4_CopyPosition(t *testing.T) {
	m1 := namedExample.Clone()
	m2 := new(Matrix4)

	m2.CopyPosition(m1)
	if m2.elements[11] != 0 || m2.elements[12] != m1.elements[12] || m2.elements[13] != m1.elements[13] || m2.elements[14] != m1.elements[14] || m2.elements[15] != 0 {
		t.Errorf("The method don't copy the position correctly from %v to %v", m1.elements, m2.elements)
	}
}

func TestMatrix4_GetDeterminant(t *testing.T) {
	m1 := mixedExample.Clone()
	det := m1.GetDeterminant()

	if det != mixedExampleDet {
		t.Errorf("The determinant of the matrix %v must be %v but %v calculated", m1.elements, mixedExampleDet, det)
	}
}

func TestMatrix4_GetInverse(t *testing.T) {
	m1 := mixedExample.Clone()
	err := m1.Inverse()

	if err != nil {
		t.Errorf(err.Error())
	} else if !m1.EqualsRound(mixedExampleInv.Clone(), 6) {
		t.Errorf("The inverse of the matrix\n%s\nbut is\n%s", mixedExampleInv.ToString(), m1.ToString())
	}
}
