package main

import (
	"fmt"
	"github.com/ridthyself/cwtest/critter"
	"github.com/ridthyself/cwtest/item"
	"github.com/ridthyself/cwtest/measure"
)

func main() {
	Helm := &item.Stats{"Cap of Dismay", 0.8}
	Sword := &item.Stats{"Pinner", 2.2}
	Shield := &item.Stats{"Prot", 6.5}
	Inventory := item.Group{Sword, Shield, Helm}
	Wedge := &critter.Stats{"Wedge", 14, 32, 98, Inventory}
	Biggs := &critter.Stats{"Biggs", 10, 29, 90, Inventory}
	fmt.Println(measure.Weight(Wedge))
	fmt.Println(measure.Weight(Biggs))
	fmt.Println(measure.Weight(Inventory))
	fmt.Println(Biggs)
	Party := &critter.Group{Biggs, Wedge}
	fmt.Println(measure.Weight(Party))
}
