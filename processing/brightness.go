package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func calculateNewBrightness(partial uint8, amount int) uint8 {
	newPartial := int(partial) + amount
	newPartial = utilities.MaxMin(newPartial, 255, 0)
	return uint8(newPartial)
}

func Brightness(source [][]color.Color, amount int) [][]color.Color {
	amount = utilities.MaxMin(amount, 255, -255)
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			r, g, b, alpha := utilities.RGBA(source[x][y])
			cR := calculateNewBrightness(r, amount)
			cG := calculateNewBrightness(g, amount)
			cB := calculateNewBrightness(b, amount)
			destination[x][y] = color.RGBA{cR, cG, cB, alpha}
		}
	}
	return destination
}
