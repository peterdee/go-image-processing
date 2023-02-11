package filters

import (
	"go-image-processing/utilities"
	"math"
	"time"
)

func distance(x, y, i, j int) float64 {
	return math.Sqrt(math.Pow(float64(x-i), 2) + math.Pow(float64(y-j), 2))
}

func gaussian(x int, sigma float64) float64 {
	return (1.0 / (2 * math.Pi * (math.Pow(sigma, 2)))) * math.Exp(-(math.Pow(float64(x), 2))/(2*math.Pow(sigma, 2)))
}

func BilateralSlow(path string, radius uint, sigmaI, sigmaS float64) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	radiusInt := int(radius)
	diameter := radiusInt * 2
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := getCoordinates(i/4, width)

		b := 20
		if x < b || x > width-b || y < b || y > height-b {
			continue
		}

		fR, fG, fB := 0.0, 0.0, 0.0
		wR, wG, wB := 0.0, 0.0, 0.0

		for i := 0; i < diameter; i += 1 {
			for j := 0; j < diameter; j += 1 {
				nx := x - (radiusInt - i)
				ny := y - (radiusInt - j)

				npx := getPixel(nx, ny, width)
				nr, ng, nb := img.Pix[npx], img.Pix[npx+1], img.Pix[npx+2]

				giR := gaussian(int(nr-img.Pix[i]), sigmaI)
				giG := gaussian(int(ng-img.Pix[i+1]), sigmaI)
				giB := gaussian(int(nb-img.Pix[i+2]), sigmaI)
				gs := gaussian(int(distance(nx, ny, x, y)), sigmaS)
				wiR := giR * gs
				wiG := giG * gs
				wiB := giB * gs
				fR += float64(nr) * wiR
				fG += float64(ng) * wiG
				fB += float64(nb) * wiB
				wR += wiR
				wG += wiG
				wB += wiB
			}
		}
		fR /= wR
		fG /= wG
		fB /= wB
		R := uint8(utilities.MaxMin(fR, 255, 0))
		G := uint8(utilities.MaxMin(fG, 255, 0))
		B := uint8(utilities.MaxMin(fB, 255, 0))
		img.Pix[i], img.Pix[i+1], img.Pix[i+2] = R, G, B
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("f open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
