package main

import (
	"math"
)

// AABB is an axis-aligned bounding box.
type AABB struct {
	min Vec3
	max Vec3
}

// SurroundingBox returns an AABB containing two boxes.
func SurroundingBox(box0, box1 AABB) *AABB {
	small := Vec3{
		math.Min(box0.min.x(), box1.min.x()),
		math.Min(box0.min.y(), box1.min.y()),
		math.Min(box0.min.z(), box1.min.z()),
	}

	big := Vec3{
		math.Max(box0.max.x(), box1.max.x()),
		math.Max(box0.max.y(), box1.max.y()),
		math.Max(box0.max.z(), box1.max.z()),
	}

	return &AABB{small, big}
}

func (box AABB) hit(r Ray, tMin, tMax float64) bool {
	for a := 0; a < 3; a++ {
		t0 := math.Min(
			(box.min.get(a)-r.origin().get(a))/r.direction().get(a),
			(box.max.get(a)-r.origin().get(a))/r.direction().get(a),
		)

		t1 := math.Max(
			(box.min.get(a)-r.origin().get(a))/r.direction().get(a),
			(box.max.get(a)-r.origin().get(a))/r.direction().get(a),
		)

		tMin := math.Max(t0, tMin)
		tMax := math.Min(t1, tMax)

		if tMax <= tMin {
			return false
		}
	}

	return true
}
