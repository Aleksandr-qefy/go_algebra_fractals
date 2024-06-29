package clrplot

type Color struct {
	R, G, B uint8
}

func (c Color) plus(another Color) Color {
	return Color{
		R: c.R + another.R,
		G: c.G + another.G,
		B: c.B + another.B,
	}
}

func (c Color) portion(dible, der int) Color {
	div := func(c uint8) uint8 {
		return uint8(int(c) * dible / der)
	}
	return Color{
		R: div(c.R),
		G: div(c.G),
		B: div(c.B),
	}
}
