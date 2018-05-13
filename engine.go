package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	nx := 1024
	ny := 768
	ns := 100

	header := getHeader(nx, ny)
	fmt.Print(header)

	world := RandomScene()

	from := Vec3{13, 2, 3}
	at := Vec3{0, 0, 0}
	up := Vec3{0, 1, 0}
	fvov := 75.0
	aspect := float64(nx) / float64(ny)
	distToFocus := from.subtract(at).length()
	aperture := 0.01

	camera := NewCamera(
		from,
		at,
		up,
		fvov,
		aspect,
		aperture,
		distToFocus,
	)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := Vec3Zero()

			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)

				r := camera.getRay(u, v)

				newColor := Color(r, world, 0)

				color.inPlaceAdd(newColor)
			}

			color.inPlaceDivideScalar(float64(ns))

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
