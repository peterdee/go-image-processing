package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"time"

	"go-image-processing/processing"
)

func save(name string, grid [][]color.Color) {
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
	jpeg.Encode(newFile, img.SubImage(img.Rect), nil)
}

func open(path string) [][]color.Color {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open the file")
	}
	defer file.Close()
	content, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("Could not decode the file", err)
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
			// c := rgba.At(i, j)
			// cc := content.At(i, j)
			// fmt.Println(c)
		}
		grid = append(grid, y)
	}
	return grid
}

func main() {
	img := open("images/7-x640.jpeg")
	now := time.Now().UnixMilli()
	// flipped := processing.FlipVertical(img)
	// gray := processing.Grayscale(img)
	// grayLum := processing.GrayscaleLuminocity(img)
	boxBlur := processing.BoxBlur(img, 7)
	// binary := processing.Binary(img, 185)
	// inverted := processing.Invert(img)
	est := time.Now().UnixMilli() - now
	println(est)
	name := fmt.Sprintf(`file-%d.jpeg`, time.Now().Unix())
	// save("gray-"+name, gray)
	// save("flipped-"+name, flipped)
	// save("binary-"+name, binary)
	// save("gray-lum-"+name, grayLum)
	save("box-blur-"+name, boxBlur)
	// save("inverted-"+name, inverted)
}
