package core

import (
	"github.com/tokkenno/seed/core/types"
	"github.com/tokkenno/seed/core/event"
	"github.com/tokkenno/seed/math"
)

const (
	defaultMatrixAutoUpdate bool = true
)

var (
	object3IdGenerator = new(IdGenerator)
)

type CameraObject interface {
	IsCamera() bool
}

type Object3 struct {
	id                     Id
	Name                   string
	parent                 *Object3
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

	beforeRenderEvent *event.Emitter
	afterRenderEvent  *event.Emitter
	addEvent          *event.Emitter // Event fired when the object is added to other
	removeEvent       *event.Emitter // Event fired when the object is removed from other
}

func NewObject() *Object3 {
	return newObjectWithId(object3IdGenerator.Next())
}

func newObjectWithId(id Id) *Object3 {
	obj := &Object3{
		id:                     id,
		Name:                   "",
		parent:                 nil,
		children:               []*Object3{},
		up:                     *math.NewVector3(0, 1, 0),
		Position:               math.Vector3{Vector2: math.Vector2{X: 0, Y: 0}, Z: 0},
		Rotation:               *math.NewDefaultEuler(),
		quaternion:             *math.NewDefaultQuaternion(),
		Scale:                  math.Vector3{Vector2: math.Vector2{X: 1.0, Y: 1.0}, Z: 1.0},
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

		beforeRenderEvent: event.NewEvent(),
		afterRenderEvent:  event.NewEvent(),
		addEvent:          event.NewEvent(),
		removeEvent:       event.NewEvent(),
	}

	obj.Rotation.OnChange().Always(obj.onEulerChange)
	obj.quaternion.OnChange().Always(obj.onQuaternionChange)

	return obj
}

func (obj *Object3) GetId() Id {
	return obj.id
}

func (obj *Object3) IsCamera() bool {
	return false
}

func (obj *Object3) IsVisible() bool {
	return obj.Visible
}

func (obj *Object3) OnBeforeRender() *event.Handler {
	return obj.beforeRenderEvent.GetHandler()
}

func (obj *Object3) OnAfterRender() *event.Handler {
	return obj.afterRenderEvent.GetHandler()
}

func (obj *Object3) onQuaternionChange(sender interface{}, args *event.Args) {

}

func (obj *Object3) onEulerChange(sender interface{}, args *event.Args) {

}

func (obj *Object3) ApplyMatrix(matrix *math.Matrix4) {
	obj.matrix.MultiplyMatrices(matrix, &obj.matrix)
	obj.matrix.DeCompose(&obj.Position, &obj.quaternion, &obj.Scale)
}

func (obj *Object3) ApplyQuaternion(q *math.Quaternion) {
	obj.quaternion.PreMultiply(q)
}

func (obj *Object3) SetRotationFromAxisAngle(axis *math.Vector3, angle types.Angle) {
	obj.quaternion.SetFromAxisAngle(axis, angle)
}

func (obj *Object3) SetRotationFromEuler(euler *math.Euler) {
	obj.quaternion.SetFromEuler(euler, true)
}

func (obj *Object3) SetRotationFromMatrix(m *math.Matrix4) {
	obj.quaternion.SetFromRotationMatrix(m)
}

func (obj *Object3) SetRotationFromQuaternion(q *math.Quaternion) {
	obj.quaternion.Copy(q)
}

func (obj *Object3) RotateOnAxis(axis *math.Vector3, angle types.Angle) {
	q1 := math.NewDefaultQuaternion()
	q1.SetFromAxisAngle(axis, angle)
	obj.quaternion.Multiply(q1)
}

func (obj *Object3) RotateOnWorldAxis(axis *math.Vector3, angle types.Angle) {
	q1 := math.NewDefaultQuaternion()
	q1.SetFromAxisAngle(axis, angle)
	obj.quaternion.PreMultiply(q1)
}

func (obj *Object3) RotateX(angle types.Angle) {
	obj.RotateOnAxis(math.NewVector3(1, 0, 0), angle)
}

func (obj *Object3) RotateY(angle types.Angle) {
	obj.RotateOnAxis(math.NewVector3(0, 1, 0), angle)
}

func (obj *Object3) RotateZ(angle types.Angle) {
	obj.RotateOnAxis(math.NewVector3(0, 0, 1), angle)
}

func (obj *Object3) TranslateOnAxis(axis *math.Vector3, distance float32) {
	tmp := axis.Clone()
	tmp.ApplyQuaternion(&obj.quaternion)
	tmp.MultiplyScalar(distance)
	obj.Position.Add(tmp)
}

