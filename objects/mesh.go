package objects

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/core/geometry"
	"github.com/tokkenno/seed/core/material"
	"github.com/tokkenno/seed/render/draw"
)

type Mesh struct {
	core.Object3

	geometry *geometry.Geometry
	material *core.Material
	drawMode draw.Mode

	morphTargetInfluences []int // FIXME: type
	morphTargetDictionary map[int]int // FIXME: type
}

func (mesh *Mesh) SetDrawMode(mode draw.Mode) {
	mesh.drawMode = mode
}

func (mesh *Mesh) Copy(source *Mesh) {
	mesh.Object3.Copy(&source.Object3, true)

	mesh.drawMode = source.drawMode

	mesh.morphTargetInfluences = make([]int, len(source.morphTargetInfluences))
	copy(mesh.morphTargetInfluences, source.morphTargetInfluences)

	mesh.morphTargetDictionary = make(map[int]int)
	for key, value := range source.morphTargetDictionary {
		mesh.morphTargetDictionary[key] = value
	}
}