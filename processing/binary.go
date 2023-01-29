package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func Binary(source [][]color.Color, threshold uint8) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			gray, alpha := utilities.Gray(source[x][y])
			value := uint8(255)
			if gray < threshold {
				value = 0
			}
			destination[x][y] = color.RGBA{value, value, value, alpha}
		}
	}
	return destination
}
