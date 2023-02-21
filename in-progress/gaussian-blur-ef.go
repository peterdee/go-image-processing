package progress

import (
	"image/color"
	"sync"

	"go-image-processing/utilities"
)

func GaussianBlurEF(source [][]color.Color, radius uint) [][]color.Color {
	if radius == 0 {
		radius = 1
	}
	width, height := len(source), len(source[0])
	horizontal := utilities.CreateGrid(width, height)
	destination := utilities.CreateGrid(width, height)

	k := createKernel(4)
	j := len(k)/2 + 25

	var wg sync.WaitGroup

	processHorizontal := func(x, y int, px, res *[][]color.Color) {
		defer wg.Done()
		pixels := *px
		result := *res
		sumR := 0.0
		sumG := 0.0
		sumB := 0.0
		for m := 0; m < len(k); m += 1 {
			mx := x - (len(k)/2 - m)
			r, g, b, _ := utilities.RGBA(pixels[mx][y])
			sumR += float64(r) * float64(k[m])
			sumG += float64(g) * float64(k[m])
			sumB += float64(b) * float64(k[m])
		}

		result[x][y] = color.RGBA{
			uint8(utilities.MaxMin(sumR, 255, 0)),
			uint8(utilities.MaxMin(sumG, 255, 0)),
			uint8(utilities.MaxMin(sumB, 255, 0)),
			255,
		}
	}

	for x := 0; x < width-1; x += 2 {
		for y := 0; y < height-1; y += 2 {
			if x < j || x > width-j || y < j || y > height-j {
				horizontal[x][y] = source[x][y]
				continue
			}
			wg.Add(4)
			go processHorizontal(x, y, &source, &horizontal)
			go processHorizontal(x+1, y, &source, &horizontal)
			go processHorizontal(x, y+1, &source, &horizontal)
			go processHorizontal(x+1, y+1, &source, &horizontal)
		}
	}
	wg.Wait()

	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			if x < j || x > width-j || y < j || y > height-j {
				destination[x][y] = horizontal[x][y]
				continue
			}
			sumR := 0.0
			sumG := 0.0
			sumB := 0.0
			for n := 0; n < len(k); n += 1 {
				nx := y - (len(k)/2 - n)
				r, g, b, _ := utilities.RGBA(horizontal[x][nx])
				sumR += float64(r) * float64(k[n])
				sumG += float64(g) * float64(k[n])
				sumB += float64(b) * float64(k[n])
			}
			destination[x][y] = color.RGBA{
				uint8(utilities.MaxMin(sumR, 255, 0)),
				uint8(utilities.MaxMin(sumG, 255, 0)),
				uint8(utilities.MaxMin(sumB, 255, 0)),
				255,
			}
		}
	}
	return destination
}
