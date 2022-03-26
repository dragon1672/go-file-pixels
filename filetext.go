package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

var (
	fileToRead = flag.String("file", "image.png", "file to load.")
)

type Pixel struct {
	R, G, B, A int
}
type Pos struct {
	X, Y int
}

func getPixels(img image.Image) map[Pos]Pixel {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	ret := make(map[Pos]Pixel)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ret[Pos{x, y}] = colorToPixel(img.At(x, y))
		}
	}

	return ret
}

func colorToPixel(c color.Color) Pixel {
	r, g, b, a := c.RGBA()
	return Pixel{
		R: int(r >> 8),
		G: int(g >> 8),
		B: int(b >> 8),
		A: int(a >> 8),
	}
}

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	file, err := os.Open(*fileToRead)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Error decoding image: %v", err)
	}

	pixels := getPixels(img)

	fmt.Println(pixels)
}
