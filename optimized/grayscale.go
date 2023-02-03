package optimized

import (
	"math"
	"time"

	"go-image-processing/constants"
)

func Grayscale(path, grayscaleType string) {
	if grayscaleType != constants.GRAYSCALE_TYPE_AVERAGE &&
		grayscaleType != constants.GRAYSCALE_TYPE_LUMINANCE {
		grayscaleType = constants.GRAYSCALE_TYPE_AVERAGE
	}
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	for i := 0; i < len(img.Pix); i += 4 {
		var subpixel uint8
		if grayscaleType == constants.GRAYSCALE_TYPE_AVERAGE {
			subpixel = uint8((int(img.Pix[i]) + int(img.Pix[i+1]) + int(img.Pix[i+2])) / 3)
		} else {
			subpixel = uint8(
				(float64(img.Pix[i])*0.21 + float64(img.Pix[i+1])*0.72 + float64(img.Pix[i+2])*0.07),
			)
		}
		img.Pix[i] = subpixel
		img.Pix[i+1] = subpixel
		img.Pix[i+2] = subpixel
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
