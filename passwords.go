package people

import (
	"math/rand"
)

func Password() string {
	return Words[rand.Int31n(int32(len(Words)))]
}
