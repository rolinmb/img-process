package main

import (
	"fmt"
    "log"
    "math"
    "math/rand"
    "image"
	"image/draw"
    "image/color"
    "image/png"
    "reflect"
    "time"
    "os"
)

const (
    width = 1500
    height = 1500
    geometrySize = 50  // trippyPng3()
	scale = 0.001      // trippyPng4()
    complexity  = 10   // trippyPng4()
	colorFactor = 40   // trippyPng4()
	NoiseScale = 0.01  // trippyPng6() & trippyPng7()
	NumLayers = 5	   // trippyPng6()
	MaxAlpha = 200	   // trippyPng6() & trippyPng7()
	MinAlpha = 100     // trippyPng6() & trippyPng7()
	NumShapes = 500    // trippyPng7()
	MaxShapeSize = 50  // trippyPng7()
	MinShapeSize = 10  // trippyPng7()
	ScaleFactor = 0.05 // trippyPng7()
	MaxDistortion = 30 // trippyPng7()
	DistortionScale = 0.1 // trippyPng7()
	MaxLineWidth = 5   // trippyPng7()
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
	// cmdPrint("odnodi.png")
	// noisePng("new_pngs/new_noisy.png", width, height)
    // trippyPng("new_pngs/trippy33.png", width, height)
    // trippyPng2("new_pngs/trippy_v2_10.png", width, height)
    // trippyPng3("new_pngs/trippy_v3_10.png", width, height)
    trippyPng4("new_pngs/trippy_v4_45pm.png", width, height)
    // trippyPng5("new_pngs/trippy_v5_12.png", width, height)
	// trippyPng6("new_pngs/trippy_v6_1.png", width, height)
	// trippyPng7("new_pngs/trippy_v7_1.png", width, height)
	// randomPng("new_pngs/new_rand.png", width, height)
	// trippyFx("new_pngs/trippy31.png", "new_pngs/trippy31_fx2.png")
	// interpolatePngs("new_pngs/trippy_v5_0.png", "new_pngs/trippy_v4_3.png", "new_pngs/trippy_interp_19.png", 0.00875)
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
			r := uint8((x+(y*y))*((x)+(y)) % 256)
			g := uint8(((x*y))*((x)-(y*y)) % 256)
			b := uint8(((y)+(x))*((x)-(y)) % 256)
			a := uint8(255)
			newPng.SetRGBA(x, y, color.RGBA{r, g, b, a})
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
	amplitude := 50.0
	frequency := 0.02
	phase := 0.0
	dx := x + int(amplitude*math.Sin(frequency*float64(x)+phase))
	dy := y + int(amplitude*math.Sin(frequency*float64(y)+phase))
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
	angle := math.Pi * scale * float64((x)-(y))
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

func trippyPng6(fname string, width int, height int) {
	newPng := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for layer := 0; layer < NumLayers; layer++ {
				offsetX := float64(x) * NoiseScale * float64(layer+1)
				offsetY := float64(y) * NoiseScale * float64(layer+1)
				noiseValue := noise(offsetX, offsetY)
				r := uint8(math.Abs(noiseValue) * 255)
				g := uint8(math.Abs(math.Sin(noiseValue)) * 255)
				b := uint8(math.Abs(math.Cos(noiseValue)) * 255)
				a := uint8(rand.Intn(MaxAlpha-MinAlpha) + MinAlpha)
				c := color.RGBA{r, g, b, a}
				newPng.Set(x, y, c)
			}
		}
	}
	savePng(fname, newPng)
}

func noise(x, y float64) float64 {
	return math.Sin(x*y) + math.Cos(x/y)
}

func trippyPng7(fname string, width int, height int) {
	newPng := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			x := rand.Intn(width)
			y := rand.Intn(height)
			size := rand.Intn(MaxShapeSize-MinShapeSize) + MinShapeSize
			alpha := rand.Intn(MaxAlpha-MinAlpha) + MinAlpha
			red := rand.Intn(256)
			green := rand.Intn(256)
			blue := rand.Intn(256)
			color := color.RGBA{uint8(red), uint8(green), uint8(blue), uint8(alpha)}
			drawShape(newPng, x, y, size, color)
		}
	}
	savePng(fname, newPng)
}

