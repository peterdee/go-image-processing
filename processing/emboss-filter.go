package processing

import (
	"image/color"
	"math"
)

var horizontal = [3][3]int{
	{0, 2, 0},
	{0, 0, 0},
	{0, -2, 0},
}
var vertical = [3][3]int{
	{0, 0, 0},
	{2, 0, -2},
	{0, 0, 0},
}

func EmbossFilter(grid [][]color.Color) [][]color.Color {
	gray := Grayscale(grid)

	gridLen := len(gray)
	colLen := len(gray[0])
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
					r, _, _, _ := gray[x+i][j+y].RGBA()
					gradientX += int(uint8(r)) * vertical[i][j]
					gradientY += int(uint8(r)) * horizontal[i][j]
				}
			}

			colorCode := 255 - uint8(int(math.Sqrt(
				float64((gradientX*gradientX)+(gradientY*gradientY)),
			)))
			pixelColor := color.RGBA{colorCode, colorCode, colorCode, 255}
			gray[x][y] = pixelColor
		}
	}
	return gray
}
