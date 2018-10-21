package core

type Scene struct {
	Object3
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
