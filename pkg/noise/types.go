package noise

type NoiseType uint8

const (
	Discrete NoiseType = iota
	Bicubic
	Perlin
	Simplex
)

func (n NoiseType) String() string {
	switch n {
	case Discrete:
		return "discrete"
	case Bicubic:
		return "tricubic"
	case Perlin:
		return "perlin"
	case Simplex:
		return "simplex"
	default:
		return "unknown"
	}
}
