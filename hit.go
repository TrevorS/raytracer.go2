package main

// Hitable represents hitable graphical objects.
type Hitable interface {
	hit(r Ray, tMin, tMax float64) (bool, *Hit)
	boundingBox(t0, t1 float64) (bool, *AABB)
}

// Hit is a record of a Hitable object being hit.
type Hit struct {
	t        float64
	u        float64
	v        float64
	p        Vec3
	normal   Vec3
	material Material
}

// FlipNormals accepts a Hitable and reverses the normal.
type FlipNormals struct {
	hitable Hitable
}

func (fn FlipNormals) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	didHit, hit := fn.hitable.hit(r, tMin, tMax)

	if didHit {
		hit.normal = hit.normal.negate()

		return true, hit
	}

	return false, nil
}

func (fn FlipNormals) boundingBox(t0, t1 float64) (bool, *AABB) {
	return fn.hitable.boundingBox(t0, t1)
}
