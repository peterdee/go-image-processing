package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

type Color struct {
	R, G, B int
}

var COLORS = [8]Color{
	{255, 0, 0},
	{0, 255, 0},
	{0, 0, 255},
	{255, 255, 0},
	{255, 0, 255},
	{0, 255, 255},
	{255, 255, 255},
	{0, 0, 0},
}

func EightColors(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			r, g, b, alpha := utilities.RGBA(source[x][y])
			minDelta := 195076
			var selectedColor Color
			for i := range COLORS {
				indexColor := COLORS[i]
				rDifference := int(r) - indexColor.R
				gDifference := int(g) - indexColor.G
				bDifference := int(b) - indexColor.B
				delta := rDifference*rDifference + gDifference*gDifference + bDifference*bDifference
				if delta < minDelta {
					minDelta = delta
					selectedColor = indexColor
				}
			}
			destination[x][y] = color.RGBA{
				uint8(selectedColor.R),
				uint8(selectedColor.G),
				uint8(selectedColor.B),
				alpha,
			}
		}
	}
	return destination
}
