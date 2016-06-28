package critter

import (
	"github.com/ridthyself/cwtest/item"
	"github.com/ridthyself/cwtest/measure"
)

type New struct {
	Fat, Grit, Water float64
	Inventory        item.Group
}

func (c *New) Weigh() float64 {
	return c.Fat + c.Grit + c.Water + measure.Weight(&c.Inventory)
}
