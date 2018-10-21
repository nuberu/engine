package constant

type Factor uint8

const (
	NilFactor              Factor = 0
	ZeroFactor             Factor = 1
	OneFactor              Factor = 2
	SrcColorFactor         Factor = 3
	OneMinusSrcColorFactor Factor = 4
	SrcAlphaFactor         Factor = 5
	OneMinusSrcAlphaFactor Factor = 6
	DstAlphaFactor         Factor = 7
	OneMinusDstAlphaFactor Factor = 8
	DstColorFactor         Factor = 9
	OneMinusDstColorFactor Factor = 10
	SrcAlphaSaturateFactor Factor = 11
)
