package progress

import (
	"image/color"
	"math"

	"go-image-processing/utilities"
)

func distance(x, y, i, j int) float64 {
	return math.Sqrt(math.Pow(float64(x-i), 2) + math.Pow(float64(y-j), 2))
}

func gaussian(x int, sigma float64) float64 {
	return (1.0 / (2 * math.Pi * (math.Pow(sigma, 2)))) * math.Exp(-(math.Pow(float64(x), 2))/(2*math.Pow(sigma, 2)))
}

func Bilateral(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)

	// TODO: these should be arguments
	radius := 8
	diameter := radius * 2
	sigmaI := 24.0
	sigmaS := 32.0

	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {

			b := 20
			if x < b || x > width-b || y < b || y > height-b {
				destination[x][y] = source[x][y]
				continue
			}

			filteredR := 0.0
			filteredG := 0.0
			filteredB := 0.0
			WpR := 0.0
			WpG := 0.0
			WpB := 0.0

			for i := 0; i < diameter; i += 1 {
				for j := 0; j < diameter; j += 1 {
					neighbour_x := x - (radius - i)
					neighbour_y := y - (radius - j)

					// pi, _ := utilities.Gray(source[x][y])
					// px, _ := utilities.Gray(source[neighbour_x][neighbour_y])

					pr, pg, pb, _ := utilities.RGBA(source[x][y])
					nr, ng, nb, _ := utilities.RGBA(source[neighbour_x][neighbour_y])

					giR := gaussian(int(nr-pr), sigmaI)
					giG := gaussian(int(ng-pg), sigmaI)
					giB := gaussian(int(nb-pb), sigmaI)
					gs := gaussian(int(distance(neighbour_x, neighbour_y, x, y)), sigmaS)
					wR := giR * gs
					wG := giG * gs
					wB := giB * gs
					filteredR += float64(nr) * wR
					filteredG += float64(ng) * wG
					filteredB += float64(nb) * wB
					WpR += wR
					WpG += wG
					WpB += wB
				}
			}
			filteredR /= WpR
			filteredG /= WpG
			filteredB /= WpB
			R := uint8(utilities.MaxMin(filteredR, 255, 0))
			G := uint8(utilities.MaxMin(filteredG, 255, 0))
			B := uint8(utilities.MaxMin(filteredB, 255, 0))
			destination[x][y] = color.RGBA{R, G, B, 255}
		}
	}
	return destination
}
