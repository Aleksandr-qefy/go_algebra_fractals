package usertypes

import cpt "github.com/Aleksandr-qefy/clrplot"

type Config struct {
	MaxIter    int
	ColorMap   cpt.ColorMap
	GoroutineN int
	ImgHeight  int
	ImgWidth   int
	FileName   string
}
