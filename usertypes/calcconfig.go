package usertypes

import "clrplot/clrplot"

type Config struct {
	MaxIter    int
	ColorMap   clrplot.ColorMap
	GoroutineN int
	ImgHeight  int
	ImgWidth   int
	FileName   string
}
