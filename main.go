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
	img, f := utilities.OpenFile("images/1.png")
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
	// solarize := processing.Solarize(img, 168)
	// sepia := processing.Sepia(img)
	// eight := processing.EightColors(img)
	rotateN := processing.RotateAngle(img, 52)
	// kuwahara := processing.KuwaharaFilter(img)
	// laplasian := processing.LaplasianFilter(img)
	est := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	println(est)
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
	// utilities.SaveFile("solar-"+name, FORMAT, solarize)
	// utilities.SaveFile("sepia-"+name, FORMAT, sepia)
	// utilities.SaveFile("8colors-"+name, FORMAT, eight)
	utilities.SaveFile("rotateN-"+name, FORMAT, rotateN)
	// utilities.SaveFile("kuwahara-"+name, FORMAT, kuwahara)
	// utilities.SaveFile("laplasian-"+name, FORMAT, laplasian)
	// save("rotate180-"+name, rotate180)
	// save("rotate270-"+name, rotate270)
	// save("binary-"+name, binary)
	// save("gray-lum-"+name, grayLum)
	// utilities.SaveFile("box-blur-"+name, FORMAT, boxBlur)
	// save("inverted-"+name, inverted)
}
