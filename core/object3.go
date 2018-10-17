package core

import (
	"github.com/tokkenno/seed/core/math"
)

type Object3 struct {
	id       uint64
	Name     string
	Parent   *Object3
	children []*Object3

	position *math.Vector3
	rotation *math.Vector3
	scale    *math.Vector3

	matrix      *math.Matrix4
	matrixWorld *math.Matrix4

	matrixAutoUpdate       bool
	matrixWorldNeedsUpdate bool
}

func NewObject() *Object3 {
	obj := new(Object3)

	obj.matrixAutoUpdate = true

	return obj
}

func (obj *Object3) GetId() uint64 {
	return obj.id
}

/**
 Recursive = TRUE by default
 */
func (obj *Object3) Copy(source *Object3, recursive bool) {

}

func (obj *Object3) UpdateMatrix() {

}

func (obj *Object3) UpdateMatrixWorld(force bool) {
	if obj.matrixAutoUpdate {
		obj.UpdateMatrix()
	}

	if obj.matrixWorldNeedsUpdate || force {

		if obj.Parent == nil {
			obj.matrixWorld.Copy(obj.matrix)
		} else {
			obj.matrixWorld.MultiplyMatrices(obj.Parent.matrixWorld, obj.matrix)
		}

		obj.matrixWorldNeedsUpdate = false
		force = true
	}

	// update children
	for _, child := range obj.children {
		child.UpdateMatrixWorld(force)
	}
}

func (obj *Object3) GetMatrixWorld() *math.Matrix4 {
	return obj.matrixWorld
}