package filters

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"go-image-processing/constants"
	"go-image-processing/utilities"
)

func RotateFixed(path string, angle uint) {
	if angle != constants.ROTATE_FIXED_90 &&
		angle != constants.ROTATE_FIXED_180 &&
		angle != constants.ROTATE_FIXED_270 {
		angle = constants.ROTATE_FIXED_90
	}
	img, format, openMS, convertMS := open(path)
	now := math.Round(float64(time.Now().UnixNano()) / 1000000)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	gridWidth, gridHeight := width, height
	if angle == constants.ROTATE_FIXED_180 {
		gridWidth, gridHeight = height, width
	}
	destination := utilities.CreateGrid(gridWidth, gridHeight)
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := getCoordinates(i/4, width)
		dx, dy := x, y
		if angle == constants.ROTATE_FIXED_90 {
			dx, dy = height-y-1, x
		}
		if angle == constants.ROTATE_FIXED_180 {
			dx, dy = width-x-1, height-y-1
		}
		if angle == constants.ROTATE_FIXED_270 {
			dx, dy = y, width-x-1
		}
		destination[dx][dy] = color.RGBA{
			img.Pix[i],
			img.Pix[i+1],
			img.Pix[i+2],
			img.Pix[i+3],
		}
	}
	processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
	saveMS := utilities.SaveFile(
		fmt.Sprintf(`file-%d.%s`, time.Now().Unix(), format),
		format,
		destination,
	)
	sum := openMS + convertMS + processMS + saveMS
	println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
}
