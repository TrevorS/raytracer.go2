package main

// Hitable represents hitable graphical objects.
type Hitable interface {
	hit(r Ray, tMin, tMax float64) (bool, *Hit)
}

// Hit is a record of a Hitable object being hit.
type Hit struct {
	t      float64
	p      Vec3
	normal Vec3
}

// HitableList is an array of Hitable graphics objects.
type HitableList []Hitable

func (hList HitableList) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	hitAnything := false
	closest := tMax

	var closestHit Hit

	for _, hitable := range hList {
		didHit, hit := hitable.hit(r, tMin, closest)

		if didHit {
			hitAnything = true
			closest = hit.t
			closestHit = *hit
		}
	}

	return hitAnything, &closestHit
}
