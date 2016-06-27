package main

import (
	"fmt"
	"github.com/ridthyself/cwtest/critter"
	"github.com/ridthyself/cwtest/measure"
)

func main() {
	Wedge := &critter.Stats{14, 32, 98}
	fmt.Println(Wedge, measure.Weight(Wedge))
}
