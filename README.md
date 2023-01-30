## go-image-processing

Various image processing functions written in Golang

This module is a base for https://github.com/julyskies/brille

### Filters to add

1. Gaussian blur (static)
2. Image rotation for any given angle
3. Bilateral filter (static / dynamic)

### Optimized filters

These filters are available in [/optimized](/optimized) directory

Optimization can be applied to simple color filters

Speed comparison for `Solarize` filter (JPEG, 640x640):

```text
// optimized (values in ms)
open 0 convert 14 process 1 save 17 sum 32

// regular (values in ms)
open 0 convert 33 process 12 save 42 sum 87
```

Optimized filters can perform up to 10 times faster

### Added filters (non-optimized)

These filters are available in [/processing](/processing) directory

- Binary
- Box blur (dynamic)
- Brightness
- Contrast
- Eight colors (color reduction filter)
- Emboss filter (edge detection, static)
- Flip horizontal
- Flip vertical
- Gamma correction
- Grayscale (average)
- Grayscale (luminocity)
- Hue rotate
- Inversion
- Kuwahara filter (edge detection, dynamic)
- Laplasian filter (edge detection, static)
- Rotate image (90 degrees)
- Rotate image (180 degrees)
- Rotate image (270 degrees)
- Rotate image (any given angle) - **in progress**
- Sepia
- Sharpen
- Sobel filter (edge detection, static)
- Solarize

### License

[MIT](./LICENSE.md)
