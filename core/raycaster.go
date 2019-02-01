package core

import (
	"github.com/nuberu/engine/math"
	nativeMath "math"
	"sort"
)

type RaycastObject interface {
	IsVisible() bool
	RayCast(*Raycaster) []Intersection
	GetChildren() []*Object3
}

type RaycastCamera interface {
	GetRay(coordinates *math.Vector2) *math.Ray
}

type Raycaster struct {
	ray math.Ray
	near float32
	far float32
	linePrecision uint
}

func NewDefaultRaycaster(origin *math.Vector3, direction *math.Vector3) *Raycaster {
	return NewRaycaster(origin, direction, 0, float32(nativeMath.Inf(1)), 1)
}

func NewRaycaster(origin *math.Vector3, direction *math.Vector3, near float32, far float32, linePrecision uint) *Raycaster {
	return &Raycaster{
		ray: *math.NewRay(origin, direction),
		near: near,
		far: far,
		linePrecision: linePrecision,
	}
}

func (ray *Raycaster) Set(origin *math.Vector3, direction *math.Vector3) {
	ray.ray.Set(origin, direction)
}

func (ray *Raycaster) SetFromCamera(coordinates *math.Vector2, camera RaycastCamera) {
	ray.ray.Copy(camera.GetRay(coordinates))
}

func (ray *Raycaster) IntersectObject(object interface{RaycastObject}, recursive bool) []Intersection {
	if object.IsVisible() == false {
		return []Intersection{}
	}

	intersections := object.RayCast(ray)

	if recursive {
		for _, child := range object.GetChildren() {
			// Raycast the child objects that implements the RaycastObject interface
			var anonChild interface{} = child
			rayCastChild, ok := anonChild.(RaycastObject)
			if ok {
				intersections = append(intersections, ray.IntersectObject(rayCastChild, recursive)...)
			}
		}
	}

	return intersections
}

func (ray *Raycaster) IntersectObjects(objects []interface{RaycastObject}, recursive bool) []Intersection {
	var intersections []Intersection
	for _, object := range objects {
		intersections = append(intersections, ray.IntersectObject(object, recursive)...)
	}
	sort.Sort(SortableIntersections(intersections))
	return intersections
}