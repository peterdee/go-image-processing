package utilities

import "image/color"

func CreateGrid(width, height int) [][]color.Color {
	gridCopy := make([][]color.Color, width)
	for i := range gridCopy {
		gridCopy[i] = make([]color.Color, height)
	}
	return gridCopy
}
