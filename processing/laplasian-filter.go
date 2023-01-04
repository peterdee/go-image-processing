package processing

import (
	"image/color"
	"math"

	"go-image-processing/utilities"
)

var laplasianKernel = [3][3]int{
	{0, -1, 0},
	{-1, 4, -1},
	{0, -1, 0},
}

func LaplasianFilter(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			gradientX := 0
			for i := 0; i < 3; i += 1 {
				for j := 0; j < 3; j += 1 {
					k := utilities.GetGradientPoint(x, i, width)
					l := utilities.GetGradientPoint(y, j, height)
					grayColor, _ := utilities.Gray(source[x+k][y+l])
					gradientX += int(grayColor) * laplasianKernel[i][j]
				}
			}
			colorCode := 255 - uint8(int(math.Sqrt(
				float64((gradientX * gradientX)),
			)))
			destination[x][y] = color.RGBA{colorCode, colorCode, colorCode, 255}
		}
	}
	return destination
}
