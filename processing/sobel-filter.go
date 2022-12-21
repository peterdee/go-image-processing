package processing

import (
	"go-image-processing/utilities"
	"image/color"
	"math"
)

var SOBEL_HORIZONTAL = [3][3]int{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}

var SOBEL_VERTICAL = [3][3]int{
	{1, 2, 1},
	{0, 0, 0},
	{-1, -2, -1},
}

func SobelFilter(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])
	for x := 3; x < gridLen-3; x += 1 {
		for y := 3; y < colLen-3; y += 1 {
			gradientX := 0
			gradientY := 0

			// iStart, iEnd := utilities.GetPoints(x, 1, gridLen)
			// jStart, jEnd := utilities.GetPoints(y, 1, colLen)

			// if iEnd >= gridLen {
			// 	iEnd = gridLen - 1
			// }
			// if jEnd >= colLen {
			// 	jEnd = colLen - 1
			// }

			// fmt.Println(iStart, iEnd, jStart, jEnd)
			for i := 0; i < 3; i += 1 {
				for j := 0; j < 3; j += 1 {
					grayColor, _ := utilities.Gray(grid[x+i][y+j])
					gradientX += int(grayColor) * SOBEL_HORIZONTAL[i][j]
					gradientY += int(grayColor) * SOBEL_VERTICAL[i][j]
				}
			}

			colorCode := 255 - uint8(int(math.Sqrt(
				float64((gradientX*gradientX)+(gradientY*gradientY)),
			)))
			pixelColor := color.RGBA{colorCode, colorCode, colorCode, 255}
			grid[x][y] = pixelColor
		}
	}
	return grid
}
