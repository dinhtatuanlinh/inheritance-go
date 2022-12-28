package main

import "inheritance-go/creature"
import "inheritance-go/flight-creature"

func main() {
	Creature := creature.NewCreature("fish", true)
	Creature.Dump()
	FlightCreature := flight_creature.NewFlightCreature("bird", true, 2)
	FlightCreature.Dump()
	FlightCreature.Fly()
}
