package processing

import (
	"image/color"

	"go-image-processing/constants"
)

func Blur(grid [][]color.Color, amount uint) [][]color.Color {
	if amount == 0 {
		amount = constants.DEFAULT_BLUR_AMOUNT
	}
	for x := 0; x < len(grid); x += 1 {
		col := grid[x]
		if x > int(amount) && x < (len(grid)-int(amount)) {
			for y := 0; y < len(col); y += 1 {
				if y > int(amount) && y < (len(col)-int(amount)) {
					// R, G, B, A := grid[x][y].RGBA()
					// for i := 0; i < int(amount); i += 1 {
					// 	for j := 0; j < int(amount); j += 1 {
					// 		p1 := grid[x - i][y - j]
					// 		p2 := grid[x + i][y + j]
					// 		gray := math.Round((float64(uint8(R)) + float64(uint8(G)) + float64(uint8(B))) / 3.0)
					// 		value := uint8(255)
					// 		if uint(gray) < amount {
					// 			value = 0
					// 		}
					// 	}
					// }
					// col[y] = color.RGBA{value, value, value, uint8(A)}
				}
			}
		}
	}
	return grid
}
