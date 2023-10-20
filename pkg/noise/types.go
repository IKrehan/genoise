package noise

type NoiseType uint8

const (
	Hash NoiseType = iota
	Value
	Perlin
	Worley
	Voronoi
	Fractal
)

func (n NoiseType) String() string {
	switch n {
	case Hash:
		return "hash"
	case Value:
		return "value"
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
