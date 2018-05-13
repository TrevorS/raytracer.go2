package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	nx := 200
	ny := 100
	ns := 100

	header := getHeader(nx, ny)
	fmt.Print(header)

	sphere1 := Sphere{
		center: Vec3{0.0, 0.0, -1.0},
		radius: 0.5,
		material: NewLambertian(
			Vec3{0.1, 0.2, 0.5},
		),
	}

	sphere2 := Sphere{
		center: Vec3{0, -100.5, -1},
		radius: 100,
		material: NewLambertian(
			Vec3{0.8, 0.8, 0.0},
		),
	}

	sphere3 := Sphere{
		center: Vec3{1.0, 0.0, -1.0},
		radius: 0.5,
		material: NewMetal(
			Vec3{0.8, 0.6, 0.2},
			0.2,
		),
	}

	sphere4 := Sphere{
		center: Vec3{-1.0, 0.0, -1.0},
		radius: 0.5,
		material: NewDielectric(
			1.5,
		),
	}

	sphere5 := Sphere{
		center: Vec3{-1.0, 0.0, -1.0},
		radius: -0.45,
		material: NewDielectric(
			1.5,
		),
	}

	world := HitableList{
		sphere1,
		sphere2,
		sphere3,
		sphere4,
		sphere5,
	}

	from := Vec3{3, 3, 2}
	at := Vec3{0, 0, -1}
	up := Vec3{0, 1, 0}
	fvov := 20.0
	aspect := float64(nx) / float64(ny)
	distToFocus := from.subtract(at).length()
	aperture := 2.0

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
