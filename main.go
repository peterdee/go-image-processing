package main

import (
	"go-image-processing/optimized"
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

	// optimized.Binary(path, 185)
	// optimized.BoxBlur(path, 7)
	// optimized.Brightness(path, 56)
	// optimized.Contrast(path, 225)
	// optimized.EightColors(path)
	// optimized.Emboss(path)
	// optimized.Flip(path, constants.FLIP_TYPE_VERTICAL)
	// optimized.GammaCorrection(path, 0.7)
	// optimized.Grayscale(path, constants.GRAYSCALE_TYPE_LUMINOSITY)
	// optimized.HueRotate(path, 252)
	// optimized.Invert(path)
	// optimized.Kuwahara(path, 7)
	optimized.Laplacian(path)
	// optimized.RotateFixed(path, constants.ROTATE_FIXED_90)
	// optimized.Sepia(path)
	// optimized.Sharpen(path, 92)
	// optimized.Solarize(path, 175)
	// optimized.Sobel(path)

	// sum := openMS + convertMS + processMS + saveMS
	// println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
