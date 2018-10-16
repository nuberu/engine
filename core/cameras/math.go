package cameras

import "github.com/tokkenno/seed/core/math"

func Project(vec *math.Vector3, camera *Camera) {
	vec.ApplyMatrix4(camera.GetMatrixWorldInverse())
	vec.ApplyMatrix4(camera.GetProjectionMatrix())
}

func UnProject(vec *math.Vector3, camera *Camera) {
	matrix := math.Matrix4Identity()
	matrix.Inverse(camera.GetProjectionMatrix(), false)
	vec.ApplyMatrix4(matrix)
	vec.ApplyMatrix4(camera.GetMatrixWorld())
}
