package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	nx := 200
	ny := 100
	ns := 100

	header := getHeader(nx, ny)
	fmt.Print(header)

	sphere1 := Sphere{Vec3{0.0, 0.0, -1.0}, 0.5}
	sphere2 := Sphere{Vec3{0, -100.5, -1}, 100}

	world := HitableList{sphere1, sphere2}
	camera := NewCamera()

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := Vec3Zero()

			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)

				r := camera.getRay(u, v)

				newColor := Color(r, world)

				color.inPlaceAdd(newColor)
			}

			color.inPlaceDivideScalar(float64(ns))

			ir := int(255.99 * color.r())
			ig := int(255.99 * color.g())
			ib := int(255.99 * color.b())

			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}

func getHeader(nx, ny int) string {
	sNx := strconv.Itoa(nx)
	sNy := strconv.Itoa(ny)

	return "P3\n" + sNx + " " + sNy + "\n255\n"
}
