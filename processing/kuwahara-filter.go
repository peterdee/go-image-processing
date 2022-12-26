package processing

import (
	"image/color"
	"math"

	"go-image-processing/utilities"
)

type Point struct {
	x, y int
}

type GridPartial3x3 [3][3]Point

type GridPartial5x5 [5][5]Point

func createGridPartials5x5(x, y int) [4]GridPartial5x5 {
	tl := GridPartial5x5{
		{{x - 4, y - 4}, {x - 3, y - 4}, {x - 2, y - 4}, {x - 1, y - 4}, {x, y - 4}},
		{{x - 4, y - 3}, {x - 3, y - 3}, {x - 2, y - 3}, {x - 1, y - 3}, {x, y - 3}},
		{{x - 4, y - 2}, {x - 3, y - 2}, {x - 2, y - 2}, {x - 1, y - 2}, {x, y - 2}},
		{{x - 4, y - 1}, {x - 3, y - 1}, {x - 2, y - 1}, {x - 1, y - 1}, {x, y - 1}},
		{{x - 4, y}, {x - 3, y}, {x - 2, y}, {x - 1, y}, {x, y}},
	}
	tr := GridPartial5x5{
		{{x, y - 4}, {x + 1, y - 4}, {x + 2, y - 4}, {x + 3, y - 4}, {x + 4, y - 4}},
		{{x, y - 3}, {x + 1, y - 3}, {x + 2, y - 3}, {x + 3, y - 3}, {x + 4, y - 3}},
		{{x, y - 2}, {x + 1, y - 2}, {x + 2, y - 2}, {x + 3, y - 2}, {x + 4, y - 2}},
		{{x, y - 1}, {x + 1, y - 1}, {x + 2, y - 1}, {x + 3, y - 1}, {x + 4, y - 1}},
		{{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y}, {x + 4, y}},
	}
	bl := GridPartial5x5{
		{{x - 4, y}, {x - 3, y}, {x - 2, y}, {x - 1, y}, {x, y}},
		{{x - 4, y - 1}, {x - 3, y - 1}, {x - 2, y - 1}, {x - 1, y - 1}, {x, y - 1}},
		{{x - 4, y - 2}, {x - 3, y - 2}, {x - 2, y - 2}, {x - 1, y - 2}, {x, y - 2}},
		{{x - 4, y - 3}, {x - 3, y - 3}, {x - 2, y - 3}, {x - 1, y - 3}, {x, y - 3}},
		{{x - 4, y - 4}, {x - 3, y - 4}, {x - 2, y - 4}, {x - 1, y - 4}, {x, y - 4}},
	}
	br := GridPartial5x5{
		{{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y}, {x + 4, y}},
		{{x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1}, {x + 3, y + 1}, {x + 4, y + 1}},
		{{x, y + 2}, {x + 1, y + 2}, {x + 2, y + 2}, {x + 3, y + 2}, {x + 4, y + 2}},
		{{x, y + 3}, {x + 1, y + 3}, {x + 2, y + 3}, {x + 3, y + 3}, {x + 4, y + 3}},
		{{x, y + 4}, {x + 1, y + 4}, {x + 2, y + 4}, {x + 3, y + 4}, {x + 4, y + 4}},
	}
	return [4]GridPartial5x5{tl, tr, bl, br}
}

func createGridPartials(x, y int) [4]GridPartial3x3 {
	tl := GridPartial3x3{
		{{x - 2, y - 2}, {x - 1, y - 2}, {x, y - 2}},
		{{x - 2, y - 1}, {x - 1, y - 1}, {x, y - 1}},
		{{x - 2, y}, {x - 1, y}, {x, y}},
	}
	tr := GridPartial3x3{
		{{x, y - 2}, {x + 1, y - 2}, {x + 2, y - 2}},
		{{x, y - 1}, {x + 1, y - 1}, {x + 2, y - 1}},
		{{x, y}, {x + 1, y}, {x + 2, y}},
	}
	bl := GridPartial3x3{
		{{x - 2, y}, {x - 1, y}, {x, y}},
		{{x - 2, y - 1}, {x - 1, y - 1}, {x, y - 1}},
		{{x - 2, y - 2}, {x - 1, y - 2}, {x, y - 2}},
	}
	br := GridPartial3x3{
		{{x, y}, {x + 1, y}, {x + 2, y}},
		{{x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1}},
		{{x, y + 2}, {x + 1, y + 2}, {x + 2, y + 2}},
	}
	return [4]GridPartial3x3{tl, tr, bl, br}
}

func KuwaharaFilter(grid [][]color.Color) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])
	for x := 5; x < gridLen-5; x += 1 {
		for y := 5; y < colLen-5; y += 1 {
			grayColor, alpha := utilities.Gray(grid[x][y])

			partialsSum := [4][3]int{}

			partials := createGridPartials5x5(x, y)
			for index, partial := range partials {
				for i := 0; i < 5; i += 1 {
					for j := 0; j < 5; j += 1 {
						gridPartialPoint := partial[i][j]
						r, g, b, _ := utilities.RGBA(grid[gridPartialPoint.x][gridPartialPoint.y])
						partialsSum[index][0] += int(r)
						partialsSum[index][1] += int(g)
						partialsSum[index][2] += int(b)
					}
				}
			}

			average := [3]int{}
			difference := 0
			for index, pix := range partialsSum {
				partialAverageR := math.Round(float64(pix[0]) / 25)
				partialAverageG := math.Round(float64(pix[1]) / 25)
				partialAverageB := math.Round(float64(pix[2]) / 25)
				partialAverage := int(
					math.Round((partialAverageB + partialAverageG + partialAverageR) / 3.0),
				)
				localDifference := int(math.Abs(float64(partialAverage - int(grayColor))))
				if index == 0 {
					average = [3]int{int(partialAverageR), int(partialAverageG), int(partialAverageB)}
					difference = localDifference
				}
				if localDifference <= difference {
					average = [3]int{int(partialAverageR), int(partialAverageG), int(partialAverageB)}
					difference = localDifference
				}
			}

			grid[x][y] = color.RGBA{uint8(average[0]), uint8(average[1]), uint8(average[2]), alpha}
		}
	}
	return grid
}
