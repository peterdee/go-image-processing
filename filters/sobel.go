package filters

import (
	"math"
	"time"

	"go-image-processing/utilities"
)

var sobelHorizontal = [3][3]int{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}

var sobelVertical = [3][3]int{
	{1, 2, 1},
	{0, 0, 0},
	{-1, -2, -1},
}

func Sobel(path string) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := getCoordinates(i/4, width)
		gradientX, gradientY := 0, 0
		for m := 0; m < 3; m += 1 {
			for n := 0; n < 3; n += 1 {
				k := getGradientPoint(x, m, width)
				l := getGradientPoint(y, n, height)
				px := getPixel(x+k, y+l, width)
				average := (int(img.Pix[px]) + int(img.Pix[px+1]) + int(img.Pix[px+2])) / 3
				gradientX += average * sobelHorizontal[m][n]
				gradientY += average * sobelVertical[m][n]
			}
		}
		channel := uint8(
			255 - utilities.MaxMin(
				math.Sqrt(float64(gradientX*gradientX+gradientY*gradientY)),
				255,
				0,
			),
		)
		img.Pix[i], img.Pix[i+1], img.Pix[i+2] = channel, channel, channel
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
