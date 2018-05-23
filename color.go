package main

import (
	"math"
)

// Color returns a color from a Ray.
func Color(r Ray, hitable Hitable, depth int) Vec3 {
	didHit, hit := hitable.hit(r, 0.001, math.MaxFloat64)

	if didHit {
		didScatter, albedo, scattered, pdf := hit.material.scatter(r, *hit)
		emitted := hit.material.emitted(hit.u, hit.v, hit.p)

		if depth < 50 && didScatter {
			addition := albedo.multiplyScalar(
				hit.material.scatteringPdf(r, *hit, scattered),
			).multiply(Color(scattered, hitable, depth+1)).divideScalar(pdf)

			return emitted.add(addition)
		}

		return emitted
	}

	return EmitBlack()
}

// EmitBlack emits the Color black.
func EmitBlack() Vec3 {
	return Vec3Zero()
}
