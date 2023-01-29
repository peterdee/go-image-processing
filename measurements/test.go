package measurements

import (
	"go-image-processing/utilities"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
	"time"
)

func OpenFile(path string) (image.Image, string, int) {
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open the file: ", err)
	}
	defer file.Close()
	content, format, err := image.Decode(file)
	if err != nil {
		log.Fatal("Could not decode the file: ", err)
	}
	est := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	return content, format, est
}

func Binary(name, path string, threshold uint8) {
	img, format, openEST := OpenFile(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	width, height := img.Bounds().Max.X, img.Bounds().Max.Y
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			gray, alpha := utilities.Gray(img.At(x, y))
			partial := uint8(255)
			if gray < threshold {
				partial = 0
			}
			destination[x][y] = color.RGBA{partial, partial, partial, alpha}
		}
	}
	println("binary", openEST, int(math.Round(float64(time.Now().UnixNano())/1000000)-now))
	utilities.SaveFile(name, format, destination)
}
