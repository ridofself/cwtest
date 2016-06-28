package main

import (
	"fmt"
	"github.com/ridthyself/cwtest/critter"
	"github.com/ridthyself/cwtest/item"
	"github.com/ridthyself/cwtest/measure"
)

func main() {
	Sword := &item.New{"Glamdrig strom bringer", 2.1}
	Shield := &item.New{"blockhead", 5.3}
	Inventory := &item.Group{*Sword, *Shield}
	Wedge := &critter.New{14, 32, 98, *Inventory}
	fmt.Println(Wedge)
	fmt.Println(Inventory)
	fmt.Println(measure.Weight(Wedge))
	fmt.Println(measure.Weight(Sword))
	fmt.Println(measure.Weight(Inventory))
}
