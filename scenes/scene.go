package scenes

import (
	"github.com/tokkenno/seed/core"
)

type Scene struct {
	core.Object3
	AutoUpdateRender bool
}

func NewScene() *Scene {
	scene := new(Scene)
	scene.AutoUpdateRender = true
	return scene
}

func (scene *Scene) Copy(source *Scene, recursive bool) {
	scene.Object3.Copy(&source.Object3, recursive)
}
