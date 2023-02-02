package optimized

import (
	"go-image-processing/constants"
	"math"
	"time"
)

func Flip(path, flipType string) {
	if flipType != constants.FLIP_TYPE_HORIZONTAL &&
		flipType != constants.FLIP_TYPE_VERTICAL {
		flipType = constants.FLIP_TYPE_HORIZONTAL
	}
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	width, _ := img.Rect.Max.X, img.Rect.Max.Y
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := getCoordinates(i/4, width)
		if x < width/2 && flipType == constants.FLIP_TYPE_HORIZONTAL {
			z := width - x - 1
			px := getPixel(z, y, width)
			r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
			img.Pix[i] = img.Pix[px]
			img.Pix[i+1] = img.Pix[px+1]
			img.Pix[i+2] = img.Pix[px+2]
			img.Pix[px] = r
			img.Pix[px+1] = g
			img.Pix[px+2] = b
		}
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
