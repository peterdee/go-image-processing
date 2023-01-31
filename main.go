package main

import (
	"fmt"
	"math"
	"time"

	"go-image-processing/optimized"
	"go-image-processing/processing"
	"go-image-processing/utilities"
)

var FORMAT string

func main() {
	path := "images/5.png"
	img, f, openMS, convertMS := utilities.OpenFile(path)
	FORMAT = f
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	// flippedV := processing.FlipVertical(img)
	// boxBlur := processing.BoxBlur(img, 7)
	// flippedH := processing.FlipHorizontal(img)
	// rotate90 := processing.Rotate90(img)
	// rotate270 := processing.Rotate270(img)
	// rotate180 := processing.Rotate180(img)
	sobel := processing.SobelFilter(img)
	// emboss := processing.EmbossFilter(img)
	// rotateN := processing.RotateAngle(img, 52)
	// kuwahara := processing.KuwaharaFilter(img, 5)
	// gauss := processing.GaussianBlur(img, 25)
	// laplasian := processing.LaplasianFilter(img)
	// sharpen := processing.SharpenFilter(img, 99)
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	// println(est)
	name := fmt.Sprintf(`file-%d.%s`, time.Now().Unix(), FORMAT)
	// utilities.SaveFile("flippedV-"+name, FORMAT, flippedV)
	// utilities.SaveFile("flippedH-"+name, FORMAT, flippedH)
	// save("rotate90-"+name, rotate90)
	saveMS := utilities.SaveFile("sobel-"+name, FORMAT, sobel)
	// utilities.SaveFile("emboss-"+name, FORMAT, emboss)
	// utilities.SaveFile("rotateN-"+name, FORMAT, rotateN)
	// saveMS := utilities.SaveFile("kuwahara-"+name, FORMAT, kuwahara)
	// utilities.SaveFile("gauss-"+name, FORMAT, gauss)
	// utilities.SaveFile("laplasian-"+name, FORMAT, laplasian)
	// saveMS := utilities.SaveFile("sharp-"+name, FORMAT, sharpen)
	// save("rotate180-"+name, rotate180)
	// save("rotate270-"+name, rotate270)
	// utilities.SaveFile("box-blur-"+name, FORMAT, boxBlur)

	/* Optimized filters */

	// optimized.Binary(path, 185)
	// optimized.Brightness(path, 56)
	// optimized.Contrast(path, 225)
	// optimized.EightColors(path)
	// optimized.GammaCorrection(path, 0.7)
	// optimized.Grayscale(path, constants.GRAYSCALE_TYPE_LUMINOSITY)
	// optimized.HueRotate(path, 252)
	// optimized.Invert(path)
	// optimized.Sepia(path)
	// optimized.Solarize(path, 175)
	optimized.Sobel(path)

	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
