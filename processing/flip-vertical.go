package processing

import "image/color"

func FlipVertical(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])
	for x := 0; x < gridLen; x += 1 {
		col := grid[x]
		for y := 0; y < colLen/2; y += 1 {
			z := colLen - y - 1
			col[y], col[z] = col[z], col[y]
		}
		grid[x] = col
	}
	return grid
}
