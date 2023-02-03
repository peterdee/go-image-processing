package filters

import (
	"math"
	"time"

	"go-image-processing/utilities"
)

func Brightness(path string, amount int) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	amount = utilities.MaxMin(amount, 255, -255)
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = uint8(utilities.MaxMin(int(img.Pix[i])+amount, 255, 0))
		img.Pix[i+1] = uint8(utilities.MaxMin(int(img.Pix[i+1])+amount, 255, 0))
		img.Pix[i+2] = uint8(utilities.MaxMin(int(img.Pix[i+2])+amount, 255, 0))
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
