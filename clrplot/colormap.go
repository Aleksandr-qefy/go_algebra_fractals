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

type ColorMap struct {
	vals []ColorMapVal
}

type ColorMapValSlc []ColorMapVal

func (c ColorMapValSlc) Len() int {
	return len(c)
}

func (c ColorMapValSlc) Less(i, j int) bool {
	return c[i].Ratio < c[j].Ratio
}

func (c ColorMapValSlc) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func NewColorMap(
	val0 ColorMapVal,
	val1 ColorMapVal,
	vals ...ColorMapVal,
) (ColorMap, error) {
	valSlc := append(ColorMapValSlc{val0, val1}, vals...)
	for i := 0; i < len(valSlc); i++ {
		if valSlc[i].Ratio < 0. {
			return ColorMap{}, errors.New(fmt.Sprintf(".Ratio %v value < 0", valSlc[i].Ratio))
		}
		if valSlc[i].Ratio > 1. {
			return ColorMap{}, errors.New(fmt.Sprintf(".Ratio %v value > 1", valSlc[i].Ratio))
		}
	}
	sort.Sort(valSlc)
	for i := 1; i < len(valSlc); i++ {
		if valSlc[i-1].Ratio == valSlc[i].Ratio {
			return ColorMap{}, errors.New(fmt.Sprintf("identical .Ratio == %v values", valSlc[i].Ratio))
		}
	}
	return ColorMap{vals: valSlc}, nil
}
