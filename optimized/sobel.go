package optimized

import (
	"fmt"
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

func getCoordinates(pixel, width int) (int, int) {
	return pixel % width, int(math.Floor(float64(pixel) / float64(width)))
}

func getGradientPoint(axisValue, shift, axisLength int) int {
	if (axisValue + shift) >= axisLength {
		return axisLength - axisValue - 1
	}
	return shift
}

func getPixel(x, y, width int) int {
	return ((y * width) + x) * 4
}

func Sobel(path string) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := getCoordinates(i/4, width)
		gradientX := 0
		gradientY := 0
		for m := 0; m < 3; m += 1 {
			for n := 0; n < 3; n += 1 {
				k := getGradientPoint(x, m, width)
				l := getGradientPoint(y, n, height)
				px := getPixel(x+k, y+l, width)
				average := int((img.Pix[px] + img.Pix[px+1] + img.Pix[px+2]) / 3)
				if x == 128 && y == 122 {
					fmt.Println("fast", m, n, k, l, average, px)
				}
				gradientX += utilities.MaxMin(average, 255, 0) * sobelHorizontal[m][n]
				gradientY += utilities.MaxMin(average, 255, 0) * sobelVertical[m][n]
			}
		}
		colorCode := 255 - utilities.MaxMin(math.Sqrt(float64(gradientX*gradientX+gradientY*gradientY)), 255, 0)
		img.Pix[i] = uint8(colorCode)
		img.Pix[i+1] = uint8(colorCode)
		img.Pix[i+2] = uint8(colorCode)
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
