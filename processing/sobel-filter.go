package processing

import (
	"fmt"
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

func getPoints(axisValue, axisLength int) (int, int) {
	start := 0
	end := axisValue + 2
	if axisValue >= 1 {
		start = axisValue - 1
	}
	if end > axisLength {
		end = axisLength
	}
	return start, end
}

func SobelFilter(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])
	for x := 0; x < gridLen; x += 1 {
		for y := 0; y < colLen; y += 1 {
			gradientX := 0
			gradientY := 0

			iS, iE := getPoints(x, gridLen)
			jS, jE := getPoints(y, colLen)
			fmt.Println(iS, iE, jS, jE)
			mi, mj := 0, 0
			for i := iS; i < iE; i += 1 {
				mj = 0
				for j := jS; j < jE; j += 1 {
					grayColor, _ := utilities.Gray(grid[i][j])
					gradientX += int(grayColor) * SOBEL_HORIZONTAL[mi][mj]
					gradientY += int(grayColor) * SOBEL_VERTICAL[mi][mj]
					mj += 1
				}
				mi += 1
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
