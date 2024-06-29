# Algebra Fractals

## Description
In this project I write code to calculate fractals like 
Mandelbrot and Burning Ship.


To speed up calculations image rows are calculated in goroutines.
So, I use sync.Mutex and sync.WaitGroup types variables.


Extra information about goroutines number and calculating time
is printed in terminal. 


To draw resulting calculations I use my
[clrplot module](https://github.com/Aleksandr-qefy/clrplot).

## How to run

```
cd ./go_algebra_fractals
go run .\main.go
```


You can use your fractal and config by initializing variables of 
`Fractal` and `Config` struct types and pass them to `DrawFractal` function,
as in main function right now:
```
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
```


## Examples

You can see resulting images in [`/images` dir](https://github.com/Aleksandr-qefy/go_algebra_fractals/tree/main/images).
