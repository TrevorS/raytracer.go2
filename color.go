package main

import (
	"math"
	"math/rand"
)

// Color returns a color from a Ray.
func Color(r Ray, world HitableList) Vec3 {
	didHit, hit := world.hit(r, 0.0, math.MaxFloat64)

	if didHit {
		target := hit.p.add(hit.normal).add(randomInUnitSphere())
		ray := Ray{hit.p, target.subtract(hit.p)}

		return Color(ray, world).multiplyScalar(0.5)
	}

	unitDirection := r.direction().unitVector()

	t := 0.5 * (unitDirection.y() + 1.0)

	v1 := Vec3{1.0, 1.0, 1.0}.multiplyScalar(1.0 - t)
	v2 := Vec3{0.5, 0.7, 1.0}.multiplyScalar(t)

	return v1.add(v2)
}

func randomInUnitSphere() Vec3 {
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
