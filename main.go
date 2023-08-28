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

const (
    width = 1500
    height = 1500
)

func savePng(fname string, newPng *image.RGBA) {
    out,err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
    err = png.Encode(out, newPng)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nSuccessfully created/rewritten",fname)
}

func main(){
    rand.Seed(time.Now().UnixNano())
	cmdPrint("new_pngs/newton_fractal.png")
	noisePng("new_pngs/noisy08272023.png", width, height)
    trippyPng("new_pngs/trippy08272023.png", width, height)
}

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

func noisePng(fname string, width int, height int) {
    newPng := image.NewRGBA(image.Rect(0,0,width,height))
    for i := 0; i < width; i++ {
        for j := 0; j < height; j++ {
            r := uint8(math.Sin((float64(i-j)+float64(i+j))/float64(rand.Intn(50))) * 127 + 128)
			g := uint8(math.Cos((float64(i-j)-float64(i-j))/float64(rand.Intn(50))) * 127 + 128)
			b := uint8(math.Sin((float64(i+j)+float64(i+j))/float64(rand.Intn(100))) * 127 + 128)
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
    savePng(fname, newPng)
}

func trippyPng(fname string, width int, height int) {
    newPng := image.NewRGBA(image.Rect(0,0,width,height))
    for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8((x+(y*y))*((x)+(y)) % 256) // you can adjust these functions/use other functions
			g := uint8(((x*y))*((x)-(y*y)) % 256)
			b := uint8(((y)+(x))*((x)-(y)) % 256)
			a := uint8(255)
			newPng.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}
    savePng(fname, newPng)
}