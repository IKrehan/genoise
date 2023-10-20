package noise

import "math/rand"

func whiteNoiseStrategy(x, y int) (float32, error) {
	return rand.Float32(), nil
}
