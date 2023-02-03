package utilities

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"time"
)

func OpenFile(path string) ([][]color.Color, string, int, int) {
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
	rgba := image.NewRGBA(image.Rect(0, 0, rect.Dx(), rect.Dy()))
	draw.Draw(rgba, rgba.Bounds(), content, rect.Min, draw.Src)

	var grid [][]color.Color
	size := rgba.Bounds().Size()
	for i := 0; i < size.X; i += 1 {
		var y []color.Color
		for j := 0; j < size.Y; j += 1 {
			y = append(y, rgba.At(i, j))
		}
		grid = append(grid, y)
	}
	convertMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now2)
	return grid, format, openMS, convertMS
}

func SaveFile(name, format string, grid [][]color.Color) int {
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	xLen, yLen := len(grid), len(grid[0])
	img := image.NewNRGBA(image.Rect(0, 0, xLen, yLen))
	for x := 0; x < xLen; x += 1 {
		for y := 0; y < yLen; y += 1 {
			img.Set(x, y, grid[x][y])
		}
	}
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

func SaveImage(img *image.RGBA, format string) int {
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
