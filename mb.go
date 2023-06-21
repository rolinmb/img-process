package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	width  = 5000
	height = 5000
    fout = "new_pngs/mb.png"
)

func main() {
	newPng := image.NewRGBA(image.Rect(0, 0, width, height))
	// Define the boundaries and scaling factors for the Mandelbrot Set
	xmin, xmax := -2.0, 1.0
	ymin, ymax := -1.5, 1.5
	xscale := (xmax - xmin) / float64(width)
	yscale := (ymax - ymin) / float64(height)
	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			x := float64(px)*xscale + xmin
			y := float64(py)*yscale + ymin
			mandelbrotValue := mandelbrot(x, y)
			// Map the Mandelbrot value to a color
			color := getColor(mandelbrotValue)
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
	log.Println("Mandelbrot Set image saved as",fout)
}

func mandelbrot(x, y float64) int {
	const maxIterations = 1000
	const escapeRadius = 2
	var zx, zy, zx2, zy2 float64
	for i := 0; i < maxIterations; i++ {
		if zx2+zy2 > escapeRadius*escapeRadius {
			return i
		}
		zy = 2*zx*zy + y
		zx = zx2 - zy2 + x
		zx2 = zx * zx
		zy2 = zy * zy
	}
	return maxIterations
}

func getColor(iterations int) color.Color {
	// Assign a color based on the number of iterations
	// You can modify this function to customize the color mapping
	// based on your preference
	if iterations == 1000 {
		return color.Black
	}
	return color.RGBA{uint8(iterations % 255), 0, 0, 255}
}