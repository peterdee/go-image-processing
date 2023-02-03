package main

import (
	"go-image-processing/filters"
)

var FORMAT string

func main() {
	path := "images/11.jpeg"
	// img, f, openMS, convertMS := utilities.OpenFile(path)
	// FORMAT = f
	// now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	// rotateN := processing.RotateAngle(img, 52)
	// gauss := processing.GaussianBlur(img, 25)
	// processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	// println(est)
	// name := fmt.Sprintf(`file-%d.%s`, time.Now().Unix(), FORMAT)
	// utilities.SaveFile("rotateN-"+name, FORMAT, rotateN)
	// saveMS := utilities.SaveFile("gauss-"+name, FORMAT, gauss)

	/* Optimized filters */

	// filters.Binary(path, 185)
	// filters.BoxBlur(path, 7)
	// filters.Brightness(path, 56)
	// filters.Contrast(path, 225)
	// filters.EightColors(path)
	// filters.Emboss(path)
	// filters.Flip(path, constants.FLIP_TYPE_VERTICAL)
	// filters.GammaCorrection(path, 0.7)
	// filters.Grayscale(path, constants.GRAYSCALE_TYPE_LUMINOSITY)
	// filters.HueRotate(path, 252)
	// filters.Invert(path)
	// filters.Kuwahara(path, 7)
	filters.Laplacian(path)
	// filters.RotateFixed(path, constants.ROTATE_FIXED_90)
	// filters.Sepia(path)
	// filters.Sharpen(path, 92)
	// filters.Solarize(path, 175)
	// filters.Sobel(path)

	// sum := openMS + convertMS + processMS + saveMS
	// println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
