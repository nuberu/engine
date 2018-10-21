package buffer

// Memory optimized buffer
type buffer struct {
	typeSize   uint8
	size       uint
	normalized bool
	dynamic    bool
}

func (buf *buffer) SetDynamic(dynamic bool) {
	buf.dynamic = dynamic
}
