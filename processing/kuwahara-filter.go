package processing

import (
	"fmt"
	"image/color"
)

type Point struct {
	x, y int
}

type GridPartial [3][3]Point

func createGridPartials(x, y int) [4]GridPartial {
	tl := [3][3]Point{
		{{x - 2, y - 2}, {x - 1, y - 2}, {x, y - 2}},
		{{x - 2, y - 1}, {x - 1, y - 1}, {x, y - 1}},
		{{x - 2, y}, {x - 1, y}, {x, y}},
	}
	tr := [3][3]Point{
		{{x, y - 2}, {x + 1, y - 2}, {x + 2, y - 2}},
		{{x, y - 1}, {x + 1, y - 1}, {x + 2, y - 1}},
		{{x, y}, {x + 1, y}, {x + 2, y}},
	}
	bl := [3][3]Point{
		{{x - 2, y}, {x - 1, y}, {x, y}},
		{{x - 2, y - 1}, {x - 1, y - 1}, {x, y - 1}},
		{{x - 2, y - 2}, {x - 1, y - 2}, {x, y - 2}},
	}
	br := [3][3]Point{
		{{x, y}, {x + 1, y}, {x + 2, y}},
		{{x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1}},
		{{x, y + 2}, {x + 1, y + 2}, {x + 2, y + 2}},
	}
	return [4]GridPartial{tl, tr, bl, br}
}

func KuwaharaFilter(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])
	for x := 3; x < gridLen-3; x += 1 {
		for y := 3; y < colLen-3; y += 1 {
			// gradientX := 0
			// gradientY := 0

			partials := createGridPartials(x, y)
			for _, partial := range partials {
				fmt.Println(partial)
			}
		}
	}
	return grid
}
