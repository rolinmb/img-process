package main

import (
  "fmt"
  "image"
  "image/color"
  "image/png"
  "log"
  "os"
)

const (
  width = 6000
  height = 6000
  fout = "new_pngs/mandelbox.png"
)

func main() {
  newPng := image.NewRGBA(image.Rect(0, 0, width, height))
  centerX, centerY := 0.0, 0.0
  scaleX, scaleY := 4.0 / float64(width), 4.0 / float64(height)
  for px := 0; px < width; px++ {
    for py := 0; py < height; py++ {
      x := float64(px)*scaleX - 2.0 + centerX
      y := float64(py)*scaleY - 2.0 + centerY
      mandelboxVal := mandelbox(x, y, 8) // adjust power
      color := getColor(mandelboxVal)
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
	fmt.Println("Mandelbox fractal image saved as", fout)
}

func mandelbox(x, y float64, iterations int) int {
  const bailout = 2.0
  const minR = -1.0
  const maxR = 1.0
  const minI = -1.0
  const maxI = 1.0
  const scale = 1.5
  r, i := x, y
  for iter := 0; iter < iterations; iter++ {
    r2, i2 := r*r, i*i
    if r2+i2 > bailout*bailout {
      return iter
    }
    theta := scale * r
    phi := scale * i
    r = r2 - i2 + theta
    i = phi + phi*r2 - phi*i2 + y
    if r < minR {
      r = 2*minR - r
    } else if r > maxR {
      r = 2*maxR - r
    }
    if i < minI {
      i = 2*minI - i
    } else if i > maxI {
      i = 2*maxI - i
    }
  }
  return iterations
}

func getColor(iterations int) color.Color {
	if iterations == 8 {
		return color.Black
	}
	return color.RGBA{uint8(iterations % 255), uint8(iterations % 255), uint8(iterations % 255), 255}
}