package progress

import (
	"math"
	"runtime"
	"sync"
	"time"
)

func processingSlices(path string, threshold uint8, iteration int) int {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)

	sa := len(img.Pix) / 4

	var wg sync.WaitGroup
	wg.Add(4)

	process := func(threshold uint8, pixels *[]uint8) {
		defer wg.Done()
		pxs := *pixels
		for i := 0; i < len(pxs); i += 4 {
			average := uint8((int(pxs[i]) + int(pxs[i+1]) + int(pxs[i+2])) / 3)
			channel := uint8(255)
			if average < threshold {
				channel = 0
			}
			pxs[i], pxs[i+1], pxs[i+2] = channel, channel, channel
		}
	}

	s1 := img.Pix[:sa]
	s2 := img.Pix[sa : sa*2]
	s3 := img.Pix[sa*2 : sa*3]
	s4 := img.Pix[sa*3:]
	go process(threshold, &s1)
	go process(threshold, &s2)
	go process(threshold, &s3)
	go process(threshold, &s4)

	wg.Wait()

	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format, iteration)
	sum := openMS + convertMS + processMS + saveMS
	// println("f open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
	return sum
}

func BinaryEFSlices(path string, threshold uint8) {
	iterations := 100
	total := 0
	for i := 0; i < iterations; i += 1 {
		total += processingSlices(path, threshold, i)
	}
	println("avg", total/iterations, "CPUs:", runtime.NumCPU())
}
