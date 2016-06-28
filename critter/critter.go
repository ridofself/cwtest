package critter

import (
	"github.com/ridthyself/cwtest/item"
	"github.com/ridthyself/cwtest/measure"
)

type Stats struct {
	Name             string
	Fat, Grit, Water float64
	Inventory        item.Group
}

func (own *Stats) Weigh() float64 {
	return (own.Fat + own.Grit + own.Water +
		measure.Weight(&own.Inventory))
}

type Group []*Stats // identical in package: item

func (members Group) Weigh() float64 { // identical in package: item
	var Weight float64
	for _, member := range members {
		Weight += measure.Weight(member)
	}
	return Weight
}
