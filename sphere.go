package main

import (
	"math"
)

// Sphere is a Hitable graphics object.
type Sphere struct {
	centerStart  Vec3
	centerFinish Vec3
	radius       float64
	material     Material
	timeStart    float64
	timeFinish   float64
}

// NewStationarySphere returns a Sphere that does not move.
func NewStationarySphere(center Vec3, radius float64, material Material) Sphere {
	return Sphere{
		center,
		center,
		radius,
		material,
		0,
		0,
	}
}

// NewMovingSphere returns a Sphere that moves from centerStart to centerFinish over timeStart to timeFinish.
func NewMovingSphere(centerStart, centerFinish Vec3, radius float64, material Material, timeStart, timeFinish float64) Sphere {
	return Sphere{
		centerStart,
		centerFinish,
		radius,
		material,
		timeStart,
		timeFinish,
	}
}

func (s Sphere) center(time float64) Vec3 {
	if s.centerStart == s.centerFinish {
		return s.centerStart
	}

	return s.centerStart.add((s.centerFinish.subtract(s.centerStart)).multiplyScalar((time - s.timeStart) / (s.timeFinish - s.timeStart)))
}

func (s Sphere) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	oc := r.origin().subtract(s.center(r.time()))

	a := r.direction().dot(r.direction())
	b := oc.dot(r.direction())
	c := oc.dot(oc) - s.radius*s.radius

	discriminant := b*b - a*c

	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a

		if temp < tMax && temp > tMin {
			p := r.pointAtParameter(temp)

			u, v := GetSphereUV(p.subtract(s.center(r.time())).divideScalar(s.radius))

			hit := Hit{
				t:        temp,
				p:        p,
				u:        u,
				v:        v,
				normal:   p.subtract(s.center(r.time())).divideScalar(s.radius),
				material: s.material,
			}

			return true, &hit
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a

		if temp < tMax && temp > tMin {
			p := r.pointAtParameter(temp)

			u, v := GetSphereUV(p.subtract(s.center(r.time())).divideScalar(s.radius))

			hit := Hit{
				t:        temp,
				p:        p,
				u:        u,
				v:        v,
				normal:   p.subtract(s.center(r.time())).divideScalar(s.radius),
				material: s.material,
			}

			return true, &hit
		}
	}

	return false, nil
}

func (s Sphere) boundingBox(t0, t1 float64) (hasBox bool, box *AABB) {
	t0Box := AABB{
		s.center(t0).subtract(Vec3{s.radius, s.radius, s.radius}),
		s.center(t0).add(Vec3{s.radius, s.radius, s.radius}),
	}

	t1Box := AABB{
		s.center(t1).subtract(Vec3{s.radius, s.radius, s.radius}),
		s.center(t1).add(Vec3{s.radius, s.radius, s.radius}),
	}

	return true, SurroundingBox(t0Box, t1Box)
}

func (s Sphere) pdfValue(o, direction Vec3) float64 {
	return 0.0
}

func (s Sphere) random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
