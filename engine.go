package main

import (
	"math/rand"
)

func main() {
	config := Config{
		width:    500,
		height:   500,
		samples:  100,
		from:     Vec3{13, 2, 3},
		at:       Vec3{0, 0, 0},
		up:       Vec3{0, 1, 0},
		fov:      75.0,
		aperture: 0.01,
		filename: "output.png",
	}

	world := RandomScene()

	camera := NewCamera(
		config.from,
		config.at,
		config.up,
		config.fov,
		config.aspectRatio(),
		config.aperture,
		config.focusDistance(),
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
