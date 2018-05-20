package main

import (
	"math"
)

// Color returns a color from a Ray.
func Color(r Ray, hitable Hitable, depth int) Vec3 {
	didHit, hit := hitable.hit(r, 0.001, math.MaxFloat64)

	if didHit {
		didScatter, attenuation, scattered := hit.material.scatter(r, *hit)
		emitted := hit.material.emitted(hit.u, hit.v, hit.p)

		if depth < 50 && didScatter {
			return emitted.add(attenuation.multiply(Color(scattered, hitable, depth+1)))
		}

		return emitted
	}

	return Vec3Zero()
}
