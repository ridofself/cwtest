package item

import (
	"github.com/ridthyself/cwtest/measure"
)

type Stats struct { // defines the properties of a new item
	Name   string
	Weight float64
}

func (own *Stats) Weigh() float64 {
	return own.Weight
}

type Group []*Stats // identical in package: critter

func (members Group) Weigh() float64 { // identical in package: critter
	var Weight float64
	for _, member := range members {
		Weight += measure.Weight(member)
	}
	return Weight
}
