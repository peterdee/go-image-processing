package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func Rotate270(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(height, width)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			destination[y][width-x-1] = source[x][y]
		}
	}
	return destination
}
