package processing

import (
	"fmt"
	"image/color"
	"math"

	"go-image-processing/constants"
)

func BoxBlur(grid [][]color.Color, amount uint) [][]color.Color {
	if amount == 0 {
		amount = constants.DEFAULT_BLUR_AMOUNT
	}
	for x := int(amount); x < len(grid)-int(amount); x += 1 {
		col := grid[x]
		for y := int(amount); y < len(col)-int(amount); y += 1 {
			var tR, tG, tB float64
			var pA uint32
			for i := int(amount) * (-1); i < int(amount); i += 1 {
				for j := int(amount) * (-1); j < int(amount); j += 1 {
					R, G, B, A := grid[x+i][y+j].RGBA()
					tR += float64(uint8(R))
					tG += float64(uint8(G))
					tB += float64(uint8(B))
					pA = A
				}
			}

			blurredR := math.Round(float64(uint8(tR)) / float64(amount*amount))
			blurredG := math.Round(float64(uint8(tG)) / float64(amount*amount))
			blurredB := math.Round(float64(uint8(tB)) / float64(amount*amount))
			fmt.Println(uint8(blurredR), uint8(blurredG), uint8(blurredB), blurredR)
			col[y] = color.RGBA{uint8(blurredR), uint8(blurredG), uint8(blurredB), uint8(pA)}
		}
		grid[x] = col
	}
	return grid
}
