package filters

import (
	"math"
	"time"
)

type Color struct {
	R, G, B int
}

var COLORS = [8]Color{
	{255, 0, 0},
	{0, 255, 0},
	{0, 0, 255},
	{255, 255, 0},
	{255, 0, 255},
	{0, 255, 255},
	{255, 255, 255},
	{0, 0, 0},
}

func EightColors(path string) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	for i := 0; i < len(img.Pix); i += 4 {
		minDelta := 195076
		var selectedColor Color
		for j := range COLORS {
			indexColor := COLORS[j]
			rDifference := int(img.Pix[i]) - indexColor.R
			gDifference := int(img.Pix[i+1]) - indexColor.G
			bDifference := int(img.Pix[i+2]) - indexColor.B
			delta := rDifference*rDifference + gDifference*gDifference + bDifference*bDifference
			if delta < minDelta {
				minDelta = delta
				selectedColor = indexColor
			}
		}
		img.Pix[i] = uint8(selectedColor.R)
		img.Pix[i+1] = uint8(selectedColor.G)
		img.Pix[i+2] = uint8(selectedColor.B)
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
