package item

import (
	"github.com/ridthyself/cwtest/measure"
)

type Stats struct {
	Name   string
	Weight float64
}

func (item *Stats) Weigh() float64 {
	return item.Weight
}

type Inventory []Stats

func (items *Inventory) Weigh() float64 {
	var Weight float64
	for _, item := range *items {
		Weight += measure.Weight(&item)
	}
	return Weight
}
