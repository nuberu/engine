package constant

type MagFilter uint8

const (
	NearestMagFilter MagFilter = 3
	LinearMagFilter  MagFilter = 6
)

type MinFilter uint8

const (
	NearestMinFilter           MinFilter = 3
	NearestMipMapNearestFilter MinFilter = 4
	NearestMipMapLinearFilter  MinFilter = 5
	LinearMinFilter            MinFilter = 6
	LinearMipMapNearestFilter  MinFilter = 7
	LinearMipMapLinearFilter   MinFilter = 8
)
