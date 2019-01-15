package constant

type Condition uint8

const (
	Never        Condition = 1
	Always       Condition = 2
	Less         Condition = 3
	LessEqual    Condition = 4
	Equal        Condition = 5
	GreaterEqual Condition = 6
	Greater      Condition = 7
	NotEqual     Condition = 8
)
