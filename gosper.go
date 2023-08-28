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
  width = 6000
  height = 6000
  fout = "new_pngs/gosper.png"
)

func main() {
  newPng := image.NewRGBA(image.Rect(0, 0, width, height))
  x, y := 50.0, 700.0 // Starting position
  length := 600.0     // Length of the initial segment
  iterations := 4     // Number of iterations
  angle := 60.0       // Angle between segments in degrees
  drawGosperCurve(newPng, x, y, length, iterations, angle)
  newFile, err := os.Create(fout)
  if err != nil {
	log.Fatal(err)
  }
  defer newFile.Close()
  err = png.Encode(newFile, newPng)
  if err != nil {
	log.Fatal(err)
  }
  fmt.Println("Gosper curve image saved as", fout)
}

func drawGosperCurve(img *image.RGBA, x, y, length float64, iterations int, angle float64) {
  if iterations == 0 {
	drawSegment(img, x, y, length, angle)
  } else {
	length /= 2.0
	drawGosperCurve(img, x, y, length, iterations-1, angle)
	x += length * cos(angle)
	y += length * sin(angle)
	drawGosperCurve(img, x, y, length, iterations-1, -angle)
	x += length * cos(-angle)
	y += length * sin(-angle)
	angle -= 60.0
	drawGosperCurve(img, x, y, length, iterations-1, angle)
	x += length * cos(angle)
	y += length * sin(angle)
	drawGosperCurve(img, x, y, length, iterations-1, -angle)
	x += length * cos(-angle)
	y += length * sin(-angle)
	angle += 60.0
	drawGosperCurve(img, x, y, length, iterations-1, angle)
	x += length * cos(angle)
	y += length * sin(angle)
	drawGosperCurve(img, x, y, length, iterations-1, angle)
  }
}

func drawSegment(img *image.RGBA, x, y, length, angle float64) {
  x1, y1 := x+length*cos(angle), y+length*sin(angle)
  drawLine(img, x, y, x1, y1)
}

func drawLine(img *image.RGBA, x0, y0, x1, y1 float64) {
  dx, dy := x1-x0, y1-y0
  steps := max(abs(int(dx)), abs(int(dy)))
  xIncrement, yIncrement := dx/float64(steps), dy/float64(steps)
  x, y := x0, y0
  for i := 0; i < steps; i++ {
	img.Set(int(x+0.5), int(y+0.5), color.Black)
	x += xIncrement
	y += yIncrement
  }
}

func cos(degrees float64) float64 {
  return cosPi(degrees / 180.0)
}

func sin(degrees float64) float64 {
  return sinPi(degrees / 180.0)
}

func cosPi(radians float64) float64 {
  return cosApprox(radians * math.Pi)
}

func sinPi(radians float64) float64 {
  return sinApprox(radians * math.Pi)
}

func cosApprox(x float64) float64 {
  x = mod2Pi(x)
  sign := 1.0
  if x > math.Pi {
	x -= math.Pi
	sign = -1.0
  }
  x = (x * x * (3 - 2*x)) // Polynomial approximation
  return sign * x
}

func sinApprox(x float64) float64 {
  x = mod2Pi(x)
  sign := 1.0
  if x > math.Pi {
	x -= math.Pi
	sign = -1.0
  }
  x = (3.14159265*x - x*x*x) // Polynomial approximation
  return sign * x
}

func mod2Pi(x float64) float64 {
  x = x / (2.0 * math.Pi)
  x = 2.0*math.Pi*(x - math.Floor(x))
  if x < 0 {
	x += 2.0 * math.Pi
  }
  return x
}

func max(a, b int) int {
  if a > b {
	return a
  }
  return b
}

func abs(x int) int {
  if x < 0 {
	return -x
  }
  return x
}