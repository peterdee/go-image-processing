package processing

import (
	"image/color"

	"go-image-processing/constants"
)

func getPoints(current, amount, total int) (int, int) {
	start, end := 0, total
	if current >= amount {
		start = current - amount
	}
	if current < total-amount {
		end = current + amount
	}
	return start, end
}

func BoxBlur(grid [][]color.Color, amount uint) [][]color.Color {
	if amount == 0 {
		amount = constants.DEFAULT_BLUR_AMOUNT
	}
	if int(amount) > (len(grid) / 2) {
		amount = uint(len(grid) / 2)
	}
	amountInt := int(amount)
	var denominator uint
	for x := 0; x < len(grid); x += 1 {
		col := grid[x]
		for y := 0; y < len(col); y += 1 {
			var tR, tG, tB uint
			_, _, _, A := grid[x][y].RGBA()
			denominator = 0

			iStart, iEnd := getPoints(x, amountInt, len(grid))
			jStart, jEnd := getPoints(y, amountInt, len(col))

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
