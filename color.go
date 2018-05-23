package main

import (
	"math"
)

// Color returns a color from a Ray.
func Color(r Ray, hitable Hitable, depth int) Vec3 {
	didHit, hit := hitable.hit(r, 0.001, math.MaxFloat64)

	if didHit {
		didScatter, albedo, scattered, _ := hit.material.scatter(r, *hit)
		emitted := hit.material.emitted(r, *hit, hit.u, hit.v, hit.p)

		if depth < 50 && didScatter {
			lightShape := XZRectangle{
				213,
				343,
				227,
				332,
				554,
				NewLambertian(
					ConstantTexture{Vec3{1.0, 1.0, 1.0}},
				),
			}

			hitablePdf := HitablePdf{lightShape, hit.p}
			cosinePdf := NewCosinePdf(hit.normal)
			pdf := NewMixturePdf(hitablePdf, cosinePdf)

			scattered = Ray{hit.p, pdf.generate(), r.time()}
			pdfVal := pdf.value(scattered.direction())

			addition := albedo.multiplyScalar(
				hit.material.scatteringPdf(r, *hit, scattered),
			).multiply(Color(scattered, hitable, depth+1)).divideScalar(pdfVal)

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
