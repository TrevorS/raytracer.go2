package main

import (
	"math"
)

// Color returns a color from a Ray.
func Color(r Ray, hitable Hitable, depth int) Vec3 {
	didHit, hit := hitable.hit(r, 0.001, math.MaxFloat64)

	if didHit {
		didScatter, attenuation, scattered := hit.material.scatter(r, *hit)

		if depth < 50 && didScatter {
			return attenuation.multiply(Color(scattered, hitable, depth+1))
		}

		return Vec3Zero()
	}

	unitDirection := r.direction().unitVector()

	t := 0.5 * (unitDirection.y() + 1.0)

	v1 := Vec3{1.0, 1.0, 1.0}.multiplyScalar(1.0 - t)
	v2 := Vec3{0.5, 0.7, 1.0}.multiplyScalar(t)

	return v1.add(v2)
}
