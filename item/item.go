package item

import (
	"github.com/ridthyself/cwtest/measure"
)

type New struct {
	Name   string
	Weight float64
}

func (item *New) Weigh() float64 {
	return item.Weight
}

type Group []New

func (items *Group) Weigh() float64 {
	var Weight float64
	for _, item := range *items {
		Weight += measure.Weight(&item)
	}
	return Weight
}
