## go-image-processing

Various image processing functions written in Golang

This module is a base for https://github.com/julyskies/brille

### Available filters

These filters are ready to be used and were optimized compared to previous implementations

- Binary **(optimized, uses WaitGroup)**
- Box blur (dynamic)
- Brightness
- Contrast
- Eight colors (color reduction filter)
- Emboss filter (edge detection, static)
- Flip image (horizontal, vertical)
- Gamma correction
- Gaussian blur (dynamic) **(optimized, uses WaitGroup)**
- Grayscale (average, luminance)
- Hue rotate
- Inversion
- Kuwahara filter (edge detection / coloring, dynamic)
- Laplacian filter (edge detection, static)
- Rotate image by fixed angle (90 / 180 / 270 degrees)
- Sepia **(optimized, uses WaitGroup)**
- Sharpen filter (dynamic)
- Sobel filter (edge detection, static)
- Solarize **(optimized, uses WaitGroup)**

### In progress

These filters are not ready yet

- Bilateral filter (static / dynamic)
- Rotate image by any given angle

### License

[MIT](./LICENSE.md)
