package cameras

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/math"
)

type orthographicView struct {
	enabled    bool
	fullWidth  float32
	fullHeight float32
	offsetX    float32
	offsetY    float32
	width      float32
	height     float32
}

func newOrthographicView() *orthographicView {
	return &orthographicView{
		enabled:    true,
		fullWidth:  1,
		fullHeight: 1,
		offsetX:    0,
		offsetY:    0,
		width:      1,
		height:     1,
	}
}

func (ov *orthographicView) Clone() *orthographicView {
	return &orthographicView{
		enabled:    ov.enabled,
		fullWidth:  ov.fullWidth,
		fullHeight: ov.fullHeight,
		offsetX:    ov.offsetX,
		offsetY:    ov.offsetY,
		width:      ov.width,
		height:     ov.height,
	}
}

type Orthographic struct {
	core.Camera

	zoom float32
	view orthographicView

	left   float32
	right  float32
	top    float32
	bottom float32
	near   float32
	far    float32
}

func NewOrthographic(left, top, right, bottom float32, near, far float32) *Orthographic {
	return &Orthographic{
		Camera: *core.NewCamera(),
		zoom:   0.0,
		view:   *newOrthographicView(),
		left:   left,
		right:  right,
		top:    top,
		bottom: bottom,
		near:   near,
		far:    far,
	}
}

func (camera *Orthographic) Copy(source *Orthographic, recursive bool) {
	camera.Camera.Copy(&source.Camera, recursive)

	camera.left = source.left
	camera.right = source.right
	camera.top = source.top
	camera.bottom = source.bottom
	camera.near = source.near
	camera.far = source.far

	camera.zoom = source.zoom
	camera.view = *source.view.Clone()
}

func (camera *Orthographic) SetViewOffset(fullWidth, fullHeight, x, y, width, height float32) {
	camera.view.enabled = true
	camera.view.fullWidth = fullWidth
	camera.view.fullHeight = fullHeight
	camera.view.offsetX = x
	camera.view.offsetY = y
	camera.view.width = width
	camera.view.height = height

	camera.UpdateProjectionMatrix()
}

func (camera *Orthographic) ClearViewOffset() {
	camera.view.enabled = false
	camera.UpdateProjectionMatrix()
}

func (camera *Orthographic) UpdateProjectionMatrix() {
	dx := (camera.right - camera.left) / (2 * camera.zoom)
	dy := (camera.top - camera.bottom) / (2 * camera.zoom)
	cx := (camera.right + camera.left) / 2
	cy := (camera.top + camera.bottom) / 2

	left := cx - dx
	right := cx + dx
	top := cy + dy
	bottom := cy - dy

	if camera.view.enabled {

		var zoomW = camera.zoom / (camera.view.width / camera.view.fullWidth)
		var zoomH = camera.zoom / (camera.view.height / camera.view.fullHeight)
		var scaleW = (camera.right - camera.left) / camera.view.width
		var scaleH = (camera.top - camera.bottom) / camera.view.height
		left += scaleW * (camera.view.offsetX / zoomW)
		right = left + scaleW*(camera.view.width/zoomW)
		top -= scaleH * (camera.view.offsetY / zoomH)
		bottom = top - scaleH*(camera.view.height/zoomH)
	}

	camera.GetProjectionMatrix().MakeOrthographic(left, right, top, bottom, camera.near, camera.far)

	camera.GetProjectionMatrixInverse().SetInverseOf(camera.GetProjectionMatrix(), false)
}

func (camera *Orthographic) GetRay(coordinates *math.Vector2) *math.Ray {
	origin := math.NewVector3(coordinates.X, coordinates.Y,( camera.near + camera.far ) / ( camera.near - camera.far ))
	math.UnProject(origin, &camera.Camera)

	direction := math.NewVector3(0, 0, -1)
	direction.TransformDirection(camera.GetMatrixWorld())

	return math.NewRay(origin, direction)

}