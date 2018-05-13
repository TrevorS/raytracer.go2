package main

import "math/rand"

// RandomInUnitSphere returns a random Vector within the unit sphere.
func RandomInUnitSphere() Vec3 {
	for {
		p := Vec3{
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		}.multiplyScalar(2.0).subtract(Vec3{1.0, 1.0, 1.0})

		if p.squaredLength() >= 1.0 {
			return p
		}
	}
}
