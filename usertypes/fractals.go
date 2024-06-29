package usertypes

import (
	"math"
	"math/cmplx"
)

type Fractal struct {
	maxX, minX float64
	minY, maxY float64
	function   func(complex128, complex128) complex128
}

var Mandelbrot = Fractal{
	minX: -3, maxX: 2,
	minY: -1.5, maxY: 1.5,
	function: func(z, c complex128) complex128 {
		return cmplx.Pow(z, 2) + c
	},
}

var BurningShip = Fractal{
	minX: -3, maxX: 2,
	minY: -2, maxY: 1,
	function: func(z, c complex128) complex128 {
		return cmplx.Pow(complex(math.Abs(real(z)), math.Abs(imag(z))), 2) + c
	},
}
