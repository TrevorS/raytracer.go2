package main

import (
	"math"
)

// Color returns a color from a Ray.
func Color(r Ray, hitable Hitable, lightShape Hitable, depth int) Vec3 {
	didHit, hit := hitable.hit(r, 0.001, math.MaxFloat64)

	if didHit {
		didScatter, scatter := hit.material.scatter(r, *hit)
		emitted := hit.material.emitted(r, *hit, hit.u, hit.v, hit.p)

		if depth < 50 && didScatter {
			if scatter.isSpecular {
				return scatter.attenuation.multiply(
					Color(scatter.specularRay, hitable, lightShape, depth+1),
				)
			}

			hitablePdf := HitablePdf{lightShape, hit.p}
			pdf := NewMixturePdf(hitablePdf, scatter.pdf)

			scattered := Ray{hit.p, pdf.generate(), r.time()}
			pdfVal := pdf.value(scattered.direction())

			addition := scatter.attenuation.multiplyScalar(
				hit.material.scatteringPdf(r, *hit, scattered),
			).multiply(Color(scattered, hitable, lightShape, depth+1)).divideScalar(pdfVal)

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
