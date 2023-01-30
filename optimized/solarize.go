package optimized

import (
	"math"
	"time"
)

func applyThreshold(subpixel, threshold uint8) uint8 {
	if subpixel < threshold {
		return 255 - subpixel
	}
	return subpixel
}

func Solarize(path string, threshold uint8) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = applyThreshold(img.Pix[i], threshold)
		img.Pix[i+1] = applyThreshold(img.Pix[i+1], threshold)
		img.Pix[i+2] = applyThreshold(img.Pix[i+2], threshold)
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
