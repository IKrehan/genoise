package noise

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
)

type getPositionValueStrategy func(int, int) (float64, error)

type Noise struct {
	noiseType        NoiseType
	GetPositionValue getPositionValueStrategy
}

func New(noiseType NoiseType) *Noise {
	var strategy getPositionValueStrategy

	switch noiseType {
	case Hash:
		strategy = hashStrategy
	case Perlin:
		strategy = perlinStrategy
	case Worley:
		strategy = hashStrategy
	case Voronoi:
		strategy = hashStrategy
	case Fractal:
		strategy = hashStrategy
	default:
		strategy = hashStrategy
	}

	return &Noise{
		GetPositionValue: strategy,
	}
}

func (noise Noise) Type() string {
	return noise.noiseType.String()
}

func (noise Noise) GenerateMatrix(width, height int) [][]float64 {
	matrix := make([][]float64, width)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			value, err := noise.GetPositionValue(x, y)
			if err != nil {
				log.Fatal("Invalid Coords")
			}
			matrix[x] = append(matrix[x], value)
		}
	}

	return matrix
}

func (noise Noise) GenerateImage(fileName string, width, height int) error {
	img := image.NewGray(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{width, height},
		},
	)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			value, err := noise.GetPositionValue(x, y)
			if err != nil {
				return errors.New("Failed to get Position Value: " + err.Error())
			}

			img.Set(x, y, color.Gray{uint8(value * 255)})
		}
	}

	if !strings.Contains(fileName, ".png") {
		fileName = fileName + ".png"
	}

	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("Failed to create file: " + fileName)
	}

	err = png.Encode(file, img)
	if err != nil {
		return errors.New("Failed to encode image: " + err.Error())
	}

	return nil
}
