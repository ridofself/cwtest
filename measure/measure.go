package measure

type Weighable interface {
	Weigh() float64
}

func Weight(w Weighable) float64 {
	return w.Weigh()
}
