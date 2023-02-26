## go-image-processing

Various image processing functions written in Golang

This module is a base for https://github.com/julyskies/brille

### Available filters

These filters are ready to be used and were optimized compared to previous implementations

- Binary **(optimized, uses WaitGroup)**
- Box blur (dynamic)
- Brightness **(optimized, uses WaitGroup)**
- Contrast **(optimized, uses WaitGroup)**
- Eight colors (color reduction filter) **(optimized, uses WaitGroup)**
- Emboss filter (edge detection, static)
- Flip image (horizontal, vertical)
- Gamma correction **(optimized, uses WaitGroup)**
- Gaussian blur (dynamic) **(optimized, uses WaitGroup)**
- Grayscale (average, luminance) **(optimized, uses WaitGroup)**
- Hue rotate **(optimized, uses WaitGroup)**
- Inversion **(optimized, uses WaitGroup)**
- Kuwahara filter (edge detection / coloring, dynamic)
- Laplacian filter (edge detection, static) **(optimized, uses WaitGroup)**
- Rotate image by fixed angle (90 / 180 / 270 degrees)
- Sepia **(optimized, uses WaitGroup)**
- Sharpen filter (dynamic)
- Sobel filter (edge detection, static) **(optimized, uses WaitGroup)**
- Solarize **(optimized, uses WaitGroup)**

### In progress

These filters are not ready yet

- Bilateral filter (static / dynamic)
- Rotate image by any given angle

### License

[MIT](./LICENSE.md)
