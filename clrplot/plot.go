package clrplot

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Coords struct {
	I int
	J int
}

type Plotable interface {
	CoordsToNum(Coords) int
	GetMaxNum() int
	GetMinNum() int
	Height() int
	Width() int
}

type NumToColorInterface interface {
	NumToColor(int) Color
}

func Draw(mtrx Plotable, imgName string, cm ColorMap) error {
	m := mtrx.Height()
	n := mtrx.Width()

	cmb, err := cm.Bake(mtrx.GetMinNum(), mtrx.GetMaxNum())
	if err != nil {
		return errors.New(fmt.Sprintf("ColorMap baking faild: %v", err))
	}

	img := image.NewNRGBA(image.Rect(0, 0, n, m))

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num := mtrx.CoordsToNum(
				Coords{I: i, J: j},
			)
			col := cmb.NumToColor(num)

			img.Set(j, i, color.NRGBA{
				R: col.R,
				G: col.G,
				B: col.B,
				A: 255,
			})
		}
	}

	imgFile, err := os.Create(imgName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := imgFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := png.Encode(imgFile, img); err != nil {
		log.Fatal(err)
	}
	return nil
}
