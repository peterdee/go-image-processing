package processing

import (
	"image/color"

	"go-image-processing/constants"
	"go-image-processing/utilities"
)

func BoxBlur(grid [][]color.Color, amount uint) [][]color.Color {
	if amount == 0 {
		amount = constants.DEFAULT_BLUR_AMOUNT
	}
	gridLen := len(grid)
	if int(amount) > (gridLen / 2) {
		amount = uint(gridLen / 2)
	}
	amountInt := int(amount)
	var denominator uint
	for x := 0; x < gridLen; x += 1 {
		col := grid[x]
		colLen := len(col)
		for y := 0; y < colLen; y += 1 {
			var tR, tG, tB uint
			_, _, _, A := grid[x][y].RGBA()
			denominator = 0

			iStart, iEnd := utilities.GetPoints(x, amountInt, gridLen)
			jStart, jEnd := utilities.GetPoints(y, amountInt, colLen)

			for i := iStart; i < iEnd; i += 1 {
				for j := jStart; j < jEnd; j += 1 {
					denominator += 1
					R, G, B, _ := grid[i][j].RGBA()
					tR += uint(uint8(R))
					tG += uint(uint8(G))
					tB += uint(uint8(B))
				}
			}

			bR := tR / denominator
			bG := tG / denominator
			bB := tB / denominator
			col[y] = color.RGBA{uint8(bR), uint8(bG), uint8(bB), uint8(A)}
		}
		grid[x] = col
	}
	return grid
}
