package main

import (
	"math"
)

// Camera represents the viewpoint of our scene.
type Camera struct {
	lowerLeftCorner Vec3
	horizontal      Vec3
	vertical        Vec3
	origin          Vec3
}

// NewCamera returns an pre-initialized Camera.
func NewCamera(from, at, up Vec3, vfov, aspect float64) Camera {
	theta := vfov * math.Pi / 180

	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	origin := from

	w := from.subtract(at).unitVector()
	u := up.cross(w).unitVector()
	v := w.cross(u)

	lowerLeftCorner := origin.subtract(u.multiplyScalar(halfWidth)).subtract(v.multiplyScalar(halfHeight)).subtract(w)
	horizontal := u.multiplyScalar(halfWidth * 2)
	vertical := v.multiplyScalar(halfHeight * 2)

	return Camera{
		lowerLeftCorner,
		horizontal,
		vertical,
		origin,
	}
}

func (c Camera) getRay(u, v float64) Ray {
	direction := c.lowerLeftCorner.add(c.horizontal.multiplyScalar(u)).add(c.vertical.multiplyScalar(v)).subtract(c.origin)

	return Ray{
		c.origin,
		direction,
	}
}
