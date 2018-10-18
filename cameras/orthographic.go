package cameras

type orthographicView struct {
	enabled    bool
	fullWidth  float64
	fullHeight float64
	offsetX    float64
	offsetY    float64
	width      float64
	height     float64
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
	Camera

	zoom float64
	view *orthographicView

	left   float64
	right  float64
	top    float64
	bottom float64
	near   float64
	far    float64
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
	camera.view = source.view.Clone()
}

func (camera *Orthographic) SetViewOffset(fullWidth, fullHeight, x, y, width, height float64) {
	if camera.view == nil {
		camera.view = newOrthographicView()
	}

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
	if camera.view != nil {
		camera.view.enabled = false
	}

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

	if camera.view != nil && camera.view.enabled {

		var zoomW = camera.zoom / (camera.view.width / camera.view.fullWidth);
		var zoomH = camera.zoom / (camera.view.height / camera.view.fullHeight);
		var scaleW = (camera.right - camera.left) / camera.view.width;
		var scaleH = (camera.top - camera.bottom) / camera.view.height;
		left += scaleW * (camera.view.offsetX / zoomW);
		right = left + scaleW*(camera.view.width/zoomW);
		top -= scaleH * (camera.view.offsetY / zoomH);
		bottom = top - scaleH*(camera.view.height/zoomH);
	}

	camera.GetProjectionMatrix().MakeOrthographic(left, right, top, bottom, camera.near, camera.far)

	camera.GetProjectionMatrixInverse().GetInverse(camera.projectionMatrix)
}
