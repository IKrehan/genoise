package noise

import (
	"math"
)

func hashStrategy(x, y int) (float64, error) {
	xmod, ymod, mod := 0.129898, 0.78233, 43758.5453
	dot := math.Sin(float64(x)*xmod+float64(y)*ymod) * mod

	_, res := math.Modf(math.Abs(dot))
	return res, nil
}
