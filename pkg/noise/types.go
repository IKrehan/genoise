package noise

type NoiseType uint8

const (
	WhiteNoise NoiseType = iota
	Perlin
	Worley
	Voronoi
	Fractal
)

func (n NoiseType) String() string {
	switch n {
	case WhiteNoise:
		return "white noise"
	case Perlin:
		return "perlin"
	case Worley:
		return "worley"
	case Voronoi:
		return "voronoi"
	case Fractal:
		return "fractal"
	default:
		return "unknown"
	}
}
