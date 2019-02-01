package renderer

import "syscall/js"

type Extensions struct {
	glContext js.Value
	cache     map[string]js.Value
}

func NewExtensions(glContext js.Value) *Extensions {
	return &Extensions{
		glContext: glContext,
		cache: make(map[string]js.Value),
	}
}

func (ext *Extensions) Get(name string) js.Value {
	if value, ok := ext.cache[name]; ok {
		return value
	} else {
		switch name {
		case "WEBGL_depth_texture":
			value := ext.glContext.Call("getExtension", "WEBGL_depth_texture")
			if value == js.Undefined() {
				value = ext.glContext.Call("getExtension", "MOZ_WEBGL_depth_texture")
			}
			if value == js.Undefined() {
				value = ext.glContext.Call("getExtension", "WEBKIT_WEBGL_depth_texture")
			}
			ext.cache["WEBGL_depth_texture"] = value
			return value
		case "EXT_texture_filter_anisotropic":
			value := ext.glContext.Call("getExtension", "EXT_texture_filter_anisotropic")
			if value == js.Undefined() {
				value = ext.glContext.Call("getExtension", "MOZ_EXT_texture_filter_anisotropic")
			}
			if value == js.Undefined() {
				value = ext.glContext.Call("getExtension", "WEBKIT_EXT_texture_filter_anisotropic")
			}
			ext.cache["EXT_texture_filter_anisotropic"] = value
			return value
		case "WEBGL_compressed_texture_s3tc":
			value := ext.glContext.Call("getExtension", "WEBGL_compressed_texture_s3tc")
			if value == js.Undefined() {
				value = ext.glContext.Call("getExtension", "MOZ_WEBGL_compressed_texture_s3tc")
			}
			if value == js.Undefined() {
				value = ext.glContext.Call("getExtension", "WEBKIT_WEBGL_compressed_texture_s3tc")
			}
			ext.cache["WEBGL_compressed_texture_s3tc"] = value
			return value
		case "WEBGL_compressed_texture_pvrtc":
			value := ext.glContext.Call("getExtension", "WEBGL_compressed_texture_pvrtc")
			if value == js.Undefined() {
				value = ext.glContext.Call("getExtension", "WEBKIT_WEBGL_compressed_texture_pvrtc")
			}
			ext.cache["WEBGL_compressed_texture_pvrtc"] = value
			return value
		default:
			return ext.glContext.Call("getExtension", name)
		}
	}
}
