package main

import (
	"math/rand"
)

func main() {
	config := Config{
		width:     1024,
		height:    768,
		samples:   100,
		from:      Vec3{278, 278, -800},
		at:        Vec3{278, 278, 0},
		up:        Vec3{0, 1, 0},
		fov:       40.0,
		aperture:  0.00,
		filename:  "output.png",
		timeStart: 0,
		timeEnd:   1,
	}

	world := CornellSmoke(config)

	camera := NewCamera(
		config.from,
		config.at,
		config.up,
		config.fov,
		config.aspectRatio(),
		config.aperture,
		config.focusDistance(),
		config.timeStart,
		config.timeEnd,
	)

	framebuffer := make([]Vec3, 0)

	for j := config.height - 1; j >= 0; j-- {
		for i := 0; i < config.width; i++ {
			color := Vec3Zero()

			for s := 0; s < config.samples; s++ {
				u := (float64(i) + rand.Float64()) / float64(config.width)
				v := (float64(j) + rand.Float64()) / float64(config.height)

				r := camera.getRay(u, v)

				newColor := Color(r, world, 0)

				color.inPlaceAdd(newColor)
			}

			color.inPlaceDivideScalar(float64(config.samples))

			framebuffer = append(framebuffer, color)
		}
	}

	WriteImage(&framebuffer, config)
}
