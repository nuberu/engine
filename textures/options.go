package textures

type Options struct {
	WrapS           WrappingMode
	WrapT           WrappingMode
	MagFilter       MagFilter
	MinFilter       MinFilter
	Format          Format
	TextureType     Type
	Anisotropy      uint
	Encoding        Encoding
}
