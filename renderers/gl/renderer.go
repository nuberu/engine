package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/nuberu/engine/core/render"
)

type Renderer struct {
}

func NewRenderer(target render.Target) (*Renderer, error) {
	renderer := &Renderer{}

	err := renderer.Init()

	return renderer, err
}

func (renderer *Renderer) Init() error {
	if err := gl.Init(); err != nil {
		return err
	}

	return nil
}