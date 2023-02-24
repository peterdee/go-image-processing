package filters

import (
	"math"
	"runtime"
	"sync"
	"time"
)

func Binary(path string, threshold uint8) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)

	pixLen := len(img.Pix)
	threads := runtime.NumCPU()
	pixPerThreadRaw := float64(pixLen) / float64(threads)
	pixPerThread := int(pixPerThreadRaw + (float64(threads) - math.Mod(pixPerThreadRaw, 4.0)))

	var wg sync.WaitGroup

	processing := func(startIndex int) {
		defer wg.Done()
		endIndex := clampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			average := uint8((int(img.Pix[i]) + int(img.Pix[i+1]) + int(img.Pix[i+2])) / 3)
			channel := uint8(255)
			if average < threshold {
				channel = 0
			}
			img.Pix[i], img.Pix[i+1], img.Pix[i+2] = channel, channel, channel
		}
	}

	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(pixPerThread * t)
	}

	wg.Wait()

	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
