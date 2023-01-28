package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func KuwaharaFilter(source [][]color.Color, radius uint) [][]color.Color {
	width, height := len(source), len(source[0])

	radiusInt := int(radius)
	halfRadius := radiusInt / 2

	ApetureMinX := [4]int{-halfRadius, 0, -halfRadius, 0}
	ApetureMaxX := [4]int{0, halfRadius, 0, halfRadius}
	ApetureMinY := [4]int{-halfRadius, -halfRadius, 0, 0}
	ApetureMaxY := [4]int{0, 0, halfRadius, halfRadius}

	for x := halfRadius; x < width-halfRadius; x += 1 {
		for y := halfRadius; y < height-halfRadius; y += 1 {

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
				for x2 := ApetureMinX[i]; x2 <= ApetureMaxX[i]; x2 += 1 {
					TempX := x + x2
					for y2 := ApetureMinY[i]; y2 <= ApetureMaxY[i]; y2 += 1 {
						TempY := y + y2

						r, g, b, _ := utilities.RGBA(source[TempX][TempY])
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
				CurrentDifference := (MaxRValue[i] - MinRValue[i]) + (MaxGValue[i] - MinGValue[i]) + (MaxBValue[i] - MinBValue[i])
				if CurrentDifference < MinDifference && NumPixels[i] > 0 {
					j = i
					MinDifference = CurrentDifference
				}
			}

			cR := uint8(RValues[j] / NumPixels[j])
			cG := uint8(GValues[j] / NumPixels[j])
			cB := uint8(BValues[j] / NumPixels[j])
			source[x][y] = color.RGBA{cR, cG, cB, 255}
		}
	}
	return source
}
