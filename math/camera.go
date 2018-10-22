package math

import "github.com/tokkenno/seed/core"

func Project(vec *Vector3, camera *core.Camera) {
	vec.ApplyMatrix4(camera.GetMatrixWorldInverse())
	vec.ApplyMatrix4(camera.GetProjectionMatrix())
}

func UnProject(vec *Vector3, camera *core.Camera) {
	matrix := NewDefaultMatrix4()
	matrix.SetIdentity()
	matrix.SetInverseOf(camera.GetProjectionMatrix(), false)
	vec.ApplyMatrix4(matrix)
	vec.ApplyMatrix4(camera.GetMatrixWorld())
}

