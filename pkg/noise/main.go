package noise

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type getPositionValueStrategy func(int, int) (float32, error)

type Noise struct {
	noiseType        NoiseType
	GetPositionValue getPositionValueStrategy
}

func New(noiseType NoiseType) *Noise {
	var strategy getPositionValueStrategy

	switch noiseType {
	case Discrete:
		strategy = discreteStrategy
	case Bicubic:
		strategy = discreteStrategy
	case Perlin:
		strategy = discreteStrategy
	case Simplex:
		strategy = discreteStrategy
	}

	return &Noise{
		GetPositionValue: strategy,
	}
}

func (noise Noise) Type() string {
	return noise.noiseType.String()
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
