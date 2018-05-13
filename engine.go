package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	config := Config{
		width:    500,
		height:   500,
		samples:  10,
		from:     Vec3{13, 2, 3},
		at:       Vec3{0, 0, 0},
		up:       Vec3{0, 1, 0},
		fov:      75.0,
		aperture: 0.01,
	}

	header := getHeader(config.width, config.height)
	fmt.Print(header)

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

			gammaCorrectedColor := Vec3{
				math.Sqrt(color.r()),
				math.Sqrt(color.g()),
				math.Sqrt(color.b()),
			}

			ir := int(255.99 * gammaCorrectedColor.r())
			ig := int(255.99 * gammaCorrectedColor.g())
			ib := int(255.99 * gammaCorrectedColor.b())

			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}

func getHeader(nx, ny int) string {
	sNx := strconv.Itoa(nx)
	sNy := strconv.Itoa(ny)

	return "P3\n" + sNx + " " + sNy + "\n255\n"
}
