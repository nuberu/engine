package core

type Id uint

type IdGenerator struct {
	v uint64
}

func (id *IdGenerator) Next() Id {
	id.v++
	return Id(id.v)
}