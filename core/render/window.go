package render

import "github.com/tokkenno/seed/math"

type Window interface {
	GetSize() *math.Vector2
	SetSize(size *math.Vector2)
	SetTitle(title string)
}
