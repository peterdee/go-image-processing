package processing

import "image/color"

func FlipVertical(grid [][]color.Color) [][]color.Color {
	for x := 0; x < len(grid); x += 1 {
		col := grid[x]
		for y := 0; y < len(col)/2; y += 1 {
			z := len(col) - y - 1
			col[y], col[z] = col[z], col[y]
		}
		grid[x] = col
	}
	return grid
}
