package cameras

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/core/math"
)

type Camera struct {
	core.Object

	matrixWorld             *math.Matrix4
	matrixWorldInverse      *math.Matrix4
	projectionMatrix        *math.Matrix4
	projectionMatrixInverse *math.Matrix4
}

func (camera *Camera) Copy(source *Camera, recursive bool) {
	camera.Object.Copy(&source.Object, recursive)
	camera.matrixWorldInverse.Copy(source.matrixWorldInverse)
	camera.projectionMatrix.Copy(source.projectionMatrix)
	camera.projectionMatrixInverse.Copy(source.projectionMatrixInverse)
}

func (camera *Camera) Clone() *Camera {
	newCamera := new(Camera)
	newCamera.Copy(camera, true)
	return newCamera
}

func (camera *Camera) UpdateMatrixWorld(force bool) {
	camera.Object.UpdateMatrixWorld(force)
	camera.matrixWorldInverse.Inverse(camera.matrixWorld, false)
}

func (camera *Camera) GetMatrixWorld() *math.Matrix4 {
	return camera.matrixWorld
}

func (camera *Camera) GetMatrixWorldInverse() *math.Matrix4 {
	return camera.matrixWorldInverse
}

func (camera *Camera) GetProjectionMatrix() *math.Matrix4 {
	return camera.projectionMatrix
}
