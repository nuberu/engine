package math

type Matrix4 struct {
	elements [16]float64
}

func Matrix4Identity() *Matrix4 {
	return &Matrix4{
		elements: [16]float64{
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		},
	}
}

func (matrix *Matrix4) GetElements() [16]float64 {
	return matrix.elements
}

func (matrix *Matrix4) Copy(m *Matrix4) {
}

func (matrix *Matrix4) Inverse(m *Matrix4, errorOnDegenerate bool) error {
	return nil
}
