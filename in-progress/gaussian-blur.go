package progress

import (
	"math"
	"time"

	"go-image-processing/utilities"
)

const K float64 = 6

func createKernel(sigma float64) []float64 {
	dim := math.Max(3.0, K*sigma)
	sqrtSigmaPi2 := math.Sqrt(math.Pi*2.0) * sigma
	s2 := 2.0 * sigma * sigma
	sum := 0.0
	kDim := dim
	if int(kDim)%2 == 0 {
		kDim = dim - 1
	}
	kernel := make([]float64, int(kDim))
	half := int(len(kernel) / 2)
	i := -half
	for j := 0; j < len(kernel); j += 1 {
		kernel[j] = math.Exp(-float64(i*i)/(s2)) / sqrtSigmaPi2
		sum += kernel[j]
		i += 1
	}
	for k := 0; k < int(kDim); k += 1 {
		kernel[k] /= sum
	}
	return kernel
}

func getCoordinates(pixel, width int) (int, int) {
	return pixel % width, int(math.Floor(float64(pixel) / float64(width)))
}

func getPixel(x, y, width int) int {
	return ((y * width) + x) * 4
}

// Gaussian blur: single thread
func GaussianBlur(path string, sigma float64) {
	if sigma < 0 {
		sigma *= -1
	}
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)

	kernel := createKernel(sigma)
	kernelLen := len(kernel)
	pixLen := len(img.Pix)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	temp := make([]uint8, len(img.Pix))

	// horizontal
	for i := 0; i < pixLen; i += 4 {
		x, y := getCoordinates(i/4, width)
		sumR := 0.0
		sumG := 0.0
		sumB := 0.0
		for k := 0; k < kernelLen; k += 1 {
			dx := utilities.MaxMin(x-(kernelLen/2-k), width-1, 0)
			px := getPixel(dx, y, width)
			sumR += float64(img.Pix[px]) * kernel[k]
			sumG += float64(img.Pix[px+1]) * kernel[k]
			sumB += float64(img.Pix[px+2]) * kernel[k]
		}
		temp[i] = uint8(utilities.MaxMin(sumR, 255, 0))
		temp[i+1] = uint8(utilities.MaxMin(sumG, 255, 0))
		temp[i+2] = uint8(utilities.MaxMin(sumB, 255, 0))
	}

	// vertical
	for i := 0; i < pixLen; i += 4 {
		x, y := getCoordinates(i/4, width)
		sumR := 0.0
		sumG := 0.0
		sumB := 0.0
		for k := 0; k < kernelLen; k += 1 {
			dy := utilities.MaxMin(y-(kernelLen/2-k), height-1, 0)
			px := getPixel(x, dy, width)
			sumR += float64(temp[px]) * kernel[k]
			sumG += float64(temp[px+1]) * kernel[k]
			sumB += float64(temp[px+2]) * kernel[k]
		}
		img.Pix[i] = uint8(utilities.MaxMin(sumR, 255, 0))
		img.Pix[i+1] = uint8(utilities.MaxMin(sumG, 255, 0))
		img.Pix[i+2] = uint8(utilities.MaxMin(sumB, 255, 0))
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format, 1)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
