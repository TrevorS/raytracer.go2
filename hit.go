package main

// Hitable represents hitable graphical objects.
type Hitable interface {
	hit(r Ray, tMin, tMax float64) (bool, *Hit)
	boundingBox(t0, t1 float64) (bool, *AABB)
}

// Hit is a record of a Hitable object being hit.
type Hit struct {
	t        float64
	p        Vec3
	normal   Vec3
	material Material
}
