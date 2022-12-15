package processing

import (
	"image/color"
	"math"
)

func GrayscaleLuminocity(grid [][]color.Color) [][]color.Color {
	for x := 0; x < len(grid); x += 1 {
		col := grid[x]
		for y := 0; y < len(col); y += 1 {
			R, G, B, A := grid[x][y].RGBA()
			gray := math.Round(
				(float64(uint8(R))*0.21 + float64(uint8(G))*0.72 + float64(uint8(B))*0.07),
			)
			col[y] = color.RGBA{uint8(gray), uint8(gray), uint8(gray), uint8(A)}
		}
		grid[x] = col
	}
	return grid
}
