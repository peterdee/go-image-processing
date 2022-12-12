package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"time"
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

func open() [][]color.Color {
	file, err := os.Open("images/1.jpeg")
	if err != nil {
		log.Fatal("Could not open the file")
	}
	defer file.Close()
	content, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("Could not decode the file")
	}
	var grid [][]color.Color
	size := content.Bounds().Size()
	for i := 0; i < size.X; i += 1 {
		var y []color.Color
		for j := 0; j < size.Y; j += 1 {
			y = append(y, content.At(i, j))
		}
		grid = append(grid, y)
	}
	return grid
}

func grayscale(grid [][]color.Color) [][]color.Color {
	container := make([][]color.Color, len(grid))
	for x := 0; x < len(container); x += 1 {
		col := grid[x]
		resultCol := make([]color.Color, len(col))
		for y := 0; y < len(resultCol); y += 1 {
			px := grid[x][y].(color.YCbCr)
			gray := uint8(float64(px.Cb)/3.0 + float64(px.Cr)/3.0 + float64(px.Y)/3.0)
			resultCol[y] = color.YCbCr{gray, gray, gray}
		}
		container[x] = resultCol
	}
	return container
}

func flipVertical(grid [][]color.Color) [][]color.Color {
	flipped := make([][]color.Color, len(grid))
	for x := 0; x < len(flipped); x += 1 {
		col := grid[x]
		container := make([]color.Color, len(col))
		for y := 0; y < len(col)/2; y += 1 {
			z := len(col) - y - 1
			container[y], container[z] = col[z], col[y]
		}
		flipped[x] = container
	}
	return flipped
}

func main() {
	img := open()
	// flipped := flipVertical(img)
	now := time.Now().UnixMilli()
	gray := grayscale(img)
	est := time.Now().UnixMilli() - now
	println(est)
	name := fmt.Sprintf(`file-%d.jpeg`, time.Now().Unix())
	save(name, gray)
}
