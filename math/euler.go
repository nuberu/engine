package math

import "github.com/tokkenno/seed/event"

type EulerOrder int

const (
	EulerOrderXYZ     EulerOrder = 0
	EulerOrderYZX     EulerOrder = 1
	EulerOrderZXY     EulerOrder = 2
	EulerOrderXZY     EulerOrder = 3
	EulerOrderYXZ     EulerOrder = 4
	EulerOrderZYX     EulerOrder = 5
	EulerOrderDefault            = EulerOrderXYZ
)

type Euler struct {
	x     float64
	y     float64
	z     float64
	order EulerOrder

	changeEvent *event.Emitter
}

func (eu *Euler) OnChange() *event.Handler {
	return eu.changeEvent.GetHandler()
}

func NewDefaultEuler() *Euler {
	e := &Euler{
		x:     0,
		y:     0,
		z:     0,
		order: EulerOrderDefault,
		changeEvent: event.NewEvent(),
	}
	return e
}

func NewEuler(x float64, y float64, z float64, order EulerOrder) *Euler {
	return &Euler{
		x:     x,
		y:     y,
		z:     z,
		order: order,
	}
}
