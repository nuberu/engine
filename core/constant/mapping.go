package constant

type Mapping int

const (
	UVMapping                        Mapping = 1
	CubeReflectionMapping            Mapping = 2
	CubeRefractionMapping            Mapping = 3
	EquirectangularReflectionMapping Mapping = 4
	EquirectangularRefractionMapping Mapping = 5
	SphericalReflectionMapping       Mapping = 6
	CubeUVReflectionMapping          Mapping = 7
	CubeUVRefractionMapping          Mapping = 8
)

type ToneMapping int

const (
	NoToneMapping         ToneMapping = 0
	LinearToneMapping     ToneMapping = 1
	ReinhardToneMapping   ToneMapping = 2
	Uncharted2ToneMapping ToneMapping = 3
	CineonToneMapping     ToneMapping = 4
)