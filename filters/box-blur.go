package filters

import (
	"math"
	"time"
)

func getAperture(axisValue, axisMax, apertureMin, apertureMax int) (int, int) {
	start, end := 0, axisMax
	if axisValue+apertureMin > 0 {
		start = axisValue + apertureMin
	}
	if axisValue+apertureMax < axisMax {
		end = axisValue + apertureMax
	}
	return start, end
}

func BoxBlur(path string, radius uint) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	radiusInt := int(radius)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := getCoordinates(i/4, width)
		dR, dG, dB := 0, 0, 0
		pixelCount := 0
		x2s, x2e := getAperture(x, width, -radiusInt, radiusInt)
		y2s, y2e := getAperture(y, height, -radiusInt, radiusInt)
		for x2 := x2s; x2 < x2e; x2 += 1 {
			for y2 := y2s; y2 < y2e; y2 += 1 {
				px := getPixel(x2, y2, width)
				dR += int(img.Pix[px])
				dG += int(img.Pix[px+1])
				dB += int(img.Pix[px+2])
				pixelCount += 1
			}
		}
		img.Pix[i] = uint8(dR / pixelCount)
		img.Pix[i+1] = uint8(dG / pixelCount)
		img.Pix[i+2] = uint8(dB / pixelCount)
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
