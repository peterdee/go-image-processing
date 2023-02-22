package progress

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"runtime"
	"sync"
	"time"
)

func open(path string) (*image.RGBA, string, int, int) {
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open the file: ", err)
	}
	defer file.Close()
	openMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	now2 := math.Round(float64(time.Now().UnixNano()) / 1000000)
	content, format, err := image.Decode(file)
	if err != nil {
		log.Fatal("Could not decode the file: ", err)
	}
	rect := content.Bounds()
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), content, rect.Min, draw.Src)
	convertMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now2)
	return img, format, openMS, convertMS
}

func save(img *image.RGBA, format string, iteration int) int {
	name := fmt.Sprintf(`ef-file-%d.%s`, iteration%10, format)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	newFile, err := os.Create("images/" + name)
	if err != nil {
		log.Fatal("Could not save the file")
	}
	defer newFile.Close()
	if format == "png" {
		png.Encode(newFile, img.SubImage(img.Rect))
	} else {
		jpeg.Encode(
			newFile,
			img.SubImage(img.Rect),
			&jpeg.Options{
				Quality: 100,
			},
		)
	}
	return int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
}

func processing(path string, threshold uint8, iteration int) int {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)

	sa := len(img.Pix) / 4

	var wg sync.WaitGroup
	wg.Add(4)

	process := func(startIndex, endIndex int, threshold uint8, pixels *[]uint8) {
		defer wg.Done()
		pxs := *pixels
		for i := startIndex; i < endIndex; i += 4 {
			average := uint8((int(pxs[i]) + int(pxs[i+1]) + int(pxs[i+2])) / 3)
			channel := uint8(255)
			if average < threshold {
				channel = 0
			}
			pxs[i], pxs[i+1], pxs[i+2] = channel, channel, channel
		}
	}

	go process(0, sa, threshold, &img.Pix)
	// go process(sa, len(img.Pix), threshold, &img.Pix)
	go process(sa, sa*2, threshold, &img.Pix)
	go process(sa*2, sa*3, threshold, &img.Pix)
	go process(sa*3, len(img.Pix), threshold, &img.Pix)

	wg.Wait()

	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format, iteration)
	sum := openMS + convertMS + processMS + saveMS
	// println("f open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
	return sum
}

func BinaryEF(path string, threshold uint8) {
	iterations := 100
	total := 0
	for i := 0; i < iterations; i += 1 {
		total += processing(path, threshold, i)
	}
	println("avg", total/iterations, "CPUs:", runtime.NumCPU())
}