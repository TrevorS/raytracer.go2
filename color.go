package main

import (
	"math"
)

// Color returns a color from a Ray.
func Color(r Ray, world HitableList) Vec3 {
	didHit, hit := world.hit(r, 0.0, math.MaxFloat64)

	if didHit {
		return Vec3{
			hit.normal.x() + 1,
			hit.normal.y() + 1,
			hit.normal.z() + 1,
		}.multiplyScalar(0.5)
	}

	unitDirection := r.direction().unitVector()

	t := 0.5 * (unitDirection.y() + 1.0)

	v1 := Vec3{1.0, 1.0, 1.0}.multiplyScalar(1.0 - t)
	v2 := Vec3{0.5, 0.7, 1.0}.multiplyScalar(t)

	return v1.add(v2)
}
