package processing

import (
	"fmt"
	"image/color"

	"go-image-processing/utilities"
)

// var sharpenKernel = [3][3]float64{
// 	{-0.0023, -0.0432, -0.0023},
// 	{-0.0432, 1.182, -0.0432},
// 	{-0.0023, -0.0432, -0.0023},
// }

var sharpenKernel = [3][3]int{
	{-1, -1, -1},
	{-1, 17, -1},
	{-1, -1, -1},
}

func SharpenFilter(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			sumR := 0
			sumG := 0
			sumB := 0
			for i := 0; i < 3; i += 1 {
				for j := 0; j < 3; j += 1 {
					k := utilities.GetGradientPoint(x, i, width)
					l := utilities.GetGradientPoint(y, j, height)
					r, g, b, _ := utilities.RGBA(source[x+k][y+l])
					sumR += int(r) * sharpenKernel[i][j]
					sumG += int(g) * sharpenKernel[i][j]
					sumB += int(b) * sharpenKernel[i][j]
				}
			}
			R := uint8(sumR / 9)
			G := uint8(sumG / 9)
			B := uint8(sumB / 9)
			fmt.Println(sumR, R, sumG, G, sumB, B)
			destination[x][y] = color.RGBA{R, G, B, 255}
		}
	}
	return destination
}
