package processing

import (
	"fmt"
	"image/color"

	"go-image-processing/constants"
)

func getPoints(x, y, amount, gridLen, colLen int) (int, int, int, int) {
	iE, iS, jE, jS := 0, 0, 0, 0
	if x >= amount {
		iS = amount * (-1)
	} else if x < amount {
		iS = x * (-1)
	}
	if y >= amount {
		jS = amount * (-1)
	} else if y < amount {
		jS = y * (-1)
	}
	if x > gridLen-amount {
		iE = amount * (-1)
	} else if x < amount {
		iE = x * (-1)
	}
	if y > colLen-amount {
		jE = amount * (-1)
	} else if y < amount {
		jE = y * (-1)
	}
	return iE, iS, jE, jS
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

			iE, iS, jE, jS := getPoints(x, y, amountInt, len(grid), len(col))
			for i := iS; i <= iE; i += 1 {
				for j := jS; j <= jE; j += 1 {
					denominator += 1
					R, G, B, _ := grid[x+i][y+j].RGBA()
					tR += uint(uint8(R))
					tG += uint(uint8(G))
					tB += uint(uint8(B))
				}
			}
			fmt.Println(tR, tG, tB, denominator)

			bR := tR / denominator
			bG := tG / denominator
			bB := tB / denominator
			col[y] = color.RGBA{uint8(bR), uint8(bG), uint8(bB), uint8(A)}
		}
		grid[x] = col
	}
	return grid
}
