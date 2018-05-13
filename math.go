package main

import (
	"math"
	"math/rand"
)

// RandomInUnitSphere returns a random Vector within the unit sphere.
func RandomInUnitSphere() Vec3 {
	for {
		p := Vec3{
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		}.multiplyScalar(2.0).subtract(Vec3{1.0, 1.0, 1.0})

		if p.squaredLength() < 1.0 {
			return p
		}
	}
}

// RandomInUnitDisk returns a random Vector within the unit disk.
func RandomInUnitDisk() Vec3 {
	for {
		p := Vec3{
			rand.Float64(),
			rand.Float64(),
			0,
		}.multiplyScalar(2.0).subtract(Vec3{1.0, 1.0, 0})

		if p.dot(p) < 1.0 {
			return p
		}
	}
}

// Schlick calculates an approximation of reflectivity varied by angle.
func Schlick(cosine, refractiveIndex float64) float64 {
	r0 := (1 - refractiveIndex) / (1 + refractiveIndex)

	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
