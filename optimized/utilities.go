package optimized

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
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

func save(img *image.RGBA, format string) int {
	name := fmt.Sprintf(`file-%d.%s`, time.Now().Unix(), format)
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
