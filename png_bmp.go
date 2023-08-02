package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"golang.org/x/image/bmp"
)

func main() {
	// Input and Output file paths
	pngFilePath := "new_pngs/julia_set.png"
	bmpFilePath := "button_bg.bmp"

	// Read the PNG image from the file
	pngFile, err := os.Open(pngFilePath)
	if err != nil {
		fmt.Println("Error opening the PNG file:", err)
		return
	}
	defer pngFile.Close()

	pngImage, err := png.Decode(pngFile)
	if err != nil {
		fmt.Println("Error decoding the PNG image:", err)
		return
	}

	// Create the BMP file to write the image
	bmpFile, err := os.Create(bmpFilePath)
	if err != nil {
		fmt.Println("Error creating the BMP file:", err)
		return
	}
	defer bmpFile.Close()

	// Write the PNG image to the BMP file
	err = bmp.Encode(bmpFile, pngImage)
	if err != nil {
		fmt.Println("Error encoding the BMP image:", err)
		return
	}

	fmt.Println("PNG image converted to BMP successfully.")
}