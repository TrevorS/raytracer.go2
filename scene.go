package main

import (
	"image"
	"math/rand"
	"os"
)

// SimpleScene returns a HitableList of Spheres for testing.
func SimpleScene(config Config) Hitable {
	world := NewHitableList(0)

	sphere := NewStationarySphere(
		Vec3{0, 0, -1},
		0.5,
		NewLambertian(
			ConstantTexture{Vec3{0.1, 0.2, 0.5}},
		),
	)

	world.add(sphere)

	checkerTexture := CheckerTexture{
		odd:  ConstantTexture{Vec3{0.2, 0.3, 0.1}},
		even: ConstantTexture{Vec3{0.9, 0.9, 0.9}},
	}

	sphere = NewStationarySphere(
		Vec3{0, -100.5, -1},
		100,
		NewLambertian(
			checkerTexture,
		),
	)

	world.add(sphere)

	sphere = NewStationarySphere(
		Vec3{1, 0, -1},
		0.5,
		NewMetal(Vec3{
			0.8, 0.6, 0.2,
		}, 0.0),
	)

	world.add(sphere)

	sphere = NewStationarySphere(
		Vec3{-1, 0, -1},
		0.5,
		NewDielectric(1.5),
	)

	world.add(sphere)

	sphere = NewStationarySphere(
		Vec3{-1, 0, -1},
		-0.45,
		NewDielectric(1.5),
	)

	world.add(sphere)

	return world
}

// RandomScene returns a randomly generated HitableList.
func RandomScene(config Config) Hitable {
	var hitableList HitableList

	sphere := NewStationarySphere(
		Vec3{0, -1000, 0},
		1000,
		NewLambertian(
			ConstantTexture{Vec3{0.5, 0.5, 0.5}},
		),
	)

	hitableList.add(sphere)

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMaterial := rand.Float64()
			center := Vec3{
				float64(a) + 0.9*rand.Float64(),
				0.2,
				float64(b) + 0.9*rand.Float64(),
			}

			if center.subtract(Vec3{4, 0.2, 0}).length() > 0.9 {
				if chooseMaterial < 0.8 {
					sphere := NewMovingSphere(
						center,
						center.add(Vec3{0, 0.5 * rand.Float64(), 0}),
						0.2,
						NewLambertian(
							ConstantTexture{Vec3{
								rand.Float64() * rand.Float64(),
								rand.Float64() * rand.Float64(),
								rand.Float64() * rand.Float64(),
							}},
						),
						0,
						1,
					)

					hitableList.add(sphere)

				} else if chooseMaterial < 0.95 {
					sphere := NewStationarySphere(
						center,
						0.2,
						NewMetal(
							Vec3{
								0.5 * (1 + rand.Float64()),
								0.5 * (1 + rand.Float64()),
								0.5 * (1 + rand.Float64()),
							},
							0.5*rand.Float64(),
						),
					)

					hitableList.add(sphere)

				} else {
					sphere := NewStationarySphere(
						center,
						0.2,
						NewDielectric(1.5),
					)

					hitableList.add(sphere)
				}
			}
		}
	}

	sphere = NewStationarySphere(
		Vec3{0, 1, 0},
		1.0,
		NewDielectric(1.5),
	)

	hitableList.add(sphere)

	sphere = NewStationarySphere(
		Vec3{-4, 1, 0},
		1.0,
		NewLambertian(
			ConstantTexture{Vec3{0.4, 0.2, 0.1}},
		),
	)

	hitableList.add(sphere)

	sphere = NewStationarySphere(
		Vec3{4, 1, 0},
		1.0,
		NewMetal(
			Vec3{
				0.7,
				0.6,
				0.5,
			},
			0.0,
		),
	)

	hitableList.add(sphere)

	bvhNodes := BVHNode{}

	return bvhNodes.newBVHNode(&hitableList, config.timeStart, config.timeEnd)
}

