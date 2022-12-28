package creature

import "fmt"

type Creature struct {
	Name string
	Real bool
}

func (c Creature) Dump() {
	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

func NewCreature(name string, real bool) *Creature {
	return &Creature{
		Name: name,
		Real: real,
	}
}
