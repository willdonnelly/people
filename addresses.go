package people

import (
	"fmt"
	"math/rand"
	"strings"
)

func Address() string {
	number := rand.Int31n(10000)
	street := strings.Title(Words[rand.Int31n(int32(len(Words)))])
	sType := StreetType[rand.Int31n(int32(len(StreetType)))]
	state := State[rand.Int31n(int32(len(State)))]
	return fmt.Sprintf("%d %s %s\nExampleTown, %s", number, street, sType, state)
}

var StreetType = []string{"Ave", "St", "Ct", "Pl"}

var State = []string{
	"AL", "AK", "AZ", "AR", "CA",
	"CO", "CT", "DE", "FL", "GA",
	"HI", "ID", "IL", "IN", "IA",
	"KS", "KY", "LA", "ME", "MD",
	"MA", "MI", "MN", "MS", "MO",
	"MT", "NE", "NV", "NH", "NJ",
	"NM", "NY", "NC", "ND", "OH",
	"OK", "OR", "PA", "RI", "SC",
	"SD", "TN", "TX", "UT", "VT",
	"VA", "WA", "WV", "WI", "WY",
}
