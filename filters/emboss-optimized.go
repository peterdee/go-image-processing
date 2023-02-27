package filters

import (
	"math"
	"runtime"
	"sync"
	"time"
)

var embossHorizontal = [3][3]int{
	{0, 0, 0},
	{1, 0, -1},
	{0, 0, 0},
}

var embossVertical = [3][3]int{
	{0, 1, 0},
	{0, 0, 0},
	{0, -1, 0},
}

func Emboss(path string) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	width, height := img.Rect.Max.X, img.Rect.Max.Y

	result := make([]uint8, len(img.Pix))

	pixLen := len(img.Pix)
	threads := runtime.NumCPU()
	pixPerThread := getPixPerThread(pixLen, threads)

	var wg sync.WaitGroup

	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := clampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			x, y := getCoordinates(i/4, width)
			gradientX := 0
			gradientY := 0
			for m := 0; m < 3; m += 1 {
				for n := 0; n < 3; n += 1 {
					k := getGradientPoint(x, m, width)
					l := getGradientPoint(y, n, height)
					px := getPixel(x+k, y+l, width)
					average := (int(img.Pix[px]) + int(img.Pix[px+1]) + int(img.Pix[px+2])) / 3
					gradientX += average * embossHorizontal[m][n]
					gradientY += average * embossVertical[m][n]
				}
			}
			channel := uint8(
				255 - clampMax(
					math.Sqrt(float64(gradientX*gradientX+gradientY*gradientY)),
					255,
				),
			)
			result[i], result[i+1], result[i+2], result[i+3] = channel, channel, channel, img.Pix[i+3]
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