func (obj *Object3) TranslateX(distance float32) {
	obj.TranslateOnAxis(math.NewVector3(1, 0, 0), distance)
}

func (obj *Object3) TranslateY(distance float32) {
	obj.TranslateOnAxis(math.NewVector3(0, 1, 0), distance)
}

func (obj *Object3) TranslateZ(distance float32) {
	obj.TranslateOnAxis(math.NewVector3(0, 0, 1), distance)
}

func (obj *Object3) LocalToWorld(vector *math.Vector3) {
	vector.ApplyMatrix4(&obj.matrixWorld)
}

func (obj *Object3) WorldToLocal(vector *math.Vector3) {
	m1 := math.NewDefaultMatrix4()
	m1.SetInverseOf(&obj.matrixWorld, false)
	vector.ApplyMatrix4(m1)
}

func (obj *Object3) LookAtComponents(x, y, z float32) {
	obj.LookAt(math.NewVector3(x, y, z))
}

func (obj *Object3) LookAt(x *math.Vector3) {
	q1 := math.NewDefaultQuaternion()
	m1 := math.NewDefaultMatrix4()
	target := math.NewDefaultVector3()
	position := math.NewDefaultVector3()

	obj.updateWorldMatrix(true, false)

	position.SetFromMatrixPosition(&obj.matrixWorld)

	if obj.IsCamera() {
		m1.LookAt(position, target, &obj.up)
	} else {
		m1.LookAt(target, position, &obj.up)
	}

	obj.quaternion.SetFromRotationMatrix(m1)

	if obj.parent != nil {
		m1.ExtractRotation(&obj.parent.matrixWorld)
		q1.SetFromRotationMatrix(m1)
		q1.Inverse()
		obj.quaternion.PreMultiply(q1)
	}
}

func (obj *Object3) Add(object *Object3) {
	if object.parent != nil {
		object.parent.Remove(object)
	}

	object.parent = obj
	object.addEvent.Emit(obj, nil)

	obj.children = append(obj.children, object)
}

func (obj *Object3) AddAll(objects []*Object3) {
	for _, object := range objects {
		obj.Add(object)
	}
}

func (obj *Object3) Remove(object *Object3) {
	for i, child := range obj.children {
		if child == object {
			object.parent = nil
			object.removeEvent.Emit(obj, nil)
			obj.children = append(obj.children[:i], obj.children[i+1:]...)
		}
	}
}

func (obj *Object3) RemoveAll(objects []*Object3) {
	for _, object := range objects {
		obj.Remove(object)
	}
}

func (obj *Object3) GetParent() *Object3 {
	return obj.parent
}

func (obj *Object3) GetChildren() []*Object3 {
	castChildren := make([]*Object3, len(obj.children))
	copy(castChildren, obj.children)
	return castChildren
}

func (obj *Object3) GetChildrenById(id Id) *Object3 {
	for _, child := range obj.children {
		if child.id == id {
			return child
		} else {
			return child.GetChildrenById(id)
		}
	}
	return nil
}

func (obj *Object3) GetChildrenByName(name string) *Object3 {
	for _, child := range obj.children {
		if child.Name == name {
			return child
		} else {
			return child.GetChildrenByName(name)
		}
	}
	return nil
}

func (obj *Object3) GetWorldPosition() *math.Vector3 {
	target := math.NewDefaultVector3()
	obj.CopyWorldPosition(target)
	return target
}

func (obj *Object3) CopyWorldPosition(target *math.Vector3) {
	obj.UpdateMatrixWorld(true)
	target.SetFromMatrixPosition(&obj.matrixWorld)
}

func (obj *Object3) GetWorldQuaternion() *math.Quaternion {
	target := math.NewDefaultQuaternion()
	obj.CopyWorldQuaternion(target)
	return target
}

func (obj *Object3) CopyWorldQuaternion(target *math.Quaternion) {
	position := math.NewDefaultVector3()
	scale := math.NewDefaultVector3()

	obj.UpdateMatrixWorld(true)

	obj.matrixWorld.DeCompose(position, target, scale)
}

func (obj *Object3) GetWorldScale() *math.Vector3 {
	target := math.NewDefaultVector3()
	obj.CopyWorldScale(target)
	return target
}

func (obj *Object3) CopyWorldScale(target *math.Vector3) {
	position := math.NewDefaultVector3()
	quaternion := math.NewDefaultQuaternion()

	obj.UpdateMatrixWorld(true)

	obj.matrixWorld.DeCompose(position, quaternion, target)
}

