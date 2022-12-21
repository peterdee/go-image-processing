package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func Grayscale(grid [][]color.Color) [][]color.Color {
	for x := 0; x < len(grid); x += 1 {
		col := grid[x]
		for y := 0; y < len(col); y += 1 {
			grayColor, alpla := utilities.Gray(grid[x][y])
			col[y] = color.RGBA{grayColor, grayColor, grayColor, alpla}
		}
		grid[x] = col
	}
	return grid
}
