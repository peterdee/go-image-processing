package filters

import (
	"math"
	"runtime"
	"sync"
	"time"
)

func Sepia(path string) {
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
			r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
			dR := clampMax(0.393*float64(r)+0.769*float64(g)+0.189*float64(b), 255.0)
			dG := clampMax(0.349*float64(r)+0.686*float64(g)+0.168*float64(b), 255.0)
			dB := clampMax(0.272*float64(r)+0.534*float64(g)+0.131*float64(b), 255.0)
			img.Pix[i], img.Pix[i+1], img.Pix[i+2] = uint8(dR), uint8(dG), uint8(dB)
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
