package progress

import (
	"math"
	"runtime"
	"time"

	"go-image-processing/utilities"
)

func clampMax[T float64 | int | uint8](value, max T) T {
	if value > max {
		return max
	}
	return value
}

// Gaussian blur: different approach (channels) - top 2
func GaussianBlurDA(path string, sigma float64) {
	if sigma < 0 {
		sigma *= -1
	}
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)

	kernel := createKernel(sigma)
	pixLen := len(img.Pix)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	temp := make([]uint8, len(img.Pix))

	threads := runtime.NumCPU()
	pixPerThreadRaw := float64(pixLen) / float64(threads)
	pixPerThread := int(pixPerThreadRaw + (float64(threads) - math.Mod(pixPerThreadRaw, 4.0)))

	processing := func(start int, direction string, ch chan int, thread int) {
		end := clampMax(start+pixPerThread, pixLen)
		for i := start; i < end; i += 4 {
			x, y := getCoordinates(i/4, width)
			sumR := 0.0
			sumG := 0.0
			sumB := 0.0
			for k := 0; k < len(kernel); k += 1 {
				var px int
				if direction == "horizontal" {
					px = getPixel(
						utilities.MaxMin(x-(len(kernel)/2-k), width-1, 0),
						y,
						width,
					)
					sumR += float64(img.Pix[px]) * kernel[k]
					sumG += float64(img.Pix[px+1]) * kernel[k]
					sumB += float64(img.Pix[px+2]) * kernel[k]
				} else {
					px = getPixel(
						x,
						utilities.MaxMin(y-(len(kernel)/2-k), height-1, 0),
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
		}
		ch <- thread
	}

	// horizontal
	ch := make(chan int)
	resH := make([]int, 0)
	for t := 0; t < threads; t += 1 {
		go processing(pixPerThread*t, "horizontal", ch, t)
	}
	for {
		result := <-ch
		resH = append(resH, result)
		if len(resH) == threads {
			break
		}
	}

	// vertical
	resV := make([]int, 0)
	for t := 0; t < threads; t += 1 {
		go processing(pixPerThread*t, "vertical", ch, t)
	}
	for {
		result := <-ch
		resV = append(resV, result)
		if len(resV) == threads {
			close(ch)
			break
		}
	}

	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format, 1)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
	println("threads", threads)
}
