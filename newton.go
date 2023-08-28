package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

const (
	width = 4000
	height = 4000
	xmin, xmax = -2, 2
	ymin, ymax = -2, 2
	epsilon = 1e-6
	maxIterations = 200
    fout = "new_pngs/newton_fractal.png"
)

func main() {
	newPng := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			iterations := calculateIterations(z)
			color := getColor(iterations)
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
	fmt.Println("Newton Fractal image saved as",fout)
}

func calculateIterations(z complex128) int {
	iterations := 0
	for iterations < maxIterations {
		z = z - (z*z*z-1)/(3*z*z)
		if cmplx.Abs(z*z*z-1) < epsilon {
			break
		}
		iterations++
	}
	return iterations
}

func getColor(iterations int) color.Color {
	if iterations == maxIterations {
		return color.Black
	}
	hue := float64(iterations%50) / 50.0
	r, g, b := hsvToRGB(hue, 1.0, 1.0)
	r = uint8(r * 255)
	g = uint8(g * 255)
	b = uint8(b * 255)
	return color.RGBA{r, g, b, 255}
}
// hsvToRGB converts an HSV color to RGB color.
// Hue: 0.0 - 1.0
// Saturation: 0.0 - 1.0
// Value: 0.0 - 1.0
func hsvToRGB(hue, saturation, value float64) (uint8, uint8, uint8) {
	hue *= 6.0
	i := math.Floor(hue)
	f := hue - i
	p := value * (1.0 - saturation)
	q := value * (1.0 - f*saturation)
	t := value * (1.0 - (1.0-f)*saturation)

	switch i {
	case 0:
		return uint8(value * 255), uint8(t * 255), uint8(p * 255)
	case 1:
		return uint8(q * 255), uint8(value * 255), uint8(p * 255)
	case 2:
		return uint8(p * 255), uint8(value * 255), uint8(t * 255)
	case 3:
		return uint8(p * 255), uint8(q * 255), uint8(value * 255)
	case 4:
		return uint8(t * 255), uint8(p * 255), uint8(value * 255)
	case 5:
		return uint8(value * 255), uint8(p * 255), uint8(q * 255)
	default:
		return uint8(value * 255), uint8(p * 255), uint8(q * 255)
	}
}