package main

type Sphere struct {
	center Vec3
	radius float64
}

func (s Sphere) didHit(r Ray) bool {
	oc := r.origin().subtract(s.center)

	a := r.direction().dot(r.direction())
	b := 2.0 * oc.dot(r.direction())
	c := oc.dot(oc) - s.radius*s.radius

	discriminant := b*b - 4*a*c

	return discriminant > 0
}

func (s Sphere) colorAt(r Ray) Vec3 {
	if s.didHit(r) {
		return Vec3{1.0, 0.0, 0.0}
	}

	unitDirection := r.direction().unitVector()

	t := 0.5 * (unitDirection.y() + 1.0)

	v1 := Vec3{1.0, 1.0, 1.0}.multiplyScalar(1.0 - t)
	v2 := Vec3{0.5, 0.7, 1.0}.multiplyScalar(t)

	return v1.add(v2)
}
