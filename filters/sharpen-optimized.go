package filters

import (
	"math"
	"runtime"
	"sync"
	"time"

	"go-image-processing/utilities"
)

var sharpenKernel = [3][3]int{
	{-1, -1, -1},
	{-1, 9, -1},
	{-1, -1, -1},
}

func Sharpen(path string, amount uint) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	mix := float64(utilities.MaxMin(amount, 100, 0)) / 100
	width, height := img.Rect.Max.X, img.Rect.Max.Y

	pixLen := len(img.Pix)
	threads := runtime.NumCPU()
	pixPerThread := getPixPerThread(pixLen, threads)
	result := make([]uint8, pixLen)

	var wg sync.WaitGroup

	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := clampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			dR, dG, dB := 0, 0, 0
			x, y := getCoordinates(i/4, width)
			for m := 0; m < 3; m += 1 {
				for n := 0; n < 3; n += 1 {
					k := getGradientPoint(x, m, width)
					l := getGradientPoint(y, n, height)
					px := getPixel(x+k, y+l, width)
					dR += int(img.Pix[px]) * sharpenKernel[m][n]
					dG += int(img.Pix[px+1]) * sharpenKernel[m][n]
					dB += int(img.Pix[px+2]) * sharpenKernel[m][n]
				}
			}
			result[i] = uint8(
				utilities.MaxMin(float64(dR)*mix+float64(img.Pix[i])*(1-mix), 255, 0),
			)
			result[i+1] = uint8(
				utilities.MaxMin(float64(dG)*mix+float64(img.Pix[i+1])*(1-mix), 255, 0),
			)
			result[i+2] = uint8(
				utilities.MaxMin(float64(dB)*mix+float64(img.Pix[i+2])*(1-mix), 255, 0),
			)
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
