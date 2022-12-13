package processing

import "image/color"

func Invert(grid [][]color.Color) [][]color.Color {
	for x := 0; x < len(grid); x += 1 {
		col := grid[x]
		for y := 0; y < len(col); y += 1 {
			R, G, B, A := grid[x][y].RGBA()
			col[y] = color.RGBA{255 - uint8(R), 255 - uint8(G), 255 - uint8(B), uint8(A)}
		}
		grid[x] = col
	}
	return grid
}
