package constant

type StencilFunc string

const (
	Keep         StencilFunc = "KEEP"
	Replace      StencilFunc = "REPLACE"
	Increment    StencilFunc = "INCR"
	Decrement    StencilFunc = "DECR"
	Inverr       StencilFunc = "INVERT"
	IncreaseWrap StencilFunc = "INCR_WRAP"
	DecreaseWrap StencilFunc = "DECR_WRAP"
)
