package main

import (
	"fmt"
    "log"
    "math"
    "math/rand"
    "image"
    "image/color"
    "image/png"
    "reflect"
    "time"
    "os"
)

func cmdPrint(fname string) {
	pngFile, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer pngFile.Close()
	decoded, err := png.Decode(pngFile)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Number of bytes in .png file: "+string(decoded.Bounds().Max.X)+"x"+string(decoded.Bounds().Max.Y))
    // fmt.Println(decoded)
    fmt.Println("\nType of pngFile =",reflect.TypeOf(pngFile))
    fmt.Println("Type of decoded =",reflect.TypeOf(decoded))
    fmt.Println()
    levels := []string{".", "o", "Æ", "&","@"} // characters to map to pixel brightness/darkness
    for y:= decoded.Bounds().Min.Y; y < decoded.Bounds().Max.Y; y++ {
        for x:= decoded.Bounds().Min.X; x < decoded.Bounds().Max.X; x++{
            c := color.GrayModel.Convert(decoded.At(x,y)).(color.Gray)
            level := c.Y / 51
            if level == 5{
                level--
            }
            fmt.Print(levels[level])
        }
        fmt.Print("\n")
    }
	pngFile.Close()
}

func main(){
    rand.Seed(time.Now().UnixNano())
	// var fname := "cramer.png"
	// cmdPrint(fname)
    dest1 := "new_pngs/trippy17.png"
    dest2 := "new_pngs/noisy5.png"
	w := 1000
	h := 1000
    newTrippyPng := image.NewRGBA(image.Rect(0,0,w,h))
    trippyPng(newTrippyPng, w, h)
    out1,err := os.Create(dest1)
	if err != nil {
		log.Fatal(err)
	}
    err = png.Encode(out1, newTrippyPng)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nSuccessfully created/rewritten",dest1)
    out1.Close()
    newNoisyPng := image.NewRGBA(image.Rect(0,0,w,h))
    noisePng(newNoisyPng, w, h)
    out2,err := os.Create(dest2)
	if err != nil {
		log.Fatal(err)
	}
    err = png.Encode(out2, newNoisyPng)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nSuccessfully created/rewritten",dest2)
    out2.Close()
}

func trippyPng(newPng *image.RGBA, width int, height int) {
    for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8((x - y) % 256)
			g := uint8((x*x*y + y*y*x) % 256)
			b := uint8((x*y*x + y*x*y) % 256)
			a := uint8(255)
			newPng.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}
}

func noisePng(newPng *image.RGBA, width int, height int) {
    for i := 0; i < width; i++ {
        for j := 0; j < height; j++ {
            r := uint8(math.Sin(float64(j+(1/(j+1)))/float64(rand.Intn(50))) * 127 + 128)
			g := uint8(math.Cos(float64(i*j)/float64(rand.Intn(50))) * 127 + 128)
			b := uint8(math.Sin(float64(i+j)/float64(rand.Intn(100))) * 127 + 128)
			a := uint8(255)
            newPng.SetRGBA(i,j,color.RGBA{r, g, b, a})
            rVariation := uint8(rand.Intn(50))
            gVariation := uint8(rand.Intn(50))
            bVariation := uint8(rand.Intn(50))
            r += rVariation
            g += gVariation
            b += bVariation
        }
    }
}