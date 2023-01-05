package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func checkSolarizationThreshold(partial uint8, threshold uint) uint8 {
	if partial <= uint8(threshold) {
		return 255 - partial
	}
	return partial
}

func Solarize(source [][]color.Color, threshold uint) [][]color.Color {
	threshold = utilities.MaxMin(threshold, 255, 0)
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			r, g, b, alpha := utilities.RGBA(source[x][y])
			cR := checkSolarizationThreshold(r, threshold)
			cG := checkSolarizationThreshold(g, threshold)
			cB := checkSolarizationThreshold(b, threshold)
			destination[x][y] = color.RGBA{cR, cG, cB, alpha}
		}
	}
	return destination
}
