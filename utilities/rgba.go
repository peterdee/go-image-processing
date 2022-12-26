package utilities

import "image/color"

func RGBA(pixel color.Color) (r, g, b, alpha uint8) {
	R, G, B, A := pixel.RGBA()
	alpha = uint8(A)
	b = uint8(B)
	g = uint8(G)
	r = uint8(R)
	return
}
