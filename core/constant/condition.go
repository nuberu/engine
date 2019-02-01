package constant

type Condition uint32

const (
	Never        Condition = 0
	Less         Condition = 1
	Equal        Condition = 2
	LessEqual    Condition = 3
	Greater      Condition = 4
	NotEqual     Condition = 5
	GreaterEqual Condition = 6
	Always       Condition = 7
)
