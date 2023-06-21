package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	Width  = 1000
	Height = 1000
	MaxIt  = 200
    fout = "demo_results/tchebichef_fractal.png"
)

func main() {
	newPng := image.NewRGBA(image.Rect(0, 0, Width, Height))
	for py := 0; py < Height; py++ {
		for px := 0; px < Width; px++ {
			// Convert pixel coordinates to fractal coordinates
			x := float64(px)/Width*3.5 - 2.5
			y := float64(py)/Height*3.5 - 2.5
			// Perform Tchebichef fractal calculation
			var sum float64
			for n := 0; n < MaxIt; n++ {
				sum += tchebichef(x, y, n)
			}
			// Normalize the sum
			sum /= MaxIt
			// Map the sum to the color
			colorVal := uint8(math.Floor(sum * 255))
			color := color.RGBA{colorVal, colorVal, colorVal, 255}
			newPng.Set(px, py, color)
		}
	}
	newFile, err := os.Create(fout)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	if err := png.Encode(newFile, newPng); err != nil {
		log.Fatal(err)
	}
	log.Println("Image generated and saved as",fout)
}

func tchebichef(x, y float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return y
	} else if n > 1 {
		return 2 * x * tchebichef(x, y, n-1) - tchebichef(x, y, n-2)
	} else {
        return 0
    }
}