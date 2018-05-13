package main

// Camera represents the viewpoint of our scene.
type Camera struct {
	lowerLeftCorner Vec3
	horizontal      Vec3
	vertical        Vec3
	origin          Vec3
}

// NewCamera returns an pre-initialized Camera.
func NewCamera() Camera {
	return Camera{
		lowerLeftCorner: Vec3{-2.0, -1.0, -1.0},
		horizontal:      Vec3{4.0, 0.0, 0.0},
		vertical:        Vec3{0.0, 2.0, 0.0},
		origin:          Vec3{0.0, 0.0, 0.0},
	}
}

func (c Camera) getRay(u, v float64) Ray {
	direction := c.lowerLeftCorner.add(c.horizontal.multiplyScalar(u)).add(c.vertical.multiplyScalar(v)).subtract(c.origin)

	return Ray{
		c.origin,
		direction,
	}
}
