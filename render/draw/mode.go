package draw

type Mode uint8

const (
	TrianglesDrawMode     Mode = 0
	TriangleStripDrawMode Mode = 1
	TriangleFanDrawMode   Mode = 2
)
