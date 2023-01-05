package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func calculateNewContrast(partial uint8, factor float64) uint8 {
	newPartial := factor*(float64(partial)-128) + 128
	newPartial = utilities.MaxMin(newPartial, 255, 0)
	return uint8(newPartial)
}

func Contrast(source [][]color.Color, amount int) [][]color.Color {
	amount = utilities.MaxMin(amount, 255, -255)
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	factor := float64(259*(amount+255)) / float64(255*(259-amount))
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			r, g, b, alpha := utilities.RGBA(source[x][y])
			cR := calculateNewContrast(r, factor)
			cG := calculateNewContrast(g, factor)
			cB := calculateNewContrast(b, factor)
			destination[x][y] = color.RGBA{cR, cG, cB, alpha}
		}
	}
	return destination
}
