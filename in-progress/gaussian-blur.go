package progress

import (
	"image/color"
	"math"

	"go-image-processing/utilities"
)

const K float64 = 6

func createKernel(radius uint) []float64 {
	dim := math.Max(3.0, K*float64(radius))
	sqrtSigmaPi2 := math.Sqrt(math.Pi*2.0) * float64(radius)
	s2 := float64(2.0 * radius * radius)
	sum := 0.0
	kDim := dim
	if int(kDim)%2 == 0 {
		kDim = dim - 1
	}
	kernel := make([]float64, int(kDim))
	half := int(len(kernel) / 2)
	i := -half
	for j := 0; j < len(kernel); j += 1 {
		kernel[j] = math.Exp(-float64(i*i)/(s2)) / sqrtSigmaPi2
		sum += kernel[j]
		i += 1
	}
	for k := 0; k < int(kDim); k += 1 {
		kernel[k] /= sum
	}
	return kernel
}

// func gauss(pixels, kernel, ch, gray) {
// var data = pixels.data;
// var w = pixels.width;
// var h = pixels.height;
// var buff = new Uint8Array(w*h);
// var mk = Math.floor(kernel.length / 2);
// kl := len(kernel.length;

// // First step process columns
// for (var j = 0, hw = 0; j < h; j++, hw += w)
// {
//   for (var i = 0; i < w; i++)
//   {
//     var sum = 0;
//     for (var k = 0; k < kl; k++)
//     {
//       var col = i + (k - mk);
//       col = (col < 0) ? 0 : ((col >= w) ? w - 1 : col);
//       sum += data[(hw + col)*4 + ch]*kernel[k];
//     }
//     buff[hw + i] = sum;
//   }
// }

// Second step process rows
// for (var j = 0, offset = 0; j < h; j++, offset += w)
// {
//   for (var i = 0; i < w; i++)
//   {
//     var sum = 0;
//     for (k = 0; k < kl; k++)
//     {
//       var row = j + (k - mk);
//       row = (row < 0) ? 0 : ((row >= h) ? h - 1 : row);
//       sum += buff[(row*w + i)]*kernel[k];
//     }
//     var off = (j*w + i)*4;
//     (!gray) ? data[off + ch] = sum :
//               data[off] = data[off + 1] = data[off + 2] = sum;
//   }
// }
// }

func clampVal(val, kH, max int, kernel []float64) ([]float64, int, int) {
	kLen := len(kernel)
	start, end := 0, kLen
	if val-kH < 0 {
		start = kH - val
	}
	if val+kH > max {
		end = kLen - (val + kH - max)
	}
	gs, ge := 0, max
	if val-kH > 0 {
		gs = val - kH
	}
	if val+kH < max {
		ge = val + kH
	}
	return kernel[start:end], gs, ge
}

func GaussianBlur(source [][]color.Color, radius uint) [][]color.Color {
	if radius == 0 {
		radius = 1
	}
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)

	kernel := createKernel(radius)
	kernelLen := len(kernel)
	mk := int(math.Floor(float64(kernelLen) / 2))

	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			sumR := 0.0
			sumG := 0.0
			sumB := 0.0

			kX, gxs, gxe := clampVal(x, mk, width, kernel)
			kY, gys, gye := clampVal(y, mk, height, kernel)
			i := 0
			for m := gxs; m < gxe; m += 1 {
				j := 0
				for n := gys; n < gye; n += 1 {
					r, g, b, _ := utilities.RGBA(source[m][n])
					sumR += (float64(r)*kX[i] + float64(r)*kY[j])
					sumG += (float64(g)*kX[i] + float64(g)*kY[j])
					sumB += (float64(b)*kX[i] + float64(b)*kY[j])
					j += 1
				}
				i += 1
			}

			destination[x][y] = color.RGBA{
				uint8(math.Sqrt(sumR)),
				uint8(math.Sqrt(sumG)),
				uint8(math.Sqrt(sumB)),
				255,
			}
		}
	}
	return destination
}
