package people

import (
	"fmt"
	"math/rand"
)

// Produce a phone number of the format '+15556667777'
func Phone() string {
	return fmt.Sprintf("+1%3d%3d%04d", 200+rand.Int31n(800), 200+rand.Int31n(800), rand.Int31n(10000))
}
