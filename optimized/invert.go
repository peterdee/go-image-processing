package optimized

import (
	"math"
	"time"
)

func Invert(path string) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = 255 - img.Pix[i]
		img.Pix[i+1] = 255 - img.Pix[i+1]
		img.Pix[i+2] = 255 - img.Pix[i+2]
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
