package noise

import "math/rand"

func hashStrategy(x, y int) (float32, error) {
	return rand.Float32(), nil
}
