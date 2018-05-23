package main

import (
	"math"
)

// Onb represents an ortho-normal basis.
type Onb struct {
	axis [3]Vec3
}

func (o Onb) get(i int) Vec3 {
	return o.axis[i]
}

func (o Onb) u() Vec3 {
	return o.axis[0]
}

func (o Onb) v() Vec3 {
	return o.axis[1]
}

func (o Onb) w() Vec3 {
	return o.axis[2]
}

func (o Onb) localScalars(a, b, c float64) Vec3 {
	return o.u().multiplyScalar(a).add(o.v().multiplyScalar(b)).add(o.w().multiplyScalar(c))
}

func (o Onb) local(v Vec3) Vec3 {
	return o.u().multiplyScalar(v.x()).add(o.v().multiplyScalar(v.y()).add(o.w().multiplyScalar(v.z())))
}

func (o *Onb) buildFromW(n Vec3) {
	o.axis[2] = n.unitVector()

	var a Vec3

	if math.Abs(o.w().x()) > 0.9 {
		a = Vec3{0, 1, 0}
	} else {
		a = Vec3{1, 0, 0}
	}

	o.axis[1] = o.w().cross(a).unitVector()
	o.axis[0] = o.w().cross(o.v())
}
