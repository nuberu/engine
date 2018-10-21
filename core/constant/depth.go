package constant

type Depth uint8

const (
	NeverDepth        Depth = 1
	AlwaysDepth       Depth = 2
	LessDepth         Depth = 3
	LessEqualDepth    Depth = 4
	EqualDepth        Depth = 5
	GreaterEqualDepth Depth = 6
	GreaterDepth      Depth = 7
	NotEqualDepth     Depth = 8
)
