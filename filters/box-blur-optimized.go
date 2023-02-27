package filters

import (
	"math"
	"runtime"
	"sync"
	"time"
)

func BoxBlur(path string, radius uint) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	radiusInt := int(radius)
	width, height := img.Rect.Max.X, img.Rect.Max.Y

	pixLen := len(img.Pix)
	threads := runtime.NumCPU()
	pixPerThread := getPixPerThread(pixLen, threads)
	result := make([]uint8, len(img.Pix))

	var wg sync.WaitGroup

	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := clampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			x, y := getCoordinates(i/4, width)
			dR, dG, dB := 0, 0, 0
			pixelCount := 0
			x2s, x2e := getAperture(x, width, -radiusInt, radiusInt)
			y2s, y2e := getAperture(y, height, -radiusInt, radiusInt)
			for x2 := x2s; x2 < x2e; x2 += 1 {
				for y2 := y2s; y2 < y2e; y2 += 1 {
					px := getPixel(x2, y2, width)
					dR += int(img.Pix[px])
					dG += int(img.Pix[px+1])
					dB += int(img.Pix[px+2])
					pixelCount += 1
				}
			}
			result[i] = uint8(dR / pixelCount)
			result[i+1] = uint8(dG / pixelCount)
			result[i+2] = uint8(dB / pixelCount)
			result[i+3] = img.Pix[i+3]
		}
	}

	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(t)
	}

	wg.Wait()

	img.Pix = result

	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
