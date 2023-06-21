package main

import (
	"fmt"
    "log"
    // "image"
    "image/color"
    "image/png"
    "reflect"
    "os"
)

func cmdPrint(fname string){
	pngFile, err := os.Open(fname)
	if err != nil{
		log.Fatal(err)
	}
	defer pngFile.Close()
	decoded, err := png.Decode(pngFile)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println("Number of bytes in .png file: "+string(decoded.Bounds().Max.X)+"x"+string(decoded.Bounds().Max.Y))
    // fmt.Println(decoded)
    fmt.Println("\nType of pngFile =",reflect.TypeOf(pngFile))
    fmt.Println("Type of decoded =",reflect.TypeOf(decoded))
    fmt.Println()
    levels := []string{".", "o", "Æ", "&","@"} // characters to map to pixel brightness/darkness
    for y:= decoded.Bounds().Min.Y; y < decoded.Bounds().Max.Y; y++{
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
	var fname = "cramer.png"
	cmdPrint(fname)
    /*
    newPng := image.NewRGBA(image.Rect(0,0,12,6)) // Empty image matrix
    out,err := os.Create("new.png")
    png.Encode(out, newPng)
    fmt.Println("\nSuccessfully created/rewritten new.png.\n")
    out.Close()
    */
}