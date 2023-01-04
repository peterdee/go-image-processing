package processing

import (
	"image/color"
	"math"

	"go-image-processing/constants"
	"go-image-processing/utilities"
)

func BoxBlur(source [][]color.Color, amount uint) [][]color.Color {
	if amount == 0 {
		amount = constants.DEFAULT_BLUR_AMOUNT
	}
	width, height := len(source), len(source[0])

	// auto-adjust amount based on the source size
	min := math.Min(float64(width), float64(height))
	if amount > (uint(min) / 2) {
		amount = uint(min / 2)
	}

	amountInt := int(amount)
	var denominator uint
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			var tR, tG, tB uint
			_, _, _, A := source[x][y].RGBA()
			denominator = 0

			iStart, iEnd := utilities.GetPoints(x, amountInt, width)
			jStart, jEnd := utilities.GetPoints(y, amountInt, height)

			for i := iStart; i < iEnd; i += 1 {
				for j := jStart; j < jEnd; j += 1 {
					denominator += 1
					R, G, B, _ := source[i][j].RGBA()
					tR += uint(uint8(R))
					tG += uint(uint8(G))
					tB += uint(uint8(B))
				}
			}

			bR := tR / denominator
			bG := tG / denominator
			bB := tB / denominator
			destination[x][y] = color.RGBA{uint8(bR), uint8(bG), uint8(bB), uint8(A)}
		}
	}
	return destination
}
