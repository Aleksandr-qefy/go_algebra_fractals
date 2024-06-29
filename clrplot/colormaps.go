package clrplot

func GenViridis() ColorMap {
	cm, _ := NewColorMap(
		ColorMapVal{C: Color{R: 41, G: 0, B: 58}, Ratio: 0.1},
		ColorMapVal{C: Color{R: 4, G: 128, B: 104}, Ratio: 0.5},
		ColorMapVal{C: Color{R: 217, G: 196, B: 10}, Ratio: 0.9},
	)
	return cm
}

func GenBlack() ColorMap {
	cm, _ := NewColorMap(
		ColorMapVal{C: Color{R: 0, G: 0, B: 0}, Ratio: 0},
		ColorMapVal{C: Color{R: 255, G: 255, B: 255}, Ratio: 1},
	)
	return cm
}

func GenPlasma() ColorMap {
	cm, _ := NewColorMap(
		ColorMapVal{C: Color{R: 41, G: 0, B: 58}, Ratio: 0.1},
		ColorMapVal{C: Color{R: 184, G: 61, B: 76}, Ratio: 0.5},
		ColorMapVal{C: Color{R: 217, G: 196, B: 10}, Ratio: 0.9},
	)
	return cm
}

func GenInferno() ColorMap {
	cm, _ := NewColorMap(
		ColorMapVal{C: Color{R: 0, G: 0, B: 0}, Ratio: 0},
		ColorMapVal{C: Color{R: 41, G: 0, B: 58}, Ratio: 0.2},
		ColorMapVal{C: Color{R: 184, G: 61, B: 76}, Ratio: 0.5},
		ColorMapVal{C: Color{R: 217, G: 196, B: 10}, Ratio: 0.8},
		ColorMapVal{C: Color{R: 255, G: 255, B: 255}, Ratio: 1},
	)
	return cm
}
