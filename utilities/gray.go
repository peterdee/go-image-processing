package utilities

import (
	"image/color"
	"math"
)

func Gray(pixel color.Color) (gray uint8, alpha uint8) {
	R, G, B, A := pixel.RGBA()
	alpha = uint8(A)
	gray = uint8(
		math.Round(
			(float64(uint8(R)) + float64(uint8(G)) + float64(uint8(B))) / 3.0,
		),
	)
	return
}
