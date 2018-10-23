package core

import "github.com/tokkenno/seed/math"

type Intersection struct {
	distance float32 // Distance between the origin of the ray and the intersection
	point math.Vector3 // Point of intersection, in world coordinates
	face Face3 // Intersected face
	faceIndex uint // Index of the intersected face
	object *Object3 // The intersected object
	uv math.Vector2 // U,V coordinates at point of intersection
}

type SortableIntersections []Intersection

func (s SortableIntersections) Len() int {
	return len(s)
}
func (s SortableIntersections) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortableIntersections) Less(i, j int) bool {
	return s[i].distance < s[j].distance
}