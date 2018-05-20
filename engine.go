package main

import (
	"math/rand"
)

// Render takes the Camera, Hitables, and Config and outputs the framebuffer.
func Render(camera Camera, world Hitable, config Config) []Vec3 {
	framebuffer := make([]Vec3, 0)

	for j := config.height - 1; j >= 0; j-- {
		for i := 0; i < config.width; i++ {
			color := sample(i, j, config, camera, world)

			framebuffer = append(framebuffer, color)
		}
	}

	return framebuffer
}

func sample(i, j int, config Config, camera Camera, world Hitable) Vec3 {
	color := Vec3Zero()

	for s := 0; s < config.samples; s++ {
		u := (float64(i) + rand.Float64()) / float64(config.width)
		v := (float64(j) + rand.Float64()) / float64(config.height)

		r := camera.getRay(u, v)

		newColor := Color(r, world, 0)

		color.inPlaceAdd(newColor)
	}

	color.inPlaceDivideScalar(float64(config.samples))

	return color
}
