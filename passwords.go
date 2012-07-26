package people

import (
	"math/rand"
)

// Return a random word which could conceivably be someone's password. This is clearly not
// a secure password generator.
func Password() string {
	return Words[rand.Int31n(int32(len(Words)))]
}
