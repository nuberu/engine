package core

import (
	"github.com/tokkenno/seed/core/math"
)

type Object struct {
	id       uint64
	Name     string
	Parent   *Object
	children []*Object

	position *math.Vector3
	rotation *math.Vector3
	scale    *math.Vector3

	matrix      *math.Matrix4
	matrixWorld *math.Matrix4

	matrixAutoUpdate       bool
	matrixWorldNeedsUpdate bool
}

func NewObject() *Object {
	obj := new(Object)

	obj.matrixAutoUpdate = true

	return obj
}

func (obj *Object) GetId() uint64 {
	return obj.id
}

func (obj *Object) Copy(source *Object, recursive bool) {

}

func (obj *Object) UpdateMatrix() {

}

func (obj *Object) UpdateMatrixWorld(force bool) {
	if obj.matrixAutoUpdate {
		obj.UpdateMatrix()
	}

	if obj.matrixWorldNeedsUpdate || force {

		if obj.Parent == nil {
			obj.matrixWorld.Copy(obj.matrix)
		} else {
			obj.matrixWorld.multiplyMatrices(obj.Parent.matrixWorld, obj.matrix)
		}

		obj.matrixWorldNeedsUpdate = false
		force = true
	}

	// update children
	for _, child := range obj.children {
		child.UpdateMatrixWorld(force)
	}
}
