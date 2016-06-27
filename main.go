package main

import (
	"fmt"
	"github.com/ridthyself/cwtest/critter"
	"github.com/ridthyself/cwtest/item"
	"github.com/ridthyself/cwtest/measure"
)

func main() {
	Sword := &item.Stats{"Glamdrig strom bringer", 2.1}
	Shield := &item.Stats{"blockhead", 5.3}
	Inventory := &item.Inventory{*Sword, *Shield}
	Wedge := &critter.Stats{14, 32, 98, *Inventory}
	fmt.Println(Wedge)
	fmt.Println(Inventory)
	fmt.Println(measure.Weight(Wedge))
	fmt.Println(measure.Weight(Sword))
	fmt.Println(measure.Weight(Inventory))
}
