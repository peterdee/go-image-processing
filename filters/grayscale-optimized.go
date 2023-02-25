package filters

import (
	"math"
	"runtime"
	"sync"
	"time"

	"go-image-processing/constants"
)

func Grayscale(path, grayscaleType string) {
	if grayscaleType != constants.GRAYSCALE_TYPE_AVERAGE &&
		grayscaleType != constants.GRAYSCALE_TYPE_LUMINANCE {
		grayscaleType = constants.GRAYSCALE_TYPE_AVERAGE
	}
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)

	pixLen := len(img.Pix)
	threads := runtime.NumCPU()
	pixPerThread := getPixPerThread(pixLen, threads)

	var wg sync.WaitGroup

	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := clampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			var channel uint8
			if grayscaleType == constants.GRAYSCALE_TYPE_AVERAGE {
				channel = uint8((int(img.Pix[i]) + int(img.Pix[i+1]) + int(img.Pix[i+2])) / 3)
			} else {
				channel = uint8(
					(float64(img.Pix[i])*0.21 + float64(img.Pix[i+1])*0.72 + float64(img.Pix[i+2])*0.07),
				)
			}
			img.Pix[i], img.Pix[i+1], img.Pix[i+2] = channel, channel, channel
		}
	}

	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(t)
	}

	wg.Wait()

	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
