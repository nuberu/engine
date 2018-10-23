package gl

import (
	"github.com/tokkenno/seed/cameras"
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/math"
)

type Nosequemierdaesesto struct {
	Value       []float32
	NeedsUpdate bool
}

type Clipping struct {
	globalState          []float32
	numGlobalPlanes      int
	localClippingEnabled bool
	renderingShadows     bool
	plane                *math.Plane
	viewNormalMatrix     *math.Matrix3
	uniform              *Nosequemierdaesesto

	numPlanes       int
	numIntersection uint64
}

func (scope *Clipping) Init(planes []*math.Plane, enableLocalClipping bool, camera *cameras.Camera) {

}

func (scope *Clipping) BeginShadows() {
	scope.renderingShadows = true
	scope.projectPlanes(nil, nil, 0, false)
}

func (scope *Clipping) EndShadows() {
	scope.renderingShadows = false
	scope.resetGlobalState()
}

func (scope *Clipping) resetGlobalState() {
	if scope.uniform.Value != scope.globalState {
		scope.uniform.Value = scope.globalState
		scope.uniform.NeedsUpdate = scope.numGlobalPlanes > 0
	}

	scope.numPlanes = scope.numGlobalPlanes
	scope.numIntersection = 0
}

func (scope *Clipping) projectPlanes(planes []*math.Plane, camera *core.Camera, dstOffset int, skipTransform bool) []float32 {
	nPlanes := 0
	if planes != nil {
		nPlanes = len(planes)
	}

	var dstArray []float32 = nil

	if nPlanes != 0 {

		dstArray = scope.uniform.Value;
		if skipTransform != true || dstArray == nil {

			flatSize := dstOffset + nPlanes*4
			viewMatrix := camera.GetMatrixWorldInverse()

			scope.viewNormalMatrix.SetNormalMatrix(viewMatrix)

			if dstArray == nil || len(dstArray) < flatSize {
				dstArray = make([]float32, flatSize)
			}

			for i := 0; i < nPlanes; i++ {
				scope.plane = planes[i].Clone()
				scope.plane.ApplyMatrix4(viewMatrix, scope.viewNormalMatrix)

				dstArray = append(dstArray, scope.plane.GetNormal().ToArray()[:]...)
				dstArray[i+4+3] = scope.plane.GetConstant()
			}
		}

		scope.uniform.Value = dstArray
		scope.uniform.NeedsUpdate = true
	}

	scope.numPlanes = nPlanes

	return dstArray
}
