package filters

import (
	"math"
	"runtime"
	"sync"
	"time"

	"go-image-processing/constants"
)

func Flip(path, flipType string) {
	if flipType != constants.FLIP_TYPE_HORIZONTAL &&
		flipType != constants.FLIP_TYPE_VERTICAL {
		flipType = constants.FLIP_TYPE_HORIZONTAL
	}
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	widthCorrection, heightCorrection := 0, 0
	if width%2 != 0 {
		widthCorrection = 1
	}
	if height%2 != 0 {
		heightCorrection = 1
	}

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
			var j int
			skip := true
			if flipType == constants.FLIP_TYPE_HORIZONTAL && x < width/2+widthCorrection {
				j = getPixel(width-x-1, y, width)
				skip = false
			}
			if flipType == constants.FLIP_TYPE_VERTICAL && y < height/2+heightCorrection {
				j = getPixel(x, height-y-1, width)
				skip = false
			}
			if !skip {
				r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
				img.Pix[i], img.Pix[i+1], img.Pix[i+2] = img.Pix[j], img.Pix[j+1], img.Pix[j+2]
				img.Pix[j], img.Pix[j+1], img.Pix[j+2] = r, g, b
			}
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
