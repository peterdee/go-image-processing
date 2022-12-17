package processing

import "image/color"

func FlipHorizontal(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])
	for x := 0; x < gridLen/2; x += 1 {
		for y := 0; y < colLen; y += 1 {
			z := gridLen - x - 1
			grid[x][y], grid[z][y] = grid[z][y], grid[x][y]
		}
	}
	return grid
}
