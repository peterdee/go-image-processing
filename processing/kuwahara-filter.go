package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func getAperture(axisValue, axisMax, apertureMin, apertureMax int) (int, int) {
	start, end := 0, axisMax
	if axisValue+apertureMin > 0 {
		start = axisValue + apertureMin
	}
	if axisValue+apertureMax < axisMax {
		end = axisValue + apertureMax
	}
	return start, end
}

func KuwaharaFilter(source [][]color.Color, radius uint) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)

	radiusInt := int(radius)
	halfRadius := radiusInt / 2

	ApetureMinX := [4]int{-halfRadius, 0, -halfRadius, 0}
	ApetureMaxX := [4]int{0, halfRadius, 0, halfRadius}
	ApetureMinY := [4]int{-halfRadius, -halfRadius, 0, 0}
	ApetureMaxY := [4]int{0, 0, halfRadius, halfRadius}

	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {

			NumPixels := [4]int{0, 0, 0, 0}
			RValues := [4]int{0, 0, 0, 0}
			GValues := [4]int{0, 0, 0, 0}
			BValues := [4]int{0, 0, 0, 0}
			MaxRValue := [4]int{0, 0, 0, 0}
			MaxGValue := [4]int{0, 0, 0, 0}
			MaxBValue := [4]int{0, 0, 0, 0}
			MinRValue := [4]int{255, 255, 255, 255}
			MinGValue := [4]int{255, 255, 255, 255}
			MinBValue := [4]int{255, 255, 255, 255}

			for i := 0; i < 4; i += 1 {
				x2s, x2e := getAperture(x, width, ApetureMinX[i], ApetureMaxX[i])
				y2s, y2e := getAperture(y, height, ApetureMinY[i], ApetureMaxY[i])
				for x2 := x2s; x2 < x2e; x2 += 1 {
					for y2 := y2s; y2 < y2e; y2 += 1 {
						r, g, b, _ := utilities.RGBA(source[x2][y2])
						RValues[i] += int(r)
						GValues[i] += int(g)
						BValues[i] += int(b)

						if int(r) > MaxRValue[i] {
							MaxRValue[i] = int(r)
						} else if int(r) < MinRValue[i] {
							MinRValue[i] = int(r)
						}

						if int(g) > MaxGValue[i] {
							MaxGValue[i] = int(g)
						} else if int(g) < MinGValue[i] {
							MinGValue[i] = int(g)
						}

						if int(b) > MaxBValue[i] {
							MaxBValue[i] = int(b)
						} else if int(b) < MinBValue[i] {
							MinBValue[i] = int(b)
						}
						NumPixels[i] += 1
					}
				}
			}

			j := 0
			MinDifference := 10000
			for i := 0; i < 4; i += 1 {
				cdR := MaxRValue[i] - MinRValue[i]
				cdG := MaxGValue[i] - MinGValue[i]
				cdB := MaxBValue[i] - MinBValue[i]
				CurrentDifference := cdR + cdG + cdB
				if CurrentDifference < MinDifference && NumPixels[i] > 0 {
					j = i
					MinDifference = CurrentDifference
				}
			}

			cR := uint8(RValues[j] / NumPixels[j])
			cG := uint8(GValues[j] / NumPixels[j])
			cB := uint8(BValues[j] / NumPixels[j])
			destination[x][y] = color.RGBA{cR, cG, cB, 255}
		}
	}
	return destination
}
