package filters

import (
	"math"
	"time"

	"go-image-processing/utilities"
)

var DEG float64 = math.Pi / 180

func HueRotate(path string, angle int) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	cos := math.Cos(float64(angle) * DEG)
	sin := math.Sin(float64(angle) * DEG)
	matrix := [3]float64{
		cos + (1-cos)/3,
		(1-cos)/3 - math.Sqrt(float64(1)/3)*sin,
		(1-cos)/3 + math.Sqrt(float64(1)/3)*sin,
	}
	for i := 0; i < len(img.Pix); i += 4 {
		r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
		rr := utilities.MaxMin(float64(r)*matrix[0]+float64(g)*matrix[1]+float64(b)*matrix[2], 255, 0)
		rg := utilities.MaxMin(float64(r)*matrix[2]+float64(g)*matrix[0]+float64(b)*matrix[1], 255, 0)
		rb := utilities.MaxMin(float64(r)*matrix[1]+float64(g)*matrix[2]+float64(b)*matrix[0], 255, 0)
		img.Pix[i] = uint8(rr)
		img.Pix[i+1] = uint8(rg)
		img.Pix[i+2] = uint8(rb)
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
