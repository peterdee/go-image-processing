package main

import progress "go-image-processing/in-progress"

var FORMAT string

func main() {
	path := "images/big.jpg"
	// img, f, openMS, convertMS := utilities.OpenFile(path)
	// FORMAT = f
	// now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	// rotateN := processing.RotateAngle(img, 52)
	// gauss := progress.GaussianBlur(img, 5)
	// bilateral := progress.Bilateral(img, 3, 10, 15)
	// processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	// name := fmt.Sprintf(`file-%d.%s`, time.Now().Unix(), FORMAT)
	// utilities.SaveFile("rotateN-"+name, FORMAT, rotateN)
	// saveMS := utilities.SaveFile("gauss-"+name, FORMAT, gauss)
	// saveMS := utilities.SaveFile("bilateral-"+name, FORMAT, bilateral)

	progress.BinaryEF(path, 185)

	/* Optimized filters */

	// filters.BilateralSlow(path, 3, 10, 15)
	// filters.Binary(path, 185)
	// filters.BoxBlur(path, 7)
	// filters.Brightness(path, 56)
	// filters.Contrast(path, 225)
	// filters.EightColors(path)
	// filters.Emboss(path)
	// filters.Flip(path, constants.FLIP_TYPE_VERTICAL)
	// filters.GammaCorrection(path, 0.7)
	// filters.Grayscale(path, constants.GRAYSCALE_TYPE_AVERAGE)
	// filters.HueRotate(path, 252)
	// filters.Invert(path)
	// filters.Kuwahara(path, 4)
	// filters.Laplacian(path)
	// filters.RotateFixed(path, constants.ROTATE_FIXED_90)
	// filters.Sepia(path)
	// filters.Sharpen(path, 92)
	// filters.Sobel(path)
	// filters.Solarize(path, 175)

	// sum := openMS + convertMS + processMS + saveMS
	// println("s open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
