package processing

import (
	"image/color"
	"math"

	"go-image-processing/utilities"
)

func GammaCorrection(source [][]color.Color, amount float64) [][]color.Color {
	if amount > 3.99 {
		amount = 3.99
	}
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	power := 1 / amount
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			r, g, b, alpha := utilities.RGBA(source[x][y])
			cR := uint8(255 * math.Pow(float64(r)/255, power))
			cG := uint8(255 * math.Pow(float64(g)/255, power))
			cB := uint8(255 * math.Pow(float64(b)/255, power))
			destination[x][y] = color.RGBA{cR, cG, cB, alpha}
		}
	}
	return destination
}
