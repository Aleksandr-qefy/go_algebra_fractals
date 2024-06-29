package usertypes

import (
	cpt "github.com/Aleksandr-qefy/clrplot"
	"math"
)

type Matrix struct {
	mtrx          [][]int
	min, max      int
	height, width int
}

func NewMatrix(height, width int) Matrix {
	mtrx := make([][]int, height)
	for i := 0; i < height; i++ {
		mtrx[i] = make([]int, width)
	}
	return Matrix{
		mtrx:   mtrx,
		height: height,
		width:  width,
		min:    math.MaxInt,
		max:    math.MinInt,
	}
}

func (m *Matrix) AddValue(value, i, j int) {
	m.min = min(m.min, value)
	m.max = max(m.max, value)
	m.mtrx[i][j] = value
}

func (m Matrix) GetMaxNum() int {
	return m.max
}

func (m Matrix) GetMinNum() int {
	return m.min
}

func (m Matrix) Height() int {
	return m.height
}

func (m Matrix) Width() int {
	return m.width
}

func (m Matrix) CoordsToNum(coords cpt.Coords) int {
	return m.mtrx[coords.I][coords.J]
}
