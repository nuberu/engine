package math

import (
	"github.com/tokkenno/seed/core/event"
	"math"
)

type Quaternion struct {
	x float64
	y float64
	z float64
	w float64

	changeEvent *event.Emitter
}

func NewEmptyQuaternion() *Quaternion {
	return NewQuaternion(0, 0, 0, 0)
}

func NewQuaternion(x float64, y float64, z float64, w float64) *Quaternion {
	return &Quaternion{
		x:           x,
		y:           y,
		z:           z,
		w:           w,
		changeEvent: event.NewEvent(),
	}
}

func (qua *Quaternion) GetX() float64 {
	return qua.x
}

func (qua *Quaternion) SetX(x float64) {
	qua.x = x
	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) GetY() float64 {
	return qua.y
}

func (qua *Quaternion) SetY(y float64) {
	qua.y = y
	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) GetZ() float64 {
	return qua.z
}

func (qua *Quaternion) SetZ(z float64) {
	qua.z = z
	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) GetW() float64 {
	return qua.w
}

func (qua *Quaternion) SetW(w float64) {
	qua.w = w
	qua.changeEvent.Emit(qua, nil)
}

func (qua *Quaternion) SetFromEuler(euler *Euler, update bool) {

}

func (qua *Quaternion) SetFromAxisAngle(axis *Vector3, angle float64) {
	halfAngle := angle / 2
	s := math.Sin(halfAngle)

	qua.x = axis.X * s
	qua.y = axis.Y * s
	qua.z = axis.Z * s
	qua.w = math.Cos(halfAngle)

	qua.changeEvent.Emit(qua, nil)
}
