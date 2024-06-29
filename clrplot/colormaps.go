package clrplot

import (
	"sort"
)

type ColorMapVal struct {
	C   Color
	Num int
}

type ColorMapValSlc []ColorMapVal

func (cms ColorMapValSlc) Len() int {
	return len(cms)
}

func (cms ColorMapValSlc) Less(i, j int) bool {
	return cms[i].Num < cms[j].Num
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

func twoColorsToColor(col1, col2 ColorMapVal, num int) {

}

func (cm ColorMap) NumToColor(num int) Color {
	n := len(cm.vals)
	idx := bisect(
		cm.vals,
		0,
		n,
		func(cms ColorMapValSlc, i int) bool { return num < cms[i].Num },
	) - 1
	if idx == -1 {
		return cm.vals[0].C
	}
	if idx == n-1 {
		return cm.vals[n-1].C
	}

	val1 := cm.vals[idx]
	val2 := cm.vals[idx+1]
	x, y := val1.Num, val2.Num
	col1, col2 := val1.C, val2.C
	z := y - x
	return col1.portion(z+num, z).plus(col2.portion(num, z))
}

type ColorMap struct {
	vals []ColorMapVal
}

func NewColorMap(
	val0 ColorMapVal,
	val1 ColorMapVal,
	vals ...ColorMapVal,
) ColorMap {
	valSlc := append(ColorMapValSlc{val0, val1}, vals...)
	sort.Sort(valSlc)
	return ColorMap{vals: valSlc}
}
