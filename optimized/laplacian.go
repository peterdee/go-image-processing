package optimized

import (
	"math"
	"time"

	"go-image-processing/utilities"
)

var laplacianKernel = [3][3]int{
	{-1, -1, -1},
	{-1, 8, -1},
	{-1, -1, -1},
}

func Laplacian(path string) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	for i := 0; i < len(img.Pix); i += 4 {
		averageSum := 0
		x, y := getCoordinates(i/4, width)
		for m := 0; m < 3; m += 1 {
			for n := 0; n < 3; n += 1 {
				k := getGradientPoint(x, m, width)
				l := getGradientPoint(y, n, height)
				px := getPixel(x+k, y+l, width)
				average := (int(img.Pix[px]) + int(img.Pix[px+1]) + int(img.Pix[px+2])) / 3
				averageSum += average * laplacianKernel[m][n]
			}
		}
		img.Pix[i] = 255 - uint8(utilities.MaxMin(averageSum, 255, 0))
		img.Pix[i+1] = 255 - uint8(utilities.MaxMin(averageSum, 255, 0))
		img.Pix[i+2] = 255 - uint8(utilities.MaxMin(averageSum, 255, 0))
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
