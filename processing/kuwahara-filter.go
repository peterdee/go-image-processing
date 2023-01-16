package processing

import (
	"image/color"
	"math"

	"go-image-processing/utilities"
)

type Point struct {
	x, y int
}

type GridPartial5x5 [5][5]Point

type Partials struct {
	LeftBottom, LeftTop, RightBottom, RightTop [][]Point
}

func createPartials(x, y, radius int) Partials {
	var partials Partials
	partials.LeftBottom = make([][]Point, radius)
	partials.LeftTop = make([][]Point, radius)
	partials.RightBottom = make([][]Point, radius)
	partials.RightTop = make([][]Point, radius)
	for i := 0; i < radius; i += 1 {
		leftBottomRow := make([]Point, radius)
		leftTopRow := make([]Point, radius)
		rightBottomRow := make([]Point, radius)
		rightTopRow := make([]Point, radius)
		for j := 0; j < radius; j += 1 {
			leftBottomRow[j] = Point{
				x: x - i,
				y: y + j,
			}
			leftTopRow[j] = Point{
				x: x - i,
				y: y - j,
			}
			rightBottomRow[j] = Point{
				x: x + i,
				y: y + j,
			}
			rightTopRow[j] = Point{
				x: x + i,
				y: y - j,
			}
		}
		partials.LeftBottom[i] = leftBottomRow
		partials.LeftTop[i] = leftTopRow
		partials.RightBottom[i] = rightBottomRow
		partials.RightTop[i] = rightTopRow
	}
	return partials
}

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

func KuwaharaFilter(grid [][]color.Color, radius uint) [][]color.Color {
	gridLen := len(grid)
	colLen := len(grid[0])
	radiusInt := int(radius)
	for x := radiusInt; x < gridLen-radiusInt; x += 1 {
		for y := radiusInt; y < colLen-radiusInt; y += 1 {
			grayColor, alpha := utilities.Gray(grid[x][y])

			partialsSum := [4][3]int{}

			partials := createPartials(x, y, radiusInt)
			for p := 0; p < 4; p += 1 {
				for i := 0; i < radiusInt; i += 1 {
					for j := 0; j < radiusInt; j += 1 {
						partial := partials.LeftBottom
						if p == 1 {
							partial = partials.LeftTop
						}
						if p == 2 {
							partial = partials.RightBottom
						}
						if p == 3 {
							partial = partials.RightTop
						}
						gridPartialPoint := partial[i][j]
						r, g, b, _ := utilities.RGBA(grid[gridPartialPoint.x][gridPartialPoint.y])
						partialsSum[p][0] += int(r)
						partialsSum[p][1] += int(g)
						partialsSum[p][2] += int(b)
					}
				}
			}

			average := [3]int{}
			difference := 0
			for index, pix := range partialsSum {
				partialAverageR := math.Round(float64(pix[0]) / float64(radiusInt*radiusInt))
				partialAverageG := math.Round(float64(pix[1]) / float64(radiusInt*radiusInt))
				partialAverageB := math.Round(float64(pix[2]) / float64(radiusInt*radiusInt))
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
