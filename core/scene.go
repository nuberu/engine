package core

type Scene struct {
	Object
	AutoUpdateRender bool
}

func NewScene() *Scene {
	scene := new(Scene)
	scene.AutoUpdateRender = true
	return scene;
}

func (scene *Scene) Copy(source *Scene, recursive bool) {
	scene.Object.Copy(&source.Object, recursive)
}
