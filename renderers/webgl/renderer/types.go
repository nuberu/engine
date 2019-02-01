package renderer

import "syscall/js"

type GLTypes struct {
	StaticDraw         js.Value
	ArrayBuffer        js.Value
	ElementArrayBuffer js.Value
	VertexShader       js.Value
	FragmentShader     js.Value
	Float              js.Value
	DepthTest          js.Value
	ColorBufferBit     js.Value
	Triangles          js.Value
	UnsignedShort      js.Value
}

func (types *GLTypes) New(glContext js.Value) {
	types.StaticDraw = glContext.Get("STATIC_DRAW")
	types.ArrayBuffer = glContext.Get("ARRAY_BUFFER")
	types.ElementArrayBuffer = glContext.Get("ELEMENT_ARRAY_BUFFER")
	types.VertexShader = glContext.Get("VERTEX_SHADER")
	types.FragmentShader = glContext.Get("FRAGMENT_SHADER")
	types.Float = glContext.Get("FLOAT")
	types.DepthTest = glContext.Get("DEPTH_TEST")
	types.ColorBufferBit = glContext.Get("COLOR_BUFFER_BIT")
	types.Triangles = glContext.Get("TRIANGLES")
	types.UnsignedShort = glContext.Get("UNSIGNED_SHORT")
}
