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
	horizontal := utilities.CreateGrid(width, height)
	destination := utilities.CreateGrid(width, height)

	k := createKernel(4)
	j := len(k)/2 + 25

	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			if x < j || x > width-j || y < j || y > height-j {
				horizontal[x][y] = source[x][y]
				continue
			}

			sumR := 0.0
			sumG := 0.0
			sumB := 0.0
			for m := 0; m < len(k); m += 1 {
				mx := x - (len(k)/2 - m)
				r, g, b, _ := utilities.RGBA(source[mx][y])
				sumR += float64(r) * float64(k[m])
				sumG += float64(g) * float64(k[m])
				sumB += float64(b) * float64(k[m])
			}
			horizontal[x][y] = color.RGBA{
				uint8(utilities.MaxMin(sumR, 255, 0)),
				uint8(utilities.MaxMin(sumG, 255, 0)),
				uint8(utilities.MaxMin(sumB, 255, 0)),
				255,
			}
		}
	}

	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			if x < j || x > width-j || y < j || y > height-j {
				destination[x][y] = horizontal[x][y]
				continue
			}
			sumR := 0.0
			sumG := 0.0
			sumB := 0.0
			for n := 0; n < len(k); n += 1 {
				nx := y - (len(k)/2 - n)
				r, g, b, _ := utilities.RGBA(horizontal[x][nx])
				sumR += float64(r) * float64(k[n])
				sumG += float64(g) * float64(k[n])
				sumB += float64(b) * float64(k[n])
			}
			destination[x][y] = color.RGBA{
				uint8(utilities.MaxMin(sumR, 255, 0)),
				uint8(utilities.MaxMin(sumG, 255, 0)),
				uint8(utilities.MaxMin(sumB, 255, 0)),
				255,
			}
		}
	}
	return destination
}
