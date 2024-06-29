package usertypes

import (
	"math"
	"math/cmplx"
)

type Fractal struct {
	MaxX, MinX float64
	MinY, MaxY float64
	Function   func(complex128, complex128) complex128
}

var Mandelbrot = Fractal{
	MinX: -3, MaxX: 2,
	MinY: -1.5, MaxY: 1.5,
	Function: func(z, c complex128) complex128 {
		return cmplx.Pow(z, 2) + c
	},
}

var BurningShip = Fractal{
	MinX: -3, MaxX: 2,
	MinY: -2, MaxY: 1,
	Function: func(z, c complex128) complex128 {
		return cmplx.Pow(complex(math.Abs(real(z)), math.Abs(imag(z))), 2) + c
	},
}
