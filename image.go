package main

import (
	"image"
	"image/png"
	"os"
)

// WriteImage writes out a PNG based on the passed in framebuffer and config.
func WriteImage(framebuffer *[]Vec3, config Config) {
	img := image.NewNRGBA(image.Rect(0, 0, config.width, config.height))

	for j := 0; j < config.height; j++ {
		for i := 0; i < config.width; i++ {
			pixelIndex := j*config.width + i

			pixel := (*framebuffer)[pixelIndex]

			img.Set(i, j, pixel.sqrt())
		}
	}

	write(img, config.filename)
}

func write(img image.Image, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
