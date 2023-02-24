package progress

import (
	"math"
	"runtime"
	"time"

	"go-image-processing/utilities"
)

// Gaussian blur that uses channels - slower version
func GaussianBlurCHSlow(path string, sigma float64) {
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
	threads := runtime.NumCPU()

	signal := make(chan struct{}, threads)
	step := 4

	processing := func(i int, direction string) {
		x, y := getCoordinates(i/4, width)
		sumR := 0.0
		sumG := 0.0
		sumB := 0.0
		for k := 0; k < kernelLen; k += 1 {
			var px int
			if direction == "horizontal" {
				px = getPixel(
					utilities.MaxMin(x-(kernelLen/2-k), width-1, 0),
					y,
					width,
				)
				sumR += float64(img.Pix[px]) * kernel[k]
				sumG += float64(img.Pix[px+1]) * kernel[k]
				sumB += float64(img.Pix[px+2]) * kernel[k]
			} else {
				px = getPixel(
					x,
					utilities.MaxMin(y-(kernelLen/2-k), height-1, 0),
					width,
				)
				sumR += float64(temp[px]) * kernel[k]
				sumG += float64(temp[px+1]) * kernel[k]
				sumB += float64(temp[px+2]) * kernel[k]
			}
		}
		if direction == "horizontal" {
			temp[i] = uint8(utilities.MaxMin(sumR, 255, 0))
			temp[i+1] = uint8(utilities.MaxMin(sumG, 255, 0))
			temp[i+2] = uint8(utilities.MaxMin(sumB, 255, 0))
		} else {
			img.Pix[i] = uint8(utilities.MaxMin(sumR, 255, 0))
			img.Pix[i+1] = uint8(utilities.MaxMin(sumG, 255, 0))
			img.Pix[i+2] = uint8(utilities.MaxMin(sumB, 255, 0))
		}
		<-signal
	}

	// horizontal
	i := 0
	for {
		if i >= pixLen {
			break
		}
		signal <- struct{}{}
		go processing(i, "horizontal")
		i += step
	}

	// vertical
	i = 0
	for {
		if i >= pixLen {
			close(signal)
			break
		}
		signal <- struct{}{}
		go processing(i, "vertical")
		i += step
	}

	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format, 1)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
	println("threads", threads)
}
