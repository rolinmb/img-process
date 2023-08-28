package main

import (
  "image"
  "image/color"
  "image/png"
  "log"
  "os"
  "math"
)

const (
  width  = 5000
  height = 5000
  fout = "new_pngs/multibrot.png"
)

func main() {
  newPng := image.NewRGBA(image.Rect(0, 0, width, height))
  xmin, xmax := -2.0, 1.0
  ymin, ymax := -1.5, 1.5
  xscale := (xmax - xmin) / float64(width)
  yscale := (ymax - ymin) / float64(height)
  for px := 0; px < width; px++ {
    for py := 0; py < height; py++ {
      x := float64(px)*xscale + xmin
      y := float64(py)*yscale + ymin
      mandelbulbVal := multibrot(x, y, 5) // adjust power
      color := getColor(mandelbulbVal)
      newPng.Set(px, py, color)
    }
  }
  newFile, err := os.Create(fout)
  if err != nil {
    log.Fatal(err)
  }
  defer newFile.Close()
  err = png.Encode(newFile, newPng)
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Multibrot Set saved at", fout)
}

func multibrot(x, y float64, power int) int {
  const maxIters = 1000
  const escapeRadius = 2
  var zx, zy, zx2, zy2 float64
  for i := 0; i < maxIters; i++ {
    if zx2+zy2 > escapeRadius*escapeRadius {
      return i
    }
    zy = math.Pow(math.Abs(zx), float64(power)) + y
    zx = zx2 - zy2 + x
    zx2 = zx * zx
    zy2 = zy * zy
  }
  return maxIters
}

func getColor(iterations int) color.Color {
  if iterations == 1000 {
    return color.Black
  }
  return color.RGBA{uint8(iterations % 255), uint8(iterations % 255), uint8(iterations % 255), 255}
}