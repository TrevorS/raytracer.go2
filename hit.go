package main

// Hitable represents hitable graphical objects.
type Hitable interface {
	hit(r Ray, tMin, tMax float64) (bool, *Hit)
}

// Hit is a record of a Hitable object being hit.
type Hit struct {
	t        float64
	p        Vec3
	normal   Vec3
	material Material
}

// HitableList is an array of Hitable graphics objects.
type HitableList []Hitable

// NewHitableList returns a dynamically sized HitableList.
func NewHitableList(size int) HitableList {
	return make(HitableList, size)
}

func (hList *HitableList) add(hitable Hitable) HitableList {
	*hList = append(*hList, hitable)

	return *hList
}

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
