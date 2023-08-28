package main

import (
  "fmt"
  "image"
  "image/color"
  "image/png"
  "log"
  "os"
  "math"
)

const (
  width = 1600
  height = 1600
  fout = "new_pngs/lyapunov.png"
)

func main() {
  newPng := image.NewRGBA(image.Rect(0, 0, width, height))
  a := 1.5 // Change this value to see different patterns
  for px := 0; px < width; px++ {
	for py := 0; py < height; py++ {
	  x := float64(px)*2.0/width - 1.0
	  y := float64(py)*2.0/height - 1.0
	  lyapunovVal := lyapunov(x, y, a, 100) // Adjust iterations
	  color := getColor(lyapunovVal)
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
  fmt.Println("Lyapunov fractal image saved as", fout)
}

func lyapunov(x, y, a float64, iterations int) int {
  const rMin = 1.0
  const rMax = 4.0
  const rStep = 0.001
  var sum float64
  for r:= rMin; r <= rMax; r += rStep {
    lx, ly := x, y
    for i := 0; i < iterations; i++ {
      lx = r * lx * (1 - lx)
      ly = r * ly * (1 - ly)
      if i > 100 {
        sum += lnAbs(r * (1 - 2 * lx))
      }
    }
  }
  avg := sum / ((rMax - rMin) / rStep)
  if avg > 0 {
    return int(avg * 500) // adjust scaling to control detail
  }
  return 0
}

func lnAbs(x float64) float64 {
  if x < 0 {
    return -lnAbs(-x)
  }
  return math.Log(x)
}

func getColor(value int) color.Color {
  if value == 0 {
    return color.Black
  }
  return color.RGBA{uint8(value % 255), uint8(value % 255), uint8(value % 255), 255}
}