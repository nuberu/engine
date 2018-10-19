package core

type Layers struct {
	mask uint64
}

func (layers *Layers) Set(channel int) {
	layers.mask = 1 << uint(channel) | 0
}

func (layers *Layers) Enable(channel int) {
	layers.mask |= 1 << uint(channel) | 0
}

func (layers *Layers) Toggle(channel int) {
	layers.mask ^= 1 << uint(channel) | 0
}

func (layers *Layers) Disable(channel int) {
	layers.mask ^= 1 << uint(channel) | 0
}

func (layers *Layers) Test(channel int) bool {
	return (layers.mask & layers.mask) != 0
}