package gl

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tokkenno/seed/math"
)

type window struct {
	width uint32
	height uint32
	title string
	windowPtr *glfw.Window
}

func (window *window) GetSize() *math.Vector2 {
	return math.NewVector2(float32(window.width), float32(window.height))
}

func (window *window) SetSize(size *math.Vector2) {
	window.width = uint32(size.X)
	window.height = uint32(size.Y)
}

func (window *window) SetTitle(title string) {
	window.title = title
}

func NewWindow(size *math.Vector2, title string) (*window, error) {
	window := &window{}

	err := glfw.Init()
	if err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window.windowPtr, err = glfw.CreateWindow(int(size.X), int(size.Y), title, nil, nil)
	if err != nil {
		return nil, err
	}
	window.windowPtr.MakeContextCurrent()

	return window, nil
}
