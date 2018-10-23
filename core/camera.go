package core

import (
	"github.com/tokkenno/seed/math"
)

type Camera struct {
	Object3

	matrixWorldInverse      math.Matrix4
	projectionMatrix        math.Matrix4
	projectionMatrixInverse math.Matrix4
}

func NewCamera() *Camera {
	cam := Camera{
		Object3: *NewObject(),
		matrixWorldInverse: *math.NewDefaultMatrix4(),
		projectionMatrix: *math.NewDefaultMatrix4(),
		projectionMatrixInverse: *math.NewDefaultMatrix4(),
	}

	return &cam
}

func (camera *Camera) IsCamera() bool {
	return true
}

func (camera *Camera) GetMatrixWorldInverse() *math.Matrix4 {
	return &camera.matrixWorldInverse
}

func (camera *Camera) GetProjectionMatrix() *math.Matrix4 {
	return &camera.projectionMatrix
}

func (camera *Camera) GetProjectionMatrixInverse() *math.Matrix4 {
	return &camera.projectionMatrixInverse
}

func (camera *Camera) Copy(source *Camera, recursive bool) {
	camera.Object3.Copy(&source.Object3, recursive)

	camera.matrixWorldInverse.Copy(&source.matrixWorldInverse)
	camera.projectionMatrix.Copy(&source.projectionMatrix)
	camera.projectionMatrixInverse.Copy(&source.projectionMatrixInverse)
}

func (camera *Camera) Clone() *Camera {
	newCamera := new(Camera)
	newCamera.Copy(camera, true)
	return newCamera
}

func (camera *Camera) UpdateMatrixWorld(force bool) {
	camera.Object3.UpdateMatrixWorld(force)
	camera.matrixWorldInverse.SetInverseOf(camera.GetMatrixWorld(), false)
}

func (camera *Camera) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	camera.UpdateMatrixWorld(true)
	e := camera.GetMatrixWorld().GetElements()

	target.Set(-e[8], -e[9], -e[10])
	target.Normalize()
	return target
}
