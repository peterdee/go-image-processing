package filters

import (
	"math"
	"runtime"
	"sync"
	"time"
)

func applySolarizeThreshold(channel, threshold uint8) uint8 {
	if channel < threshold {
		return 255 - channel
	}
	return channel
}

func Solarize(path string, threshold uint8) {
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
			img.Pix[i] = applySolarizeThreshold(img.Pix[i], threshold)
			img.Pix[i+1] = applySolarizeThreshold(img.Pix[i+1], threshold)
			img.Pix[i+2] = applySolarizeThreshold(img.Pix[i+2], threshold)
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
