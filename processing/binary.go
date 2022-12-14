package processing

import (
	"image/color"
	"math"

	"go-image-processing/constants"
)

func Binary(grid [][]color.Color, threshold uint) [][]color.Color {
	if threshold > 255 {
		threshold = constants.DEFAULT_BINARY_THRESHOLD
	}
	for x := 0; x < len(grid); x += 1 {
		col := grid[x]
		for y := 0; y < len(col); y += 1 {
			R, G, B, A := grid[x][y].RGBA()
			gray := math.Round((float64(uint8(R)) + float64(uint8(G)) + float64(uint8(B))) / 3.0)
			value := uint8(255)
			if uint(gray) < threshold {
				value = 0
			}
			col[y] = color.RGBA{value, value, value, uint8(A)}
		}
		grid[x] = col
	}
	return grid
}
