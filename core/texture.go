package core

import (
	"github.com/nuberu/engine/core/constant"
	"github.com/nuberu/engine/textures"
)

type TextureOptions struct {
	WrapS       constant.Wrapping
	WrapT       constant.Wrapping
	MagFilter   constant.MagFilter
	MinFilter   constant.MinFilter
	Format      constant.Format
	TextureType constant.Type
	Anisotropy  uint
	Encoding    constant.Encoding
}

type Texture struct {
	GenerateMipmaps bool
}

func NewTexture(image *textures.Image, mapping constant.Mapping, wrapS constant.Wrapping, wrapT constant.Wrapping,
	magFilter constant.MagFilter, minFilter constant.MinFilter, format constant.Format, ttype constant.Type, anisotropy uint, encoding constant.Encoding) *Texture {
	return nil
}

func (texture *Texture) Clone() *Texture {
	// TODO
	return nil
}