func (obj *Object3) GetWorldDirection() *math.Vector3 {
	target := math.NewDefaultVector3()
	obj.CopyWorldDirection(target)
	return target
}

func (obj *Object3) CopyWorldDirection(target *math.Vector3) {
	obj.UpdateMatrixWorld(true)
	e := obj.matrixWorld.GetElements()
	target.Set(e[8], e[9], e[10])
	target.Normalize()
}

func (obj *Object3) TraverseIterator(onlyVisible bool) *object3Iterator {
	return &object3Iterator{
		current: -1,
		data:    obj.allChildren(onlyVisible),
	}
}

func (obj *Object3) allChildren(onlyVisible bool) []*Object3 {
	var ch []*Object3

	if onlyVisible {
		ch := make([]*Object3, 0)
		for _, child := range ch {
			if child.Visible {
				ch = append(ch, child)
			}
		}
	} else {
		ch := make([]*Object3, len(obj.children))
		copy(ch, obj.children)
	}

	for _, child := range obj.children {
		ch = append(ch, child.allChildren(onlyVisible)...)
	}

	return ch
}

func (obj *Object3) TraverseAncestorsIterator() *object3Iterator {
	return &object3Iterator{
		current: -1,
		data:    obj.allParents(),
	}
}

func (obj *Object3) allParents() []*Object3 {
	if obj.parent == nil {
		return []*Object3{}
	} else {
		return append(obj.parent.allParents(), obj.parent)
	}
}

func (obj *Object3) Clone(recursive bool) *Object3 {
	no := NewObject()
	no.Copy(obj, recursive)
	return no
}

// Recursive = TRUE by default
func (obj *Object3) Copy(source *Object3, recursive bool) {
	obj.Name = source.Name
	obj.up = *source.up.Clone()

	obj.Position = *source.Position.Clone()
	obj.quaternion = *source.quaternion.Clone()
	obj.Scale = *source.Scale.Clone()

	obj.matrix = *source.matrix.Clone()
	obj.matrixWorld = *source.matrixWorld.Clone()

	obj.matrixAutoUpdate = source.matrixAutoUpdate
	obj.matrixWorldNeedsUpdate = source.matrixWorldNeedsUpdate

	obj.layers.mask = source.layers.mask
	obj.Visible = source.Visible

	obj.castShadow = source.castShadow
	obj.receiveShadow = source.receiveShadow

	obj.frustumCulled = source.frustumCulled
	obj.RenderOrder = source.RenderOrder

	if recursive {
		for _, child := range source.children {
			obj.Add(child.Clone(recursive))
		}
	}
}

func (obj *Object3) UpdateMatrix() {
	obj.matrix.Compose(&obj.Position, &obj.quaternion, &obj.Scale)
	obj.matrixWorldNeedsUpdate = true
}

func (obj *Object3) UpdateMatrixWorld(force bool) {
	if obj.matrixAutoUpdate {
		obj.UpdateMatrix()
	}

	if obj.matrixWorldNeedsUpdate || force {

		if obj.parent == nil {
			obj.matrixWorld.Copy(&obj.matrix)
		} else {
			obj.matrixWorld.MultiplyMatrices(&obj.parent.matrixWorld, &obj.matrix)
		}

		obj.matrixWorldNeedsUpdate = false
		force = true
	}

	// update children
	for _, child := range obj.children {
		child.UpdateMatrixWorld(force)
	}
}

func (obj *Object3) updateWorldMatrix(updateParents, updateChildren bool) {
	if updateParents == true && obj.parent != nil {
		obj.parent.updateWorldMatrix(true, false)
	}

	if obj.matrixAutoUpdate {
		obj.UpdateMatrix()
	}

	if obj.parent == nil {
		obj.matrixWorld.Copy(&obj.matrix)
	} else {
		obj.matrixWorld.MultiplyMatrices(&obj.parent.matrixWorld, &obj.matrix)
	}

	// update children
	if updateChildren == true {
		for i := 0; i < len(obj.children); i++ {
			obj.children[i].updateWorldMatrix(false, true)
		}
	}
}

func (obj *Object3) GetMatrixWorld() *math.Matrix4 {
	return &obj.matrixWorld
}

type object3Iterator struct {
	current int
	data    []*Object3
}

func (it *object3Iterator) Value() *Object3 {
	return it.data[it.current]
}

func (it *object3Iterator) Next() bool {
	it.current++
	if it.current >= len(it.data) {
		return false
	}
	return true
}
