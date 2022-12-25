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
	img, f := utilities.OpenFile("images/4.jpg")
	FORMAT = f
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	// flippedV := processing.FlipVertical(img)
	// gray := processing.Grayscale(img)
	// grayLum := processing.GrayscaleLuminocity(img)
	// boxBlur := processing.BoxBlur(img, 12)
	// binary := processing.Binary(img, 185)
	// inverted := processing.Invert(img)
	// flippedH := processing.FlipHorizontal(img)
	// rotate90 := processing.Rotate90(img)
	// rotate270 := processing.Rotate270(img)
	// rotate180 := processing.Rotate180(img)
	// sobel := processing.SobelFilter(img)
	// emboss := processing.EmbossFilter(img)
	kuwahara := processing.KuwaharaFilter(img)
	// laplasian := processing.LaplasianFilter(img)
	est := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	println(est)
	name := fmt.Sprintf(`file-%d.%s`, time.Now().Unix(), FORMAT)
	// save("gray-"+name, gray)
	// save("flippedV-"+name, flippedV)
	// save("flippedH-"+name, flippedH)
	// save("rotate90-"+name, rotate90)
	// save("sobel-"+name, sobel)
	// save("emboss-"+name, emboss)
	utilities.SaveFile("kuwahara-"+name, FORMAT, kuwahara)
	// utilities.SaveFile("laplasian-"+name, FORMAT, laplasian)
	// save("rotate180-"+name, rotate180)
	// save("rotate270-"+name, rotate270)
	// save("binary-"+name, binary)
	// save("gray-lum-"+name, grayLum)
	// save("box-blur-"+name, boxBlur)
	// save("inverted-"+name, inverted)
}
