package main

import (
	"fmt"
	"math"
	"time"

	"go-image-processing/processing"
	"go-image-processing/utilities"
)

var FORMAT string

func main() {
	path := "images/10.jpg"
	img, f, openMS, convertMS := utilities.OpenFile(path)
	FORMAT = f
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	// flippedV := processing.FlipVertical(img)
	// gray := processing.Grayscale(img)
	// grayLum := processing.GrayscaleLuminocity(img)
	// boxBlur := processing.BoxBlur(img, 7)
	// binary := processing.Binary(img, 185)
	// inverted := processing.Invert(img)
	// flippedH := processing.FlipHorizontal(img)
	// rotate90 := processing.Rotate90(img)
	// rotate270 := processing.Rotate270(img)
	// rotate180 := processing.Rotate180(img)
	// sobel := processing.SobelFilter(img)
	// emboss := processing.EmbossFilter(img)
	// gamma := processing.GammaCorrection(img, 0)
	// bright := processing.Brightness(img, -2225)
	// contrast := processing.Contrast(img, 225)
	// solarize := processing.Solarize(img, 175)
	// sepia := processing.Sepia(img)
	// eight := processing.EightColors(img)
	// rotateN := processing.RotateAngle(img, 52)
	// hue := processing.HueRotate(img, 52)
	kuwahara := processing.KuwaharaFilter(img, 5)
	// gauss := processing.GaussianBlur(img, 25)
	// laplasian := processing.LaplasianFilter(img)
	// sharpen := processing.SharpenFilter(img)
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	// println(est)
	name := fmt.Sprintf(`file-%d.%s`, time.Now().Unix(), FORMAT)
	// save("gray-"+name, gray)
	// utilities.SaveFile("flippedV-"+name, FORMAT, flippedV)
	// utilities.SaveFile("flippedH-"+name, FORMAT, flippedH)
	// save("rotate90-"+name, rotate90)
	// utilities.SaveFile("sobel-"+name, FORMAT, sobel)
	// utilities.SaveFile("emboss-"+name, FORMAT, emboss)
	// utilities.SaveFile("gamma-"+name, FORMAT, gamma)
	// utilities.SaveFile("bright-"+name, FORMAT, bright)
	// utilities.SaveFile("contrast-"+name, FORMAT, contrast)
	// saveMS := utilities.SaveFile("solar-"+name, FORMAT, solarize)
	// utilities.SaveFile("sepia-"+name, FORMAT, sepia)
	// saveMS := utilities.SaveFile("eight-colors-"+name, FORMAT, eight)
	// utilities.SaveFile("rotateN-"+name, FORMAT, rotateN)
	// utilities.SaveFile("hue-"+name, FORMAT, hue)
	saveMS := utilities.SaveFile("kuwahara-"+name, FORMAT, kuwahara)
	// utilities.SaveFile("gauss-"+name, FORMAT, gauss)
	// utilities.SaveFile("laplasian-"+name, FORMAT, laplasian)
	// utilities.SaveFile("sharp-"+name, FORMAT, sharpen)
	// save("rotate180-"+name, rotate180)
	// save("rotate270-"+name, rotate270)
	// saveMS := utilities.SaveFile("binary-"+name, FORMAT, binary)
	// save("gray-lum-"+name, grayLum)
	// utilities.SaveFile("box-blur-"+name, FORMAT, boxBlur)
	// save("inverted-"+name, inverted)

	// optimized.Binary(path, 185)
	// optimized.EightColors(path)
	// optimized.Solarize(path, 175)

	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
