## go-image-processing

Various image processing functions written in Golang

This module is a base for https://github.com/julyskies/brille

### Filters to add

1. Bilateral filter (static / dynamic)
2. Gaussian blur (static)
3. Image rotation for any given angle

### Available filters

These filters are ready to be used and were optimized compared to previous implementations

- Binary
- Box blur (dynamic)
- Brightness
- Contrast
- Eight colors (color reduction filter)
- Emboss filter (edge detection, static)
- Flip image (horizontal, vertical)
- Gamma correction
- Grayscale (average, luminance)
- Hue rotate
- Inversion
- Kuwahara filter (edge detection / coloring, dynamic)
- Laplacian filter (edge detection, static)
- Rotate image by fixed angle (90 / 180 / 270 degrees)
- Sepia
- Sharpen filter (dynamic)
- Sobel filter (edge detection, static)
- Solarize

### In progress

These filters are not ready yet

- Bilateral filter (static / dynamic)
- Gaussian blur (dynamic)
- Rotate image by any given angle

### License

[MIT](./LICENSE.md)