// TwoSpheres is a scene consisting of two checkered spheres.
func TwoSpheres(config Config) Hitable {
	hitables := NewHitableList(0)

	marbleTexture := NewMarbleTexture(4)

	sphere := NewStationarySphere(
		Vec3{0, -1000, 0},
		1000,
		NewLambertian(marbleTexture),
	)

	hitables.add(sphere)

	sphere = NewStationarySphere(
		Vec3{0, 2, 0},
		2,
		NewLambertian(marbleTexture),
	)

	hitables.add(sphere)

	return hitables
}

// EarthSphere returns a single Sphere wrapped in an image texture.
func EarthSphere(config Config, imageFileName string) Hitable {
	hitables := NewHitableList(0)

	imageFile, err := os.Open(imageFileName)
	defer imageFile.Close()

	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(imageFile)

	if err != nil {
		panic(err)
	}

	imageTexture := NewImageTexture(img)

	sphere := NewStationarySphere(
		Vec3{0, 0, 0},
		2,
		NewLambertian(imageTexture),
	)

	hitables.add(sphere)

	return hitables
}

// SimpleLight returns a scene with simple lighting.
func SimpleLight(config Config) Hitable {
	noiseTexture := NewNoiseTexture(4)

	hitables := NewHitableList(0)

	sphere := NewStationarySphere(
		Vec3{0, -1000, 0},
		1000,
		NewLambertian(noiseTexture),
	)

	hitables.add(sphere)

	sphere = NewStationarySphere(
		Vec3{0, 2, 0},
		2,
		NewLambertian(noiseTexture),
	)

	hitables.add(sphere)

	sphere = NewStationarySphere(
		Vec3{0, 7, 0},
		2,
		DiffuseLight{
			ConstantTexture{
				Vec3{4, 4, 4},
			},
		},
	)

	hitables.add(sphere)

	rectangle := XYRectangle{
		3,
		5,
		1,
		3,
		-2,
		DiffuseLight{
			ConstantTexture{
				Vec3{4, 4, 4},
			},
		},
	}

	hitables.add(rectangle)

	return hitables
}

// CornellBox is the standard Cornell scene.
func CornellBox(config Config) Hitable {
	hitables := NewHitableList(0)

	red := NewLambertian(
		ConstantTexture{
			Vec3{
				0.65,
				0.05,
				0.05,
			},
		},
	)

	white := NewLambertian(
		ConstantTexture{
			Vec3{
				0.73,
				0.73,
				0.73,
			},
		},
	)

	green := NewLambertian(
		ConstantTexture{
			Vec3{
				0.12,
				0.45,
				0.15,
			},
		},
	)

	light := DiffuseLight{
		ConstantTexture{
			Vec3{
				15,
				15,
				15,
			},
		},
	}

	flippedYZRectangle := FlipNormals{YZRectangle{
		0,
		555,
		0,
		555,
		555,
		green,
	}}

	hitables.add(flippedYZRectangle)

	yzRectangle := YZRectangle{
		0,
		555,
		0,
		555,
		0,
		red,
	}

	hitables.add(yzRectangle)

	xzRectangle := XZRectangle{
		213,
		343,
		227,
		332,
		554,
		light,
	}

	hitables.add(xzRectangle)

	flippedXZRectangle := FlipNormals{XZRectangle{
		0,
		555,
		0,
		555,
		555,
		white,
	}}

	hitables.add(flippedXZRectangle)

	xzRectangle = XZRectangle{
		0,
		555,
		0,
		555,
		0,
		white,
	}

	hitables.add(xzRectangle)

	flippedXYRectangle := FlipNormals{XYRectangle{
		0,
		555,
		0,
		555,
		555,
		white,
	}}

	hitables.add(flippedXYRectangle)

	box := NewBox(
		Vec3{130, 0, 65},
		Vec3{295, 165, 230},
		white,
	)

	hitables.add(box)

	box = NewBox(
		Vec3{265, 0, 295},
		Vec3{430, 330, 460},
		white,
	)

	hitables.add(box)

	return hitables
}
