package processing

import "image/color"

func Rotate180(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])

	rotatedGrid := make([][]color.Color, gridLen)
	for ri := range rotatedGrid {
		rotatedGrid[ri] = make([]color.Color, colLen)
	}

	for x := 0; x < gridLen; x += 1 {
		for y := 0; y < colLen; y += 1 {
			rotatedGrid[gridLen-x-1][colLen-y-1] = grid[x][y]
		}
	}
	return rotatedGrid
}
