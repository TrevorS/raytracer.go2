package main

// Ray represents a light ray.
type Ray struct {
	a  Vec3
	b  Vec3
	ti float64
}

func (r Ray) origin() Vec3 {
	return r.a
}

func (r Ray) direction() Vec3 {
	return r.b
}

func (r Ray) time() float64 {
	return r.ti
}

func (r Ray) pointAtParameter(t float64) Vec3 {
	return r.a.add(r.b.multiplyScalar(t))
}
