package constant

type Format uint8

const (
	AlphaFormat          Format = 1
	RGBFormat            Format = 2
	RGBAFormat           Format = 3
	LuminanceFormat      Format = 4
	LuminanceAlphaFormat Format = 5
	RGBEFormat           Format = RGBAFormat
	DepthFormat          Format = 6
	DepthStencilFormat   Format = 7
	RedFormat            Format = 8

	RGB_S3TC_DXT1_Format     Format = 20
	RGBA_S3TC_DXT1_Format    Format = 21
	RGBA_S3TC_DXT3_Format    Format = 22
	RGBA_S3TC_DXT5_Format    Format = 23
	RGB_PVRTC_4BPPV1_Format  Format = 24
	RGB_PVRTC_2BPPV1_Format  Format = 25
	RGBA_PVRTC_4BPPV1_Format Format = 26
	RGBA_PVRTC_2BPPV1_Format Format = 27
	RGB_ETC1_Format          Format = 28
	RGBA_ASTC_4x4_Format     Format = 29
	RGBA_ASTC_5x4_Format     Format = 30
	RGBA_ASTC_5x5_Format     Format = 31
	RGBA_ASTC_6x5_Format     Format = 32
	RGBA_ASTC_6x6_Format     Format = 33
	RGBA_ASTC_8x5_Format     Format = 34
	RGBA_ASTC_8x6_Format     Format = 35
	RGBA_ASTC_8x8_Format     Format = 36
	RGBA_ASTC_10x5_Format    Format = 37
	RGBA_ASTC_10x6_Format    Format = 38
	RGBA_ASTC_10x8_Format    Format = 39
	RGBA_ASTC_10x10_Format   Format = 40
	RGBA_ASTC_12x10_Format   Format = 41
	RGBA_ASTC_12x12_Format   Format = 42
)
