package types

type Subset struct {
	val float32
	def float32
	min float32
	max float32
}

func NewSubset(min, max, def float32) Subset {
	return Subset{
		val: def,
		def: def,
		min: min,
		max: max,
	}
}

func (ss *Subset) Clone() Subset {
	return Subset{
		ss.val,
		ss.def,
		ss.min,
		ss.max,
	}
}
