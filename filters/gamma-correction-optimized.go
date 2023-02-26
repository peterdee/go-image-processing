package filters

import (
	"math"
	"runtime"
	"sync"
	"time"

	"go-image-processing/utilities"
)

func GammaCorrection(path string, amount float64) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	amount = utilities.MaxMin(amount, 3.99, 0)
	power := 1 / amount

	pixLen := len(img.Pix)
	threads := runtime.NumCPU()
	pixPerThread := getPixPerThread(pixLen, threads)

	var wg sync.WaitGroup

	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := clampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			img.Pix[i] = uint8(255 * math.Pow(float64(img.Pix[i])/255, power))
			img.Pix[i+1] = uint8(255 * math.Pow(float64(img.Pix[i+1])/255, power))
			img.Pix[i+2] = uint8(255 * math.Pow(float64(img.Pix[i+2])/255, power))
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
