package processing

import (
	"image/color"
	"math"

	"go-image-processing/utilities"
)

var EMBOSS_HORIZONTAL = [3][3]int{
	{0, 0, 0},
	{1, 0, -1},
	{0, 0, 0},
}

var EMBOSS_VERTICAL = [3][3]int{
	{0, 1, 0},
	{0, 0, 0},
	{0, -1, 0},
}

func EmbossFilter(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])
	for x := 0; x < gridLen; x += 1 {
		for y := 0; y < colLen; y += 1 {
			gradientX := 0
			gradientY := 0
			for i := 0; i < 3; i += 1 {
				for j := 0; j < 3; j += 1 {
					k := utilities.GetGradientPoint(x, i, gridLen)
					l := utilities.GetGradientPoint(y, j, colLen)
					grayColor, _ := utilities.Gray(grid[x+k][y+l])
					gradientX += int(grayColor) * EMBOSS_HORIZONTAL[i][j]
					gradientY += int(grayColor) * EMBOSS_VERTICAL[i][j]
				}
			}
			colorCode := 255 - uint8(int(math.Sqrt(
				float64((gradientX*gradientX)+(gradientY*gradientY)),
			)))
			grid[x][y] = color.RGBA{colorCode, colorCode, colorCode, 255}
		}
	}
	return grid
}
