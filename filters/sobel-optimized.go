package filters

import (
	"math"
	"runtime"
	"sync"
	"time"

	"go-image-processing/utilities"
)

var sobelHorizontal = [3][3]int{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}

var sobelVertical = [3][3]int{
	{1, 2, 1},
	{0, 0, 0},
	{-1, -2, -1},
}

func Sobel(path string) {
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
			gradientX, gradientY := 0, 0
			for m := 0; m < 3; m += 1 {
				for n := 0; n < 3; n += 1 {
					px := getPixel(
						utilities.MaxMin(x-(len(laplacianKernel)/2-m), width-1, 0),
						utilities.MaxMin(y-(len(laplacianKernel)/2-n), height-1, 0),
						width,
					)
					average := (int(img.Pix[px]) + int(img.Pix[px+1]) + int(img.Pix[px+2])) / 3
					gradientX += average * sobelHorizontal[m][n]
					gradientY += average * sobelVertical[m][n]
				}
			}
			channel := uint8(
				255 - utilities.MaxMin(
					math.Sqrt(float64(gradientX*gradientX+gradientY*gradientY)),
					255,
					0,
				),
			)
			result[i], result[i+1], result[i+2] = channel, channel, channel
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
