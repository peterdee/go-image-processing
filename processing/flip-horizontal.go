package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func FlipHorizontal(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	correction := 0
	if width%2 != 0 {
		correction = 1
	}
	for x := 0; x < width/2+correction; x += 1 {
		for y := 0; y < height; y += 1 {
			z := width - x - 1
			destination[x][y], destination[z][y] = source[z][y], source[x][y]
		}
	}
	return destination
}
