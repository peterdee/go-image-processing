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

// Binary filter that uses channels - slightly slower than WaitGroup
func BinaryCH(path string, threshold uint8) {
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)

	pixLen := len(img.Pix)
	threads := runtime.NumCPU()
	pixPerThreadRaw := float64(pixLen) / float64(threads)
	pixPerThread := int(pixPerThreadRaw + (float64(threads) - math.Mod(pixPerThreadRaw, 4.0)))

	processing := func(startIndex int, ch chan int, thread int) {
		endIndex := clampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			average := uint8((int(img.Pix[i]) + int(img.Pix[i+1]) + int(img.Pix[i+2])) / 3)
			channel := uint8(255)
			if average < threshold {
				channel = 0
			}
			img.Pix[i], img.Pix[i+1], img.Pix[i+2] = channel, channel, channel
		}
		ch <- thread
	}

	ch := make(chan int)
	done := make([]int, 0)
	for t := 0; t < threads; t += 1 {
		go processing(pixPerThread*t, ch, t)
	}
	for {
		result := <-ch
		done = append(done, result)
		if len(done) == threads {
			close(ch)
			break
		}
	}

	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := save(img, format, 1)
	sum := openMS + convertMS + processMS + saveMS
	println("f open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
