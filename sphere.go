package main

import (
	"math"
)

type Sphere struct {
	center Vec3
	radius float64
}

func (s Sphere) didHit(r Ray) float64 {
	oc := r.origin().subtract(s.center)

	a := r.direction().dot(r.direction())
	b := 2.0 * oc.dot(r.direction())
	c := oc.dot(oc) - s.radius*s.radius

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return -1.0
	}

	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}

func (s Sphere) colorAt(r Ray) Vec3 {
	t := s.didHit(r)

	if t > 0.0 {
		N := r.pointAtParameter(t).subtract(Vec3{0, 0, -1}).unitVector()

		return Vec3{N.x() + 1, N.y() + 1, N.z() + 1}.multiplyScalar(0.5)
	}

	unitDirection := r.direction().unitVector()

	t = 0.5 * (unitDirection.y() + 1.0)

	v1 := Vec3{1.0, 1.0, 1.0}.multiplyScalar(1.0 - t)
	v2 := Vec3{0.5, 0.7, 1.0}.multiplyScalar(t)

	return v1.add(v2)
}
