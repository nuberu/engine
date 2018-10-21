package core

import "github.com/tokkenno/seed/core/constant"

type Options struct {
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

func NewTexture(image *Image, mapping constant.MappingMode, wrapS constant.Wrapping, wrapT constant.Wrapping,
	magFilter constant.MagFilter, minFilter constant.MinFilter, format constant.Format, ttype constant.Type, anisotropy uint, encoding constant.Encoding) *Texture {
	return nil
}

func (texture *Texture) Clone() *Texture {
	// TODO
	return nil
}
