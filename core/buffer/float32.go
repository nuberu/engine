package buffer

import "errors"

type Float32 struct {
	buffer
	data []float32
}

func NewFloat32(size int, normalized bool) *Float32 {
	return &Float32 {
		buffer: buffer{
			size: size,
			normalized: normalized,
			dynamic:    false,
		},
		data: make([]float32, size),
	}
}

func NewFloat32Dynamic() *Float32 {
	return &Float32 {
		buffer: buffer{
			size: 0,
			normalized: false,
			dynamic:    true,
		},
		data: make([]float32, 0),
	}
}

func (buffer *Float32) Set(position int, value float32) error {
	if position < buffer.size {
		buffer.data[position] = value
	} else if buffer.dynamic {
		buffer.data = append(buffer.data, make([]float32, position-len(buffer.data))...)
	} else {
		return errors.New("out of bounds")
	}
	return nil
}