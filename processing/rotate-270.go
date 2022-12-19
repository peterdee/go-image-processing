package processing

import "image/color"

func Rotate270(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])

	rotatedGrid := make([][]color.Color, colLen)
	for ri := range rotatedGrid {
		rotatedGrid[ri] = make([]color.Color, gridLen)
	}

	for x := 0; x < gridLen; x += 1 {
		for y := 0; y < colLen; y += 1 {
			rotatedGrid[y][gridLen-x-1] = grid[x][y]
		}
	}
	return rotatedGrid
}
