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
    scale = 0.001
    geometrySize = 50
    complexity  = 10
	colorFactor = 40
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

func savePng(fname string, newPng *image.RGBA) {
    out,err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
    err = png.Encode(out, newPng)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nSuccessfully created/rewritten",fname)
    out.Close()
}

func main(){
    rand.Seed(time.Now().UnixNano())
	// var fname := "cramer.png"
	// cmdPrint(fname)
    // dest1 := "new_pngs/trippy33.png"
    // destn := "new_pngs/noisy16.png"
    // dest2 := "new_pngs/trippy_v2_10.png"
    // dest3 := "new_pngs/trippy_v3_10.png"
    dest4 := "new_pngs/trippy_v4_41.png"
    // dest5 := "new_pngs/trippy_v5_12.png"
    // trippyPng(dest1, width, height)
    // noisePng(destn, width, height)
    // trippyPng2(dest2, width, height)
    // trippyPng3(dest3, width, height)
    trippyPng4(dest4, width, height)
    // trippyPng5(dest5, width, height)
}

func trippyPng(fname string, width int, height int) {
    newPng := image.NewRGBA(image.Rect(0,0,width,height))
    for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8((x+(y*y))*((x*x)+(y*y)) % 256)
			g := uint8((x-(x*y))*((x*x)-(y*y)) % 256)
			b := uint8(((x*y*y)+(x*x*y))*((x*x*x)+(y*y*y)) % 256)
			a := uint8(255)
			newPng.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}
    savePng(fname, newPng)
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

func trippyPng2(fname string, width int, height int) {
    newPng := image.NewRGBA(image.Rect(0, 0, width, height))
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            r, g, b := calcColor(x,y)
            newPng.Set(x,y,color.RGBA{r, g, b, 255})
        }
    }
    savePng(fname, newPng)
}

func calcColor(x, y int) (uint8, uint8, uint8) {
    floatX := float64(x)
    floatY := float64(y)
	r := uint8((math.Sin((floatX*floatY*floatY)*scale) + 1) * 0.5 * 255)
	g := uint8((math.Sin((floatX*floatX*floatY)*scale) + 1) * 0.5 * 255)
	b := uint8((math.Sin((math.Exp(floatX+floatY))*scale) + 1) * 0.5 * 255)
	return r, g, b
}

func trippyPng3(fname string, width int, height int) {
    newPng := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b := calcColor2(x, y)
			newPng.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
    savePng(fname, newPng)
}

func calcColor2(x, y int) (uint8, uint8, uint8) {
    floatX := float64(x)
    floatY := float64(y)
	distance := math.Sqrt(math.Pow(float64(x-width/2), 2) + math.Pow(float64(y-height/2), 2))
	angle := math.Atan2(float64(y-height/2), float64(x-width/2))
	geometryValue := math.Abs(math.Sin(distance*scale*geometrySize) * math.Cos(angle))
	r := uint8((math.Sin(math.Pow(floatX, floatY)-geometryValue*math.Pi) + 1) * 0.5 * 255)
	g := uint8((math.Cos(math.Pow(floatY, floatX)-geometryValue*math.Pi) + 1) * 0.5 * 255)
	b := uint8((math.Sin(math.Pow(floatX+floatY, floatX-floatY)-geometryValue*math.Pi*2) + 1) * 0.5 * 255)
	return r, g, b
}


func trippyPng4(fname string, width int, height int) {
	newPng := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
            dx, dy := applyDistortion(x, y)
			r, g, b := calcColor3(dx, dy)
			newPng.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
    savePng(fname, newPng)
}

func applyDistortion(x, y int) (int, int) {
	// Adjust the distortion parameters to control the strength and effect of the distortion
	amplitude := 50.0
	frequency := 0.02
	phase := 0.0
	// Apply turbulence distortion
	dx := x + int(amplitude*math.Sin(frequency*float64(y)+phase))
	dy := y + int(amplitude*math.Sin(frequency*float64(x)+phase))
	// Ensure the distorted coordinates are within the image bounds
	dx = clamp(dx, 0, width-1)
	dy = clamp(dy, 0, height-1)
	return dx, dy
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func calcColor3(x, y int) (uint8, uint8, uint8) {
	angle := math.Pi * scale * float64((x*x*x*x*x*y)-(y*y*y*y*y*x))
	distance := math.Sqrt(math.Pow(float64(x-width/2), 2) + math.Pow(float64(y-height/2), 2))
	frequency := distance * scale
	r := uint8(math.Sin(angle*complexity+frequency)*colorFactor + 128)
	g := uint8(math.Sin(angle*complexity+frequency+2*math.Pi/3)*colorFactor + 128)
	b := uint8(math.Sin(angle*complexity+frequency+4*math.Pi/3)*colorFactor + 128)
	return r, g, b
}

func trippyPng5(fname string, width int, height int) {
	newPng := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b := calcColor4(x, y)
			newPng.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	savePng(fname, newPng)
}

func calcColor4(x, y int) (uint8, uint8, uint8) {
	amplitude := 100.0
	frequency := 0.01
	phase := 0.0
	r, g, b := 0.0, 0.0, 0.0
	// Add multiple layers of fractal patterns
	for i := 0; i < 5; i++ {
		dx := float64(x)
		dy := float64(y)
		layerAmplitude := amplitude / math.Pow(2, float64(i))
		layerFrequency := frequency * math.Pow(2, float64(i))
		layerPhase := phase * math.Pow(2, float64(i))
		for j := 0; j < 5; j++ {
			dx += layerAmplitude * math.Sin(layerFrequency*dy+layerPhase)
			dy += layerAmplitude * math.Sin(layerFrequency*dx+layerPhase)
		}
		r += math.Sin(layerFrequency*(dx-dy)+layerPhase)
		g += math.Sin(layerFrequency*(dx+dy)+layerPhase)
		b += math.Sin(layerFrequency*(dx*dy)+layerPhase)
	}
	r = 0.5 + (r/5)*0.5
	g = 0.5 + (g/5)*0.5
	b = 0.5 + (b/5)*0.5
	return uint8(r*255), uint8(g*255), uint8(b*255)
}