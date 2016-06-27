package critter

type Stats struct {
	Fat, Grain, Water float64 // Weighable
}

func (s *Stats) Weigh() float64 {
	return s.Fat + s.Grain + s.Water
}
