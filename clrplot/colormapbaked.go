package clrplot

import (
	"errors"
	"fmt"
	"sort"
)

type ColorMapVal struct {
	C     Color
	Ratio float64
}

type ColorMapInterface interface {
	NumToColor(int) Color
}

type ColorMap struct {
	vals    []ColorMapVal
}

type ColorMapValSlc []ColorMapVal

func (c ColorMapValSlc) Len() int {
	return len(c)
}

func (c ColorMapValSlc) Less(i, j int) bool {
	return c[i].Ratio < c[j].Ratio
}

func (cms ColorMapValSlc) Swap(i, j int) {
	cms[i], cms[j] = cms[j], cms[i]
}

func bisect(slc ColorMapValSlc, l, r int, fn func(ColorMapValSlc, int) bool) int {
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

func (cm ColorMap) NumToColor(num int) Color {
	n := len(cm.vals)
	idx := bisect(
		cm.vals,
		0,
		n,
		func(c ColorMapValSlc, i int) bool { return num < c[i].num },
	) - 1
	if idx == -1 {
		return cm.vals[0].C
	}
	if idx == n-1 {
		return cm.vals[n-1].C
	}

	val1 := cm.vals[idx]
	val2 := cm.vals[idx+1]
	x, y := val1.num, val2.num
	col1, col2 := val1.C, val2.C
	z := y - x
	return col1.portion(z+num, z).plus(col2.portion(num, z))
}

type ColorMapBakedVal struct {
	c     Color
	num   int
}

type ColorMapBaked struct {
	vals    []ColorMapBakedVal
}

func NewColorMap(
	val0 ColorMapVal,
	val1 ColorMapVal,
	vals ...ColorMapVal,
) (ColorMap, error) {
	valSlc := append(ColorMapValSlc{val0, val1}, vals...)
	sort.Sort(valSlc)
	for i := 1; i < len(valSlc); i++ {
		if valSlc[i-1].Ratio == valSlc[i].Ratio {
			return ColorMap{}, errors.New(fmt.Sprintf("identical .Ratio == %v values", valSlc[i].Ratio))
		}
	}
	return ColorMap{vals: valSlc}, nil
}

func (cm ColorMap) Init(minNum, maxNum int) (ColorMapBaked, error) {
	if minNum == maxNum {
		return ColorMapBaked{}, errors.New("identical minNum and maxNum param values")
	}
	if minNum > maxNum {
		return ColorMapBaked{}, errors.New("param value minNum > maxNum value")
	}

	delta := float64(maxNum - minNum)
	for i := 0; i < len(cm.vals); i++ {
		cm.vals[i].num = minNum + int(delta*cm.vals[i].Ratio)
	}

	return nil
}
