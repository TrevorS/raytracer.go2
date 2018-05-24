package main

import (
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

// Render takes the Camera, Hitables, and Config and outputs the framebuffer.
func Render(camera Camera, world, lightShapes Hitable, config Config) []Vec3 {
	framebuffer := make([]Vec3, 0)

	for j := config.height - 1; j >= 0; j-- {
		for i := 0; i < config.width; i++ {
			color := sampling(i, j, config, camera, world, lightShapes)

			framebuffer = append(framebuffer, color)
		}
	}

	return framebuffer
}

func sampling(i, j int, config Config, camera Camera, world, lightShapes Hitable) Vec3 {
	color := Vec3Zero()

	samples := make(chan Vec3, config.samples)

	for s := 0; s < config.samples; s++ {
		wg.Add(1)
		go sample(i, j, config.width, config.height, camera, world, lightShapes, samples)
	}

	wg.Wait()
	close(samples)

	for newColor := range samples {
		color.inPlaceAdd(newColor)
	}

	color.inPlaceDivideScalar(float64(config.samples))

	return color
}

func sample(i, j, width, height int, camera Camera, world Hitable, lightShapes Hitable, samples chan Vec3) {
	defer wg.Done()

	u := (float64(i) + rand.Float64()) / float64(width)
	v := (float64(j) + rand.Float64()) / float64(height)

	r := camera.getRay(u, v)

	samples <- Color(r, world, lightShapes, 0)
}
