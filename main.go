package main

import (
	cpt "clrplot/clrplot"
	ut "clrplot/usertypes"
	"fmt"
	"math/cmplx"
	"runtime"
	"sync"
	"time"
)

func DrawFractal(fractal ut.Fractal, config ut.Config) {
	numCPU := runtime.NumCPU()
	fmt.Printf("CPU number:       %v\n", numCPU)
	runtime.GOMAXPROCS(config.GoroutineN)
	fmt.Printf("Goroutine number: %v\n", config.GoroutineN)

	var mutex sync.Mutex
	var waitGroup sync.WaitGroup

	matrix := ut.NewMatrix(config.ImgHeight, config.ImgWidth)
	procPixel := func(x, y float64, i, j int) {
		c := complex(x, y)
		z := 0 + 0i
		k := 0
		for ; k < config.MaxIter; k++ {
			z = fractal.Function(z, c)
			if cmplx.Abs(z) > 4 {
				mutex.Lock()
				matrix.AddValue(k, i, j)
				mutex.Unlock()
				break
			}
		}
		if k == config.MaxIter {
			matrix.AddValue(config.MaxIter, i, j)
		}
	}

	deltaW := (fractal.MaxX - fractal.MinX) / (float64)(config.ImgWidth)
	procRow := func(y float64, i int) {
		waitGroup.Add(1)
		x := fractal.MinX
		for j := 0; j < config.ImgWidth; j++ {
			procPixel(x, y, i, j)
			x += deltaW
		}
		waitGroup.Done()
	}

	start := time.Now()

	deltaH := (fractal.MaxY - fractal.MinY) / (float64)(config.ImgHeight)
	y := fractal.MinY
	for i := 0; i < config.ImgHeight; i++ {
		go procRow(y, i)
		y += deltaH
	}

	waitGroup.Wait()
	fmt.Printf(
		"Calculating time: %v ms\n",
		time.Now().Sub(start).Milliseconds(),
	)

	err := cpt.Draw(matrix, config.FileName, config.ColorMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("Image saved")
}

func main() {
	DrawFractal(
		ut.Mandelbrot,
		ut.Config{
			FileName:   "./images/image.png",
			ImgHeight:  600 * 5,
			ImgWidth:   1000 * 5,
			MaxIter:    150,
			ColorMap:   cpt.GenViridis(),
			GoroutineN: runtime.NumCPU(),
		},
	)
}
