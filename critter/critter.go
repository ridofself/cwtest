package critter

import (
	"github.com/ridthyself/cwtest/item"
	"github.com/ridthyself/cwtest/measure"
)

type Stats struct {
	Fat, Grit, Water float64
	Inventory        item.Inventory
}

func (s *Stats) Weigh() float64 {
	return s.Fat + s.Grit + s.Water + measure.Weight(&s.Inventory)
}
