package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func Sepia(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			r, g, b, alpha := utilities.RGBA(source[x][y])
			dR := 0.393*float64(r) + 0.769*float64(g) + 0.189*float64(b)
			if dR > 255 {
				dR = 255
			}
			dG := 0.349*float64(r) + 0.686*float64(g) + 0.168*float64(b)
			if dG > 255 {
				dG = 255
			}
			dB := 0.272*float64(r) + 0.534*float64(g) + 0.131*float64(b)
			if dB > 255 {
				dB = 255
			}
			destination[x][y] = color.RGBA{uint8(dR), uint8(dG), uint8(dB), alpha}
		}
	}
	return destination
}
