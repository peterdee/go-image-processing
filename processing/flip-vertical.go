package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func FlipVertical(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height/2; y += 1 {
			z := height - y - 1
			destination[x][y], destination[x][z] = source[x][z], source[x][y]
		}
	}
	return destination
}
