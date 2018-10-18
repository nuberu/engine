package textures

type Texture struct {
	id uint64
	GenerateMipmaps bool
}

func NewTexture(image *Image, mapping MappingMode, wrapS WrappingMode, wrapT WrappingMode, magFilter MagFilter, minFilter MinFilter, format Format, ttype Type, anisotropy uint, encoding Encoding) *Texture {
	return nil
}

func (texture *Texture) Clone() *Texture {
	// TODO
	return nil
}