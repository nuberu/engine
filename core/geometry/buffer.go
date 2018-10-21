package geometry

type bufferGroup struct {
	start uint
	count uint
	materialIndex uint
}

type Buffer struct {
	groups []bufferGroup
	index uint
}

func (buf *Buffer) SetIndex(index uint) {
	// TODO
}

func (buf *Buffer) AddGroup(start, count uint) {
	buf.AddGroupIndex(start, count, 0)
}

func (buf *Buffer) AddGroupIndex(start, count, index uint) {
	buf.groups = append(buf.groups, bufferGroup{
		start: start,
		count: count,
		materialIndex: index,
	})
}