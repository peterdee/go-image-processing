package filters

import (
	"math"
	"runtime"
	"sync"
	"time"

	"go-image-processing/utilities"
)

func Contrast(path string, amount int) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	amount = utilities.MaxMin(amount, 255, -255)
	factor := float64(259*(amount+255)) / float64(255*(259-amount))

	pixLen := len(img.Pix)
	threads := runtime.NumCPU()
	pixPerThread := getPixPerThread(pixLen, threads)

	var wg sync.WaitGroup

	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := clampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			img.Pix[i] = uint8(utilities.MaxMin(factor*(float64(img.Pix[i])-128)+128, 255, 0))
			img.Pix[i+1] = uint8(utilities.MaxMin(factor*(float64(img.Pix[i+1])-128)+128, 255, 0))
			img.Pix[i+2] = uint8(utilities.MaxMin(factor*(float64(img.Pix[i+2])-128)+128, 255, 0))
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
