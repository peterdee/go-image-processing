package filters

import (
	"math"
	"time"

	"go-image-processing/utilities"
)

func Sepia(path string) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	for i := 0; i < len(img.Pix); i += 4 {
		r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
		dR := utilities.MaxMin(0.393*float64(r)+0.769*float64(g)+0.189*float64(b), 255.0, 0.0)
		dG := utilities.MaxMin(0.349*float64(r)+0.686*float64(g)+0.168*float64(b), 255.0, 0.0)
		dB := utilities.MaxMin(0.272*float64(r)+0.534*float64(g)+0.131*float64(b), 255.0, 0.0)
		img.Pix[i], img.Pix[i+1], img.Pix[i+2] = uint8(dR), uint8(dG), uint8(dB)
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
