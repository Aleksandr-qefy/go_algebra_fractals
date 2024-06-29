package clrplot

import (
	"errors"
)

func (cm ColorMap) Bake(minNum, maxNum int) (ColorMapBaked, error) {
	if minNum == maxNum {
		return ColorMapBaked{}, errors.New("identical minNum and maxNum param values")
	}
	if minNum > maxNum {
		return ColorMapBaked{}, errors.New("param value minNum > maxNum value")
	}

	cmb := ColorMapBaked{
		vals:      make([]ColorMapBakedVal, len(cm.vals)),
		colorsMem: make([]Color, maxNum-minNum+1),
		min:       minNum,
		max:       maxNum,
	}
	delta := float64(maxNum - minNum)
	for i := 0; i < len(cm.vals); i++ {
		cmb.vals[i].num = minNum + int(delta*cm.vals[i].Ratio)
		cmb.vals[i].c = cm.vals[i].C
	}

	for i := 0; i < len(cmb.colorsMem); i++ {
		cmb.colorsMem[i] = cmb.numToColor(minNum + i)
	}

	return cmb, nil
}

type ColorMapBakedVal struct {
	c   Color
	num int
}

type ColorMapBakedValSlc []ColorMapBakedVal

type ColorMapBaked struct {
	vals      []ColorMapBakedVal
	colorsMem []Color
	min       int
	max       int
}

func bisect(slc ColorMapBakedValSlc, l, r int, fn func(ColorMapBakedValSlc, int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if fn(slc, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func (cmb ColorMapBaked) numToColor(num int) Color {
	n := len(cmb.vals)
	idx := bisect(
		cmb.vals,
		0,
		n,
		func(c ColorMapBakedValSlc, i int) bool { return num < c[i].num },
	) - 1
	if idx == -1 {
		return cmb.vals[0].c
	}
	if idx == n-1 {
		return cmb.vals[n-1].c
	}

	val1 := cmb.vals[idx]
	val2 := cmb.vals[idx+1]
	x, y := val1.num, val2.num
	col1, col2 := val1.c, val2.c
	z := y - x
	return col1.portion(z+num, z).plus(col2.portion(num, z))
}

func (cmb ColorMapBaked) NumToColor(num int) Color {
	if num < cmb.min {
		return cmb.colorsMem[0]
	}
	if num > cmb.max {
		return cmb.colorsMem[len(cmb.colorsMem)-1]
	}
	return cmb.colorsMem[num-cmb.min]
}
