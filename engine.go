package main

import (
	"fmt"
	"strconv"
)

func main() {
	nx := 200
	ny := 100

	header := getHeader(nx, ny)
	fmt.Print(header)

	lowerLeftCorner := Vec3{-2.0, -1.0, -1.0}
	horizontal := Vec3{4.0, 0.0, 0.0}
	vertical := Vec3{0.0, 2.0, 0.0}
	origin := Vec3{0.0, 0.0, 0.0}

	sphere1 := Sphere{Vec3{0.0, 0.0, -1.0}, 0.5}
	sphere2 := Sphere{Vec3{0, -100.5, -1}, 100}

	world := HitableList{sphere1, sphere2}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			direction := lowerLeftCorner.add(horizontal.multiplyScalar(u).add(vertical.multiplyScalar(v)))

			ray := Ray{
				origin,
				direction,
			}

			color := Color(ray, world)

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
