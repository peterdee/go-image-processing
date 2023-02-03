package filters

import (
	"math"
	"time"
)

func Binary(path string, threshold uint8) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	for i := 0; i < len(img.Pix); i += 4 {
		average := uint8((int(img.Pix[i]) + int(img.Pix[i+1]) + int(img.Pix[i+2])) / 3)
		channel := uint8(255)
		if average < threshold {
			channel = 0
		}
		img.Pix[i], img.Pix[i+1], img.Pix[i+2] = channel, channel, channel
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
