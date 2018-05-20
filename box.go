package main

// Box is a cube Hitable.
type Box struct {
	pMin     Vec3
	pMax     Vec3
	hitables HitableList
}

// NewBox returns a properly instantiated Box.
func NewBox(p0, p1 Vec3, material Material) Box {
	hitables := NewHitableList(0)

	hitables.add(
		XYRectangle{
			p0.x(),
			p1.x(),
			p0.y(),
			p1.y(),
			p1.z(),
			material,
		},
	)

	hitables.add(
		FlipNormals{
			XYRectangle{
				p0.x(),
				p1.x(),
				p0.y(),
				p1.y(),
				p0.z(),
				material,
			},
		},
	)

	hitables.add(
		XZRectangle{
			p0.x(),
			p1.x(),
			p0.z(),
			p1.z(),
			p1.y(),
			material,
		},
	)

	hitables.add(
		FlipNormals{
			XZRectangle{
				p0.x(),
				p1.x(),
				p0.z(),
				p1.z(),
				p1.y(),
				material,
			},
		},
	)

	hitables.add(
		YZRectangle{
			p0.y(),
			p1.y(),
			p0.z(),
			p1.z(),
			p1.x(),
			material,
		},
	)

	hitables.add(
		FlipNormals{
			YZRectangle{
				p0.y(),
				p1.y(),
				p0.z(),
				p1.z(),
				p1.x(),
				material,
			},
		},
	)

	return Box{
		p0,
		p1,
		hitables,
	}
}

func (b Box) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	return b.hitables.hit(r, tMin, tMax)
}

func (b Box) boundingBox(t0, t1 float64) (hasBox bool, box *AABB) {
	bbox := AABB{
		b.pMin,
		b.pMax,
	}

	return true, &bbox
}
