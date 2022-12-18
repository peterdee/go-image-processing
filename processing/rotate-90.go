package processing

import "image/color"

func Rotate90(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])
	rotatedGrid := make([][]color.Color, colLen)
	for x := 0; x < gridLen; x += 1 {
		rotatedGridCol := make([]color.Color, gridLen)
		for y := 0; y < colLen; y += 1 {
			rotatedGrid[y][colLen-x] = grid[x][y]
		}
	}
	return grid
}
