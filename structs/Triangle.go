package structs

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}
