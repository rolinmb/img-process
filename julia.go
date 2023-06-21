package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

const (
	width  = 800
	height = 800
	xmin, xmax = -2, 2
	ymin, ymax = -2, 2
	maxIterations = 200
    fout = "demo_results/julia_set.png"
)

func main() {
	newPng := image.NewRGBA(image.Rect(0, 0, width, height))
    // Julia Set Constant
    const c = complex(-0.8, 0.156)
	// Iterate over each pixel
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			iterations := calculateIterations(z, c)
			// Map the number of iterations to a color
			color := getColor(iterations)
			// Set the pixel color in the image
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
	fmt.Println("Julia Set image saved as",fout)
}

func calculateIterations(z complex128, c complex128) int {
	iterations := 0
	for iterations < maxIterations {
		if cmplx.Abs(z) > 2 {
			break
		}
		z = z*z + c
		iterations++
	}
	return iterations
}

func getColor(iterations int) color.Color {
	// Assign a color based on the number of iterations
	// You can modify this function to customize the color mapping
	// based on your preference
	if iterations == maxIterations {
		return color.Black
	}
	red := uint8(iterations % 16 * 16)
	green := uint8(iterations % 32 * 8)
	blue := uint8(iterations % 64 * 4)
	return color.RGBA{red, green, blue, 255}
}