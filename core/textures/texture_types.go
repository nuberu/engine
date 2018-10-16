package textures

// TODO: temporal name

type MappingMode int

const (
	UVMapping                        MappingMode = 300
	CubeReflectionMapping            MappingMode = 301
	CubeRefractionMapping            MappingMode = 302
	EquirectangularReflectionMapping MappingMode = 303
	EquirectangularRefractionMapping MappingMode = 304
	SphericalReflectionMapping       MappingMode = 305
	CubeUVReflectionMapping          MappingMode = 306
	CubeUVRefractionMapping          MappingMode = 307
)

type WrappingMode int

const (
	RepeatWrapping         WrappingMode = 1000
	ClampToEdgeWrapping    WrappingMode = 1001
	MirroredRepeatWrapping WrappingMode = 1002
)

type MagFilter int

const (
	NearestMagFilter MagFilter = 1003
	LinearMagFilter  MagFilter = 1006
)

type MinFilter int

const (
	NearestMinFilter           MinFilter = 1003
	NearestMipMapNearestFilter MinFilter = 1004
	NearestMipMapLinearFilter  MinFilter = 1005
	LinearMinFilter            MinFilter = 1006
	LinearMipMapNearestFilter  MinFilter = 1007
	LinearMipMapLinearFilter   MinFilter = 1008
)

type Type int

const (
	UnsignedByteType      Type = 1009
	ByteType              Type = 1010
	ShortType             Type = 1011
	UnsignedShortType     Type = 1012
	IntType               Type = 1013
	UnsignedIntType       Type = 1014
	FloatType             Type = 1015
	HalfFloatType         Type = 1016
	UnsignedShort4444Type Type = 1017
	UnsignedShort5551Type Type = 1018
	UnsignedShort565Type  Type = 1019
	UnsignedInt248Type    Type = 1020
)

type Format int

const (
	AlphaFormat          Format = 2021
	RGBFormat            Format = 1022
	RGBAFormat           Format = 1023
	LuminanceFormat      Format = 1024
	LuminanceAlphaFormat Format = 1025
	RGBEFormat           Format = RGBAFormat
	DepthFormat          Format = 1026
	DepthStencilFormat   Format = 1027
	RedFormat            Format = 1028
)

type Encoding int

const (
	LinearEncoding    Encoding = 3000
	sRGBEncoding      Encoding = 3001
	GammaEncoding     Encoding = 3007
	RGBEEncoding      Encoding = 3002
	LogLuvEncoding    Encoding = 3003
	RGBM7Encoding     Encoding = 3004
	RGBM16Encoding    Encoding = 3005
	RGBDEncoding      Encoding = 3006
	BasicDepthPacking Encoding = 3200
	RGBADepthPacking  Encoding = 3201
)
