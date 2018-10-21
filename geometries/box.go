package geometries

import (
	"github.com/tokkenno/seed/core/buffer"
	"github.com/tokkenno/seed/core/constant"
	"github.com/tokkenno/seed/core/geometry"
	"github.com/tokkenno/seed/math"
)

type Box struct {
	geometry.Basic

	width          float32
	height         float32
	depth          float32
	widthSegments  uint
	heightSegments uint
	depthSegments  uint
}

func NewBox(width, height, depth float32) *Box {
	return NewBoxWithSegments(width, height, depth, 1, 1, 1)
}

func NewBoxWithSegments(width, height, depth float32, widthSegments, heightSegments, depthSegments uint) *Box {
	return &Box{}
}

type BoxBuffer struct {
	Box
	geometry.Buffer

	numberOfVertices uint
	groupStart       uint

	indices *buffer.Float32
	vertices *buffer.Float32
	normals *buffer.Float32
	uvs *buffer.Float32
}

func newBoxBuffer(width, height, depth float32, widthSegments, heightSegments, depthSegments uint) *BoxBuffer {
	box := &BoxBuffer{
		Box:              *NewBoxWithSegments(width, height, depth, widthSegments, heightSegments, depthSegments),
		indices:          buffer.NewFloat32Dynamic(),
		vertices:         buffer.NewFloat32Dynamic(),
		normals:          buffer.NewFloat32Dynamic(),
		uvs:              buffer.NewFloat32Dynamic(),
		numberOfVertices: 0,
		groupStart:       0,
	}

	box.buildPlane(constant.AxisZ, constant.AxisY, constant.AxisX, -1, -1, depth, height, width, depthSegments, heightSegments, 0)   // px
	box.buildPlane(constant.AxisZ, constant.AxisY, constant.AxisX, 1, - 1, depth, height, - width, depthSegments, heightSegments, 1) // nx
	box.buildPlane(constant.AxisX, constant.AxisZ, constant.AxisY, 1, 1, width, depth, height, widthSegments, depthSegments, 2)      // py
	box.buildPlane(constant.AxisX, constant.AxisZ, constant.AxisY, 1, - 1, width, depth, - height, widthSegments, depthSegments, 3)  // ny
	box.buildPlane(constant.AxisX, constant.AxisY, constant.AxisZ, 1, - 1, width, height, depth, widthSegments, heightSegments, 4)   // pz
	box.buildPlane(constant.AxisX, constant.AxisY, constant.AxisZ, -1, -1, width, height, - depth, widthSegments, heightSegments, 5) // nz

	box.vertices.SetDynamic(false)
	box.normals.SetDynamic(false)
	box.uvs.SetDynamic(false)

	// build geometry
	//box.setIndex(indices)
	return box
}

func (box *BoxBuffer) buildPlane(u, v, w constant.Axis, uDir, vDir, width, height, depth float32, gridX, gridY, materialIndex uint) {
	segmentWidth := width / float32(gridX)
	segmentHeight := height / float32(gridY)

	widthHalf := width / 2
	heightHalf := height / 2
	depthHalf := depth / 2

	gridX1 := gridX + 1
	gridY1 := gridY + 1

	vertexCounter := uint(0)
	groupCount := uint(0)
	vector := math.NewVector3(0, 0, 0)

	// generate vertices, normals and uvs
	for iy := uint(0); iy < gridY1; iy++ {
		y := float32(iy)*segmentHeight - heightHalf

		for ix := uint(0); ix < gridX1; ix ++ {
			x := float32(ix)*segmentWidth - widthHalf
			depthVal := 1
			if depth <= 0 {
				depthVal = -1
			}

			// set values to correct vector component
			v1 := x * uDir
			v2 := y * vDir

			switch u {
			case constant.AxisX:
				vector.SetX(v1)
				break
			case constant.AxisY:
				vector.SetY(v1)
				break
			case constant.AxisZ:
				vector.SetZ(v1)
				break
			}

			switch v {
			case constant.AxisX:
				vector.SetX(v2)
				break
			case constant.AxisY:
				vector.SetY(v2)
				break
			case constant.AxisZ:
				vector.SetZ(v2)
				break
			}

			switch w {
			case constant.AxisX:
				vector.SetX(depthHalf)
				break
			case constant.AxisY:
				vector.SetY(depthHalf)
				break
			case constant.AxisZ:
				vector.SetZ(depthHalf)
				break
			}

			// now apply vector to vertex buffer
			box.vertices.SetVector3(vertexCounter, vector)

			// set values to correct vector component
			switch u {
			case constant.AxisX:
				vector.SetX(0)
				break
			case constant.AxisY:
				vector.SetY(0)
				break
			case constant.AxisZ:
				vector.SetZ(0)
				break
			}

			switch v {
			case constant.AxisX:
				vector.SetX(0)
				break
			case constant.AxisY:
				vector.SetY(0)
				break
			case constant.AxisZ:
				vector.SetZ(0)
				break
			}

			switch w {
			case constant.AxisX:
				vector.SetX(float32(depthVal))
				break
			case constant.AxisY:
				vector.SetY(float32(depthVal))
				break
			case constant.AxisZ:
				vector.SetZ(float32(depthVal))
				break
			}

			// now apply vector to normal buffer
			box.normals.SetVector3(vertexCounter, vector)

			// uvs
			box.uvs.Set(vertexCounter, float32(ix/gridX))
			box.uvs.Set(vertexCounter+1, float32(1-(iy/gridY)))

			// counters
			vertexCounter += 1
		}
	}

	// indices
	// 1. you need three indices to draw a single face
	// 2. a single segment consists of two faces
	// 3. so we need to generate six (2*3) indices per segment
	for iy := uint(0); iy < gridY; iy++ {
		for ix := uint(0); ix < gridX; ix ++ {
			a := float32(box.numberOfVertices + ix + gridX1*iy)
			b := float32(box.numberOfVertices + ix + gridX1*(iy+1))
			c := float32(box.numberOfVertices + (ix + 1) + gridX1*(iy+1))
			d := float32(box.numberOfVertices + (ix + 1) + gridX1*iy)

			// faces
			box.indices.Set(groupCount, a)
			box.indices.Set(groupCount+1, b)
			box.indices.Set(groupCount+2, d)
			box.indices.Set(groupCount+3, b)
			box.indices.Set(groupCount+4, c)
			box.indices.Set(groupCount+5, d)

			// increase counter
			groupCount += 6
		}
	}

	// add a group to the geometry. this will ensure multi material support
	box.AddGroupIndex(box.groupStart, groupCount, materialIndex)

	// calculate new start value for groups
	box.groupStart += groupCount

	// update total number of vertices
	box.numberOfVertices += vertexCounter
}
