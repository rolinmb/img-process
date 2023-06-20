package main

import (
	"fmt"
    "image/png"
    "log"
    "reflect"
    "os"
)

func main(){
	pngFile, err := os.Open("bg1i.png")
	if err != nil{
		log.Fatal(err)
	}
    defer pngFile.Close()
    decoded, err := png.Decode(pngFile)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println("Decoded .png file bytes:\n")
    fmt.Println(decoded)
    fmt.Println("Type of decoded:",reflect.TypeOf(decoded))
}