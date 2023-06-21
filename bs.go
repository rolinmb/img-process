package main

import (
	"image"
	"image/color"
	"image/png"
    "math"
	"log"
	"os"
)

const (
	Width  = 10000
	Height = 10000
	MaxIt  = 1000
    fout = "new_pngs/burning_ship.png"
)

func main() {
	newPng := image.NewRGBA(image.Rect(0, 0, Width, Height))
	for py := 0; py < Height; py++ {
		for px := 0; px < Width; px++ {
			// Convert pixel coordinates to complex numbers
			x0 := float64(px)/Width*3.5 - 2.5
			y0 := float64(py)/Height*3.5 - 2.5
			x := 0.0
			y := 0.0
			// Perform burning ship iteration
			var i int
            for i = 0; i < MaxIt; i++ {
				x2 := x * x
				y2 := y * y
				// Check for escape condition
				if x2+y2 > 4.0 {
					break
				}
				y = math.Abs(2.0*x*y) - y0
				x = math.Abs(x2-y2) - x0
			}
			// Color the pixel based on the number of iterations
			colorVal := uint8(i % 256)
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