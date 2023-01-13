package processing

import (
	"image/color"
	"math"

	"go-image-processing/utilities"
)

var DEG float64 = math.Pi / 180

func HueRotate(source [][]color.Color, angle int) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)

	cos := math.Cos(float64(angle) * DEG)
	sin := math.Sin(float64(angle) * DEG)
	matrix := [3]float64{
		cos + (1-cos)/3,
		(1-cos)/3 - math.Sqrt(float64(1)/3)*sin,
		(1-cos)/3 + math.Sqrt(float64(1)/3)*sin,
	}

	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			r, g, b, alpha := utilities.RGBA(source[x][y])
			rr := float64(r)*matrix[0] + float64(g)*matrix[1] + float64(b)*matrix[2]
			rg := float64(r)*matrix[2] + float64(g)*matrix[0] + float64(b)*matrix[1]
			rb := float64(r)*matrix[1] + float64(g)*matrix[2] + float64(b)*matrix[0]
			destination[x][y] = color.RGBA{
				uint8(rr),
				uint8(rg),
				uint8(rb),
				alpha,
			}
		}
	}
	return destination
}
