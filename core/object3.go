package core

import (
	"github.com/tokkenno/seed/event"
	"github.com/tokkenno/seed/math"
)

const (
	defaultMatrixAutoUpdate bool = true
)

var (
	object3IdGenerator = new(IdGenerator)
)

type Object3 struct {
	id                     Id
	Name                   string
	Parent                 *Object3
	children               []*Object3
	up                     math.Vector3
	Position               math.Vector3
	Rotation               math.Euler
	quaternion             math.Quaternion
	Scale                  math.Vector3
	modelViewMatrix        math.Matrix4
	normalMatrix           math.Matrix4
	matrix                 math.Matrix4
	matrixWorld            math.Matrix4
	matrixAutoUpdate       bool
	matrixWorldNeedsUpdate bool
	layers                 Layers
	Visible                bool
	castShadow             bool
	receiveShadow          bool
	frustumCulled          bool
	RenderOrder            uint
}

func NewObject() *Object3 {
	obj := &Object3{
		id:                     object3IdGenerator.Next(),
		Name:                   "",
		Parent:                 nil,
		children:               []*Object3{},
		up:                     *math.NewVector3(0, 1, 0),
		Position:               math.Vector3{Vector2: math.Vector2{X: 0, Y: 0}, Z: 0},
		Rotation:               *math.NewDefaultEuler(),
		quaternion:             *math.NewDefaultQuaternion(),
		Scale:                  math.Vector3{Vector2: math.Vector2{X: 0, Y: 0}, Z: 0},
		modelViewMatrix:        *math.NewDefaultMatrix4(),
		normalMatrix:           *math.NewDefaultMatrix4(),
		matrix:                 *math.NewDefaultMatrix4(),
		matrixWorld:            *math.NewDefaultMatrix4(),
		matrixAutoUpdate:       defaultMatrixAutoUpdate,
		matrixWorldNeedsUpdate: false,
		layers:                 Layers{mask: 0},
		Visible:                true,
		castShadow:             false,
		receiveShadow:          false,
		frustumCulled:          true,
		RenderOrder:            0,
	}

	obj.Rotation.OnChange().Always(obj.onEulerChange)
	obj.quaternion.OnChange().Always(obj.onQuaternionChange)

	return obj
}

func (obj *Object3) onQuaternionChange(sender interface{}, args *event.Args) {

}

func (obj *Object3) onEulerChange(sender interface{}, args *event.Args) {

}

func (obj *Object3) GetId() Id {
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
