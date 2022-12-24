package processing

import (
	"image/color"
	"math"

	"go-image-processing/utilities"
)

var LAPLASIAN_KERNEL = [3][3]int{
	{-1, -1, -1},
	{-1, 8, -1},
	{-1, -1, -1},
}

func LaplasianFilter(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])
	for x := 0; x < gridLen; x += 1 {
		for y := 0; y < colLen; y += 1 {
			gradientX := 0
			for i := 0; i < 3; i += 1 {
				for j := 0; j < 3; j += 1 {
					k := utilities.GetGradientPoint(x, i, gridLen)
					l := utilities.GetGradientPoint(y, j, colLen)
					grayColor, _ := utilities.Gray(grid[x+k][y+l])
					gradientX += int(grayColor) * LAPLASIAN_KERNEL[i][j]
				}
			}
			colorCode := 255 - uint8(int(math.Sqrt(
				float64((gradientX * gradientX)),
			)))
			grid[x][y] = color.RGBA{colorCode, colorCode, colorCode, 255}
		}
	}
	return grid
}
