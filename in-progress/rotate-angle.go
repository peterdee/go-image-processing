package progress

import (
	"fmt"
	"image/color"
	"math"

	"go-image-processing/utilities"
)

func RotateAngle(source [][]color.Color, angle float64) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)

	cos := math.Cos(angle)
	sin := math.Sin(angle)

	for x := 0; x < width-10; x += 1 {
		for y := 0; y < height-10; y += 1 {
			dx := math.Abs(float64(x)*cos - float64(y)*sin)
			dy := math.Abs(float64(x)*sin + float64(y)*cos)
			// fmt.Println(dx, dy, x, y)
			rx := int(math.Round(dx))
			ry := int(math.Round(dy))
			if rx < width-10 && rx > 0 && ry < height-10 && ry > 0 {
				fmt.Println(x, y, rx, ry)
				destination[rx][ry] = source[x][y]
			}
		}
	}
	return destination
}
