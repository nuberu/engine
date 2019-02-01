package buffer

import (
	"errors"
	"github.com/nuberu/engine/math"
)

type Float32 struct {
	buffer
	data []float32
}

func NewFloat32(size uint, normalized bool) *Float32 {
	return &Float32{
		buffer: buffer{
			size:       size,
			normalized: normalized,
			dynamic:    false,
		},
		data: make([]float32, size),
	}
}

func NewFloat32Dynamic() *Float32 {
	return &Float32{
		buffer: buffer{
			size:       0,
			normalized: false,
			dynamic:    true,
		},
		data: make([]float32, 0),
	}
}

func (buffer *Float32) Set(position uint, value float32) error {
	if position < buffer.size {
		buffer.data[position] = value
	} else if buffer.dynamic {
		buffer.data = append(buffer.data, make([]float32, position-uint(len(buffer.data)))...)
	} else {
		return errors.New("out of bounds")
	}
	return nil
}

func (buffer *Float32) SetVector3(position uint, value *math.Vector3) error {
	rPos := position * 3
	err := buffer.Set(rPos, value.X)
	if err != nil {
		return err
	}
	buffer.Set(rPos+1, value.Y)
	if err != nil {
		return err
	}
	buffer.Set(rPos+2, value.Z)
	if err != nil {
		return err
	}
	return nil
}
