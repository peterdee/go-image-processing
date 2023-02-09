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
	heightCorrection := 0
	if height%2 != 0 {
		heightCorrection = 1
	}
	if angle == constants.ROTATE_FIXED_180 {
		for i := 0; i < len(img.Pix); i += 4 {
			x, y := getCoordinates(i/4, width)
			if y < height/2+heightCorrection {
				j := getPixel(width-x-1, height-y-1, width)
				r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
				img.Pix[i], img.Pix[i+1], img.Pix[i+2] = img.Pix[j], img.Pix[j+1], img.Pix[j+2]
				if i != j {
					img.Pix[j], img.Pix[j+1], img.Pix[j+2] = r, g, b
				}
			}
		}
		processMS := int(math.Round(float64(time.Now().UnixNano())/1000000) - now)
		saveMS := save(img, format)
		sum := openMS + convertMS + processMS + saveMS
		println("open", openMS, "convert", convertMS, "process", processMS, "save", saveMS, "sum", sum)
	}
	if angle == constants.ROTATE_FIXED_90 || angle == constants.ROTATE_FIXED_270 {
		destination := utilities.CreateGrid(height, width)
		for i := 0; i < len(img.Pix); i += 4 {
			x, y := getCoordinates(i/4, width)
			if angle == constants.ROTATE_FIXED_90 {
				destination[height-y-1][x] = color.RGBA{
					img.Pix[i],
					img.Pix[i+1],
					img.Pix[i+2],
					img.Pix[i+3],
				}
			}
			if angle == constants.ROTATE_FIXED_270 {
				destination[y][width-x-1] = color.RGBA{
					img.Pix[i],
					img.Pix[i+1],
					img.Pix[i+2],
					img.Pix[i+3],
				}
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
}
