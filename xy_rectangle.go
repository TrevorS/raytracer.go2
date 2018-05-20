package main

// XYRectangle represents an axis-aligned rectangle.
type XYRectangle struct {
	x0       float64
	x1       float64
	y0       float64
	y1       float64
	k        float64
	material Material
}

func (rec XYRectangle) hit(r Ray, t0, t1 float64) (bool, *Hit) {
	t := (rec.k - r.origin().z()) / r.direction().z()

	if t < t0 || t > t1 {
		return false, nil
	}

	x := r.origin().x() + r.direction().multiplyScalar(t).x()
	y := r.origin().y() + r.direction().multiplyScalar(t).y()

	if x < rec.x0 || x > rec.x1 || y < rec.y0 || y > rec.y1 {
		return false, nil
	}

	u := (x - rec.x0) / (rec.x1 - rec.x0)
	v := (y - rec.y0) / (rec.y1 - rec.y0)

	p := r.pointAtParameter(t)

	normal := Vec3{0, 0, 1}

	hit := Hit{
		t:        t,
		p:        p,
		u:        u,
		v:        v,
		normal:   normal,
		material: rec.material,
	}

	return true, &hit
}

func (rec XYRectangle) boundingBox(t0, t1 float64) (bool, *AABB) {
	box := AABB{
		Vec3{rec.x0, rec.y0, rec.k - 0.0001},
		Vec3{rec.x1, rec.y1, rec.k + 0.0001},
	}

	return true, &box
}
