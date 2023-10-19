package noise

import "math/rand"

func discreteStrategy(x, y int) (float32, error) {
	return rand.Float32(), nil
}
