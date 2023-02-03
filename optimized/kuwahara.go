package optimized

import (
	"math"
	"time"
)

func Kuwahara(path string, radius uint) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	radiusInt := int(radius)
	destination := make([]uint8, len(img.Pix))
	apertureMinX := [4]int{-radiusInt, 0, -radiusInt, 0}
	apertureMaxX := [4]int{0, radiusInt, 0, radiusInt}
	apertureMinY := [4]int{-radiusInt, -radiusInt, 0, 0}
	apertureMaxY := [4]int{0, 0, radiusInt, radiusInt}
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := getCoordinates(i/4, width)
		rValues := [4]int{0, 0, 0, 0}
		gValues := [4]int{0, 0, 0, 0}
		bValues := [4]int{0, 0, 0, 0}
		maxRValue := [4]int{0, 0, 0, 0}
		maxGValue := [4]int{0, 0, 0, 0}
		maxBValue := [4]int{0, 0, 0, 0}
		minRValue := [4]int{255, 255, 255, 255}
		minGValue := [4]int{255, 255, 255, 255}
		minBValue := [4]int{255, 255, 255, 255}
		pixelsCount := [4]int{0, 0, 0, 0}
		for i := 0; i < 4; i += 1 {
			x2s, x2e := getAperture(x, width, apertureMinX[i], apertureMaxX[i])
			y2s, y2e := getAperture(y, height, apertureMinY[i], apertureMaxY[i])
			for x2 := x2s; x2 < x2e; x2 += 1 {
				for y2 := y2s; y2 < y2e; y2 += 1 {
					px := getPixel(x2, y2, width)
					r, g, b := img.Pix[px], img.Pix[px+1], img.Pix[px+2]
					rValues[i] += int(r)
					gValues[i] += int(g)
					bValues[i] += int(b)
					if int(r) > maxRValue[i] {
						maxRValue[i] = int(r)
					} else if int(r) < minRValue[i] {
						minRValue[i] = int(r)
					}
					if int(g) > maxGValue[i] {
						maxGValue[i] = int(g)
					} else if int(g) < minGValue[i] {
						minGValue[i] = int(g)
					}
					if int(b) > maxBValue[i] {
						maxBValue[i] = int(b)
					} else if int(b) < minBValue[i] {
						minBValue[i] = int(b)
					}
					pixelsCount[i] += 1
				}
			}
		}
		j := 0
		minDifference := 10000
		for i := 0; i < 4; i += 1 {
			cdR := maxRValue[i] - minRValue[i]
			cdG := maxGValue[i] - minGValue[i]
			cdB := maxBValue[i] - minBValue[i]
			CurrentDifference := cdR + cdG + cdB
			if CurrentDifference < minDifference && pixelsCount[i] > 0 {
				j = i
				minDifference = CurrentDifference
			}
		}
		destination[i] = uint8(rValues[j] / pixelsCount[j])
		destination[i+1] = uint8(gValues[j] / pixelsCount[j])
		destination[i+2] = uint8(bValues[j] / pixelsCount[j])
	}
	img.Pix = destination
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
