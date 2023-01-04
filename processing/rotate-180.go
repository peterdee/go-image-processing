package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func Rotate180(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			destination[width-x-1][height-y-1] = source[x][y]
		}
	}
	return destination
}
