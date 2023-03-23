package main

import (
	"fmt"
	progress "go-image-processing/in-progress"
	"go-image-processing/utilities"
	"math"
	"time"
)

var FORMAT string

func main() {
	path := "images/15.jpeg"
	img, f, openMS, convertMS := utilities.OpenFile(path)
	FORMAT = f
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	bilateral := progress.Bilateral(img, 3, 10, 15)
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	name := fmt.Sprintf(`file-%d.%s`, time.Now().Unix(), FORMAT)
	saveMS := utilities.SaveFile("bilateral-"+name, FORMAT, bilateral)

	/* Optimized filters */

	// filters.Binary(path, 122)
	// filters.BoxBlur(path, 10)
	// filters.Brightness(path, 256)
	// filters.Contrast(path, 205)
	// filters.EightColors(path)
	// filters.Emboss(path)
	// filters.Flip(path, constants.FLIP_TYPE_HORIZONTAL)
	// filters.GammaCorrection(path, 0.1)
	// filters.GaussianBlur(path, 5.2)
	// filters.Grayscale(path, constants.GRAYSCALE_TYPE_LUMINANCE)
	// filters.HueRotate(path, 252)
	// filters.Invert(path)
	// filters.Kuwahara(path, 4)
	// filters.Laplacian(path)
	// filters.RotateFixed(path, constants.ROTATE_FIXED_90)
	// filters.Sepia(path)
	// filters.Sharpen(path, 92)
	// filters.Sobel(path)
	// filters.Solarize(path, 100)

	/* Alternative implementations */

	// progress.BinaryCH(path, 122)
	// progress.GaussianBlurCH(path, 10.0)
	// progress.GaussianBlurCHSlow(path, 10.0)

	sum := openMS + convertMS + processMS + saveMS
	println("s open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
