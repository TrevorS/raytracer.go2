package main

import (
	"math"
	"math/rand"
)

// ConstantMedium is a medium of constant density.
type ConstantMedium struct {
	hitable  Hitable
	density  float64
	material Material
}

// NewConstantMedium properly initializes a ConstantMedium.
func NewConstantMedium(hitable Hitable, density float64, texture Texture) ConstantMedium {
	return ConstantMedium{
		hitable,
		density,
		Isotropic{
			texture,
		},
	}
}

func (cm ConstantMedium) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	didHit1, hit1 := cm.hitable.hit(r, -math.MaxFloat64, math.MaxFloat64)

	if didHit1 {
		didHit2, hit2 := cm.hitable.hit(r, hit1.t+0.0001, math.MaxFloat64)

		if didHit2 {
			if hit1.t < tMin {
				hit1.t = tMin
			}

			if hit2.t > tMax {
				hit2.t = tMax
			}

			if hit1.t >= hit2.t {
				return false, nil
			}

			if hit1.t < 0 {
				hit1.t = 0
			}

			distanceInsideBoundary := (hit2.t - hit1.t) * r.direction().length()
			hitDistance := -(1 / cm.density) * math.Log(rand.Float64())

			if hitDistance < distanceInsideBoundary {
				t := hit1.t + hitDistance/r.direction().length()
				p := r.pointAtParameter(t)

				normal := Vec3{1, 0, 0}
				material := cm.material

				hit := Hit{
					t,
					hit1.u,
					hit1.v,
					p,
					normal,
					material,
				}

				return true, &hit
			}
		}
	}

	return false, nil
}

func (cm ConstantMedium) boundingBox(t0, t1 float64) (bool, *AABB) {
	return cm.hitable.boundingBox(t0, t1)
}

func (cm ConstantMedium) pdfValue(o, direction Vec3) float64 {
	return 0.0
}

func (cm ConstantMedium) random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
