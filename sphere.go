package main

import (
	"math"
)

// Sphere is a Hitable graphics object.
type Sphere struct {
	center Vec3
	radius float64
}

func (s Sphere) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	oc := r.origin().subtract(s.center)

	a := r.direction().dot(r.direction())
	b := oc.dot(r.direction())
	c := oc.dot(oc) - s.radius*s.radius

	discriminant := b*b - a*c

	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a

		if temp < tMax && temp > tMin {
			p := r.pointAtParameter(temp)

			hit := Hit{
				t:      temp,
				p:      p,
				normal: p.subtract(s.center).divideScalar(s.radius),
			}

			return true, &hit
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a

		if temp < tMax && temp > tMin {
			p := r.pointAtParameter(temp)

			hit := Hit{
				t:      temp,
				p:      p,
				normal: p.subtract(s.center).divideScalar(s.radius),
			}

			return true, &hit
		}
	}

	return false, nil
}

// func (s Sphere) colorAt(r Ray) Vec3 {
// 	t := s.didHit(r)

// 	if t > 0.0 {
// 		N := r.pointAtParameter(t).subtract(Vec3{0, 0, -1}).unitVector()

// 		return Vec3{N.x() + 1, N.y() + 1, N.z() + 1}.multiplyScalar(0.5)
// 	}

// 	unitDirection := r.direction().unitVector()

// 	t = 0.5 * (unitDirection.y() + 1.0)

// 	v1 := Vec3{1.0, 1.0, 1.0}.multiplyScalar(1.0 - t)
// 	v2 := Vec3{0.5, 0.7, 1.0}.multiplyScalar(t)

// 	return v1.add(v2)
// }