func drawShape(img *image.RGBA, x, y, size int, c color.RGBA) {
	shapeType := rand.Intn(3)
	switch shapeType {
	case 0:
		drawRectangle(img, x, y, size, size, c)
	case 1:
		drawCircle(img, x, y, size/2, c)
	case 2:
		drawTriangle(img, x, y, size, c)
	}
}

func drawRectangle(img *image.RGBA, x, y, width, height int, c color.RGBA) {
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			img.Set(i, j, c)
		}
	}
}

func drawCircle(img *image.RGBA, x, y, radius int, c color.RGBA) {
	for i := x - radius; i < x+radius; i++ {
		for j := y - radius; j < y+radius; j++ {
			if dist(x, y, i, j) <= float64(radius) {
				img.Set(i, j, c)
			}
		}
	}
}

func drawTriangle(img *image.RGBA, x, y, size int, c color.RGBA) {
	halfSize := size / 2
	for i := 0; i < halfSize; i++ {
		for j := -i; j <= i; j++ {
			img.Set(x+j, y+halfSize+i, c)
		}
	}
}

func dist(x1, y1, x2, y2 int) float64 {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	return math.Sqrt(dx*dx + dy*dy)
}

func noise2(x, y float64) float64 {
	noise := math.Sin(x*NoiseScale) + math.Cos(y*NoiseScale)
	distortion := math.Sin(noise * DistortionScale)
	return distortion
}

func randomPng(fname string, width int,height int ) {
	newPng := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8(rand.Intn(256))
			g := uint8(rand.Intn(256))
			b := uint8(rand.Intn(256))
			newPng.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	savePng(fname, newPng)
}

func trippyFx(fIn string, fOut string) {
	fileIn, err := os.Open(fIn)
	if err != nil {
		fmt.Printf("Failed to decode input .png file: %v", err)
		return
	}
	defer fileIn.Close()
	pngIn, err := png.Decode(fileIn)
	if err != nil {
		fmt.Printf("Failed to decode input .png file: %v", err)
		return
	}
	pngOut := image.NewRGBA(pngIn.Bounds())
	draw.Draw(pngOut, pngIn.Bounds(), pngIn, image.Point{}, draw.Src)
	for y := pngIn.Bounds().Min.Y; y < pngIn.Bounds().Max.Y; y++ {
		for x := pngIn.Bounds().Min.X; x < pngIn.Bounds().Max.X; x++ {
			r, g, b, a := pngIn.At(x, y).RGBA()
			newR := uint8((uint32(x) ^ uint32(y)) * (r + g + b) % 256)
			newG := uint8((uint32(x) * uint32(y)) * (r ^ g ^ b) % 256)
			newB := uint8((uint32(x) + uint32(y)) * (r * g * b) % 256)
			newA := uint8((r + g + b - a) / 3)
			pngOut.Set(x, y, color.RGBA{newR, newG, newB, newA})
		}
	}
	savePng(fOut, pngOut)
	fileIn.Close()
}

func interpolatePngs(fIn1 string, fIn2 string, fOut string, factor float64) {
	fileIn1, err := os.Open(fIn1)
	if err != nil {
		fmt.Printf("Failed to decode input .png file: %v", err)
		return
	}
	defer fileIn1.Close()
	pngIn1, err := png.Decode(fileIn1)
	if err != nil {
		fmt.Printf("Failed to decode input .png file: %v", err)
		return
	}
	fileIn2, err := os.Open(fIn2)
	if err != nil {
		fmt.Printf("Failed to decode input .png file: %v", err)
		return
	}
	defer fileIn2.Close()
	pngIn2, err := png.Decode(fileIn2)
	if err != nil {
		fmt.Printf("Failed to decode input .png file: %v", err)
		return
	}
	bounds := pngIn1.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	newPng := image.NewRGBA(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c1 := pngIn1.At(x, y)
			c2 := pngIn2.At(x, y)
			r1, g1, b1, a1 := c1.RGBA()
			r2, g2, b2, a2 := c2.RGBA()
			r := uint8(float64(r1)*(1.0-factor) + float64(r2)*factor)
			g := uint8(float64(g1)*(1.0-factor) + float64(g2)*factor)
			b := uint8(float64(b1)*(1.0-factor) + float64(b2)*factor)
			a := uint8(float64(a1)*(1.0-factor) + float64(a2)*factor)
			newPng.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}
	savePng(fOut, newPng)
	fileIn1.Close()
	fileIn2.Close()
}