package flight_creature

import (
	"fmt"
	"inheritance-go/creature"
)

type FlightCreature struct {
	creature.Creature // inherit status of creature
	NumberOfWings     int
}

func (f *FlightCreature) Fly() {
	fmt.Println("flying..")
}

// inherite method Dump of creature stuct
func (f *FlightCreature) Dump() {
	fmt.Printf("Name: '%s', Real: %t, Number of wings: '%d' \n", f.Name, f.Real, f.NumberOfWings)
}

func NewFlightCreature(name string, real bool, numberofWings int) *FlightCreature {
	return &FlightCreature{
		creature.Creature{
			Name: name,
			Real: real,
		},
		numberofWings,
	}
}
