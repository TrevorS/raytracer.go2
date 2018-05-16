package main

import (
	"math"
	"math/rand"
)

// Camera represents the viewpoint of our scene.
type Camera struct {
	lowerLeftCorner Vec3
	horizontal      Vec3
	vertical        Vec3
	origin          Vec3
	u               Vec3
	v               Vec3
	w               Vec3
	lensRadius      float64
	time0           float64
	time1           float64
}

// NewCamera returns an pre-initialized Camera.
func NewCamera(from, at, up Vec3, vfov, aspect, aperture, focusDistance, t0, t1 float64) Camera {
	lensRadius := aperture / 2

	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	origin := from

	w := from.subtract(at).unitVector()
	u := up.cross(w).unitVector()
	v := w.cross(u)

	lowerLeftCorner := origin.subtract(u.multiplyScalar(halfWidth * focusDistance)).subtract(v.multiplyScalar(halfHeight * focusDistance)).subtract(w.multiplyScalar(focusDistance))

	horizontal := u.multiplyScalar(halfWidth * 2 * focusDistance)
	vertical := v.multiplyScalar(halfHeight * 2 * focusDistance)

	return Camera{
		lowerLeftCorner,
		horizontal,
		vertical,
		origin,
		u,
		v,
		w,
		lensRadius,
		t0,
		t1,
	}
}

func (c Camera) getRay(s, t float64) Ray {
	rd := RandomInUnitDisk().multiplyScalar(c.lensRadius)
	offset := c.u.multiplyScalar(rd.x()).add(c.v.multiplyScalar(rd.y()))

	origin := c.origin.add(offset)
	direction := c.lowerLeftCorner.add(c.horizontal.multiplyScalar(s)).add(c.vertical.multiplyScalar(t)).subtract(c.origin).subtract(offset)
	time := c.time0 + rand.Float64()*(c.time1-c.time0)

	return Ray{
		origin,
		direction,
		time,
	}
}
