package processing

import (
	"image/color"

	"go-image-processing/utilities"
)

func KuwaharaFilter(source [][]color.Color, radius uint) [][]color.Color {
	width, height := len(source), len(source[0])

	radiusInt := int(radius)

	ApetureMinX := [4]int{-radiusInt / 2, 0, -radiusInt / 2, 0}
	ApetureMaxX := [4]int{0, radiusInt / 2, 0, radiusInt / 2}
	ApetureMinY := [4]int{-radiusInt / 2, -radiusInt / 2, 0, 0}
	ApetureMaxY := [4]int{0, 0, radiusInt / 2, radiusInt / 2}

	for x := radiusInt; x < width-radiusInt; x += 1 {
		for y := radiusInt; y < height-radiusInt; y += 1 {

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
				for x2 := ApetureMinX[i]; x2 < ApetureMaxX[i]; x2 += 1 {
					TempX := x + x2
					for y2 := ApetureMinY[i]; y2 < ApetureMaxY[i]; y2 += 1 {
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
							MaxRValue[i] = int(r)
						} else if int(g) < MinGValue[i] {
							MinRValue[i] = int(g)
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
