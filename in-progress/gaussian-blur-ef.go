package progress

import (
	"math"
	"sync"
	"time"

	"go-image-processing/utilities"
)

func GaussianBlurEF(path string, sigma float64) {
	if sigma < 0 {
		sigma *= -1
	}
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)

	kernel := createKernel(sigma)
	pixLen := len(img.Pix)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	temp := make([]uint8, len(img.Pix))

	var wg sync.WaitGroup

	horizontal := func(i int) {
		defer wg.Done()
		x, y := getCoordinates(i/4, width)
		sumR := 0.0
		sumG := 0.0
		sumB := 0.0
		for k := 0; k < len(kernel); k += 1 {
			dx := utilities.MaxMin(x-(len(kernel)/2-k), width-1, 0)
			px := getPixel(dx, y, width)
			sumR += float64(img.Pix[px]) * kernel[k]
			sumG += float64(img.Pix[px+1]) * kernel[k]
			sumB += float64(img.Pix[px+2]) * kernel[k]
		}
		temp[i] = uint8(utilities.MaxMin(sumR, 255, 0))
		temp[i+1] = uint8(utilities.MaxMin(sumG, 255, 0))
		temp[i+2] = uint8(utilities.MaxMin(sumB, 255, 0))
	}

	vertical := func(i int) {
		defer wg.Done()
		x, y := getCoordinates(i/4, width)
		sumR := 0.0
		sumG := 0.0
		sumB := 0.0
		for k := 0; k < len(kernel); k += 1 {
			dy := utilities.MaxMin(y-(len(kernel)/2-k), height-1, 0)
			px := getPixel(x, dy, width)
			sumR += float64(temp[px]) * kernel[k]
			sumG += float64(temp[px+1]) * kernel[k]
			sumB += float64(temp[px+2]) * kernel[k]
		}
		img.Pix[i] = uint8(utilities.MaxMin(sumR, 255, 0))
		img.Pix[i+1] = uint8(utilities.MaxMin(sumG, 255, 0))
		img.Pix[i+2] = uint8(utilities.MaxMin(sumB, 255, 0))
	}

	// horizontal
	for i := 0; i < pixLen; i += 16 {
		wg.Add(1)
		go horizontal(i)
		if i+12 < pixLen {
			wg.Add(3)
			go horizontal(i + 4)
			go horizontal(i + 8)
			go horizontal(i + 12)
		} else if i+8 < pixLen {
			wg.Add(2)
			go horizontal(i + 4)
			go horizontal(i + 8)
		} else if i+4 < pixLen {
			wg.Add(1)
			go horizontal(i + 4)
		}
	}

	wg.Wait()

	// vertical
	for i := 0; i < pixLen; i += 16 {
		wg.Add(1)
		go vertical(i)
		if i+12 < pixLen {
			wg.Add(3)
			go vertical(i + 4)
			go vertical(i + 8)
			go vertical(i + 12)
		} else if i+8 < pixLen {
			wg.Add(2)
			go vertical(i + 4)
			go vertical(i + 8)
		} else if i+4 < pixLen {
			wg.Add(1)
			go vertical(i + 4)
		}
	}
	wg.Wait()

	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format, 1)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
