package constant

type Encoding uint8

const (
	LinearEncoding    Encoding = 1
	sRGBEncoding      Encoding = 2
	RGBEEncoding      Encoding = 3
	LogLuvEncoding    Encoding = 4
	RGBM7Encoding     Encoding = 5
	RGBM16Encoding    Encoding = 6
	RGBDEncoding      Encoding = 7
	GammaEncoding     Encoding = 8
	BasicDepthPacking Encoding = 20
	RGBADepthPacking  Encoding = 21
)
