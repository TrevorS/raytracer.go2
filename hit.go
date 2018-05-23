package main

import "math"

// Hitable represents hitable graphical objects.
type Hitable interface {
	hit(r Ray, tMin, tMax float64) (bool, *Hit)
	boundingBox(t0, t1 float64) (bool, *AABB)
	pdfValue(o, direction Vec3) float64
	random(o Vec3) Vec3
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

func (fn FlipNormals) pdfValue(o, direction Vec3) float64 {
	return 0.0
}

func (fn FlipNormals) random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}

// Translate moves a Hitable by an offset.
type Translate struct {
	hitable Hitable
	offset  Vec3
}

func (ts Translate) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	rayMoved := Ray{
		r.origin().subtract(ts.offset),
		r.direction(),
		r.time(),
	}

	didHit, hit := ts.hitable.hit(rayMoved, tMin, tMax)

	if didHit {
		hit.p.inPlaceAdd(ts.offset)

		return didHit, hit
	}

	return didHit, nil
}

func (ts Translate) boundingBox(t0, t1 float64) (bool, *AABB) {
	hasBox, boundingBox := ts.hitable.boundingBox(t0, t1)

	if hasBox {
		movedBox := &AABB{
			boundingBox.min.add(ts.offset),
			boundingBox.max.add(ts.offset),
		}

		return hasBox, movedBox
	}

	return hasBox, nil
}

func (ts Translate) pdfValue(o, direction Vec3) float64 {
	return 0.0
}

func (ts Translate) random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}

// RotateY is a Hitable that contains a Y rotated Hitable.
type RotateY struct {
	hitable  Hitable
	sinTheta float64
	cosTheta float64
	hasBox   bool
	bbox     AABB
}

// NewRotateY properly instantiates a RotateY Hitable.
func NewRotateY(hitable Hitable, angle float64) RotateY {
	radians := (math.Pi / 180.0) * angle
	sinTheta := math.Sin(radians)
	cosTheta := math.Cos(radians)

	hasBox, boundingBox := hitable.boundingBox(0, 1)

	min := Vec3{
		math.MaxFloat64,
		math.MaxFloat64,
		math.MaxFloat64,
	}

	max := Vec3{
		-1 * math.MaxFloat64,
		-1 * math.MaxFloat64,
		-1 * math.MaxFloat64,
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				fI := float64(i)
				fJ := float64(j)
				fK := float64(k)

				x := fI*boundingBox.max.x() + (1-fI)*boundingBox.min.x()
				y := fJ*boundingBox.max.y() + (1-fJ)*boundingBox.min.y()
				z := fK*boundingBox.max.z() + (1-fK)*boundingBox.min.z()

				newX := cosTheta*x + sinTheta*z
				newZ := -sinTheta*x + cosTheta*z

				tester := Vec3{newX, y, newZ}

				for c := 0; c < 3; c++ {
					newValue := tester.get(c)

					if newValue > max.get(c) {
						max.set(c, newValue)
					}

					if newValue < min.get(c) {
						min.set(c, newValue)
					}
				}
			}
		}
	}

	return RotateY{
		hitable,
		sinTheta,
		cosTheta,
		hasBox,
		AABB{
			min,
			max,
		},
	}
}

func (ry RotateY) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	origin := r.origin()
	direction := r.direction()

	origin.e0 = ry.cosTheta*r.origin().x() - ry.sinTheta*r.origin().z()
	origin.e2 = ry.sinTheta*r.origin().x() + ry.cosTheta*r.origin().z()

	direction.e0 = ry.cosTheta*r.direction().x() - ry.sinTheta*r.direction().z()
	direction.e2 = ry.sinTheta*r.direction().x() + ry.cosTheta*r.direction().z()

	rotatedRay := Ray{
		origin,
		direction,
		r.time(),
	}

	didHit, hit := ry.hitable.hit(rotatedRay, tMin, tMax)

	if didHit {
		p := hit.p
		normal := hit.normal

		p.e0 = ry.cosTheta*hit.p.x() + ry.sinTheta*hit.p.z()
		p.e2 = -ry.sinTheta*hit.p.x() + ry.cosTheta*hit.p.z()

		normal.e0 = ry.cosTheta*hit.normal.x() + ry.sinTheta*hit.normal.z()
		normal.e2 = -ry.sinTheta*hit.normal.x() + ry.cosTheta*hit.normal.z()

		hit.p = p
		hit.normal = normal

		return didHit, hit
	}

	return didHit, nil
}

func (ry RotateY) boundingBox(t0, t1 float64) (bool, *AABB) {
	return ry.hasBox, &ry.bbox
}

func (ry RotateY) pdfValue(o, direction Vec3) float64 {
	return 0.0
}

func (ry RotateY) random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
