package main

import (
	"math"
)

// Vec3 represents a point or a color.
type Vec3 struct {
	e0 float64
	e1 float64
	e2 float64
}

// Vec3Zero instantiates a zero'd out Vec3.
func Vec3Zero() Vec3 {
	return Vec3{0, 0, 0}
}

func (v Vec3) x() float64 {
	return v.e0
}

func (v Vec3) y() float64 {
	return v.e1
}

func (v Vec3) z() float64 {
	return v.e2
}

func (v Vec3) r() float64 {
	return v.e0
}

func (v Vec3) g() float64 {
	return v.e1
}

func (v Vec3) b() float64 {
	return v.e2
}

func (v Vec3) negate() Vec3 {
	return Vec3{
		-v.e0,
		-v.e1,
		-v.e2,
	}
}

func (v Vec3) length() float64 {
	return math.Sqrt(v.squaredLength())
}

func (v Vec3) squaredLength() float64 {
	return v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2
}

func (v Vec3) unitVector() Vec3 {
	return v.divideScalar(v.length())
}

func (v *Vec3) makeUnitVector() {
	k := 1.0 / math.Sqrt(v.e0*v.e0+v.e1*v.e1+v.e2*v.e2)

	v.e0 *= k
	v.e1 *= k
	v.e2 *= k
}

func (v Vec3) add(v2 Vec3) Vec3 {
	return Vec3{
		v.e0 + v2.e0,
		v.e1 + v2.e1,
		v.e2 + v2.e2,
	}
}

func (v Vec3) subtract(v2 Vec3) Vec3 {
	return Vec3{
		v.e0 - v2.e0,
		v.e1 - v2.e1,
		v.e2 - v2.e2,
	}
}

func (v Vec3) multiply(v2 Vec3) Vec3 {
	return Vec3{
		v.e0 * v2.e0,
		v.e1 * v2.e1,
		v.e2 * v2.e2,
	}
}

func (v Vec3) divide(v2 Vec3) Vec3 {
	return Vec3{
		v.e0 / v2.e0,
		v.e1 / v2.e1,
		v.e2 / v2.e2,
	}
}

func (v Vec3) multiplyScalar(s float64) Vec3 {
	return Vec3{
		v.e0 * s,
		v.e1 * s,
		v.e2 * s,
	}
}

func (v Vec3) divideScalar(s float64) Vec3 {
	return Vec3{
		v.e0 / s,
		v.e1 / s,
		v.e2 / s,
	}
}

func (v Vec3) dot(v2 Vec3) float64 {
	return v.e0*v2.e0 + v.e1*v2.e1 + v.e2*v2.e2
}

func (v Vec3) cross(v2 Vec3) Vec3 {
	return Vec3{
		v.e1*v2.e2 - v.e2*v2.e1,
		-(v.e0*v2.e2 - v.e2*v2.e0),
		v.e0*v2.e1 - v.e1*v2.e0,
	}
}

func (v *Vec3) inPlaceAdd(v2 Vec3) Vec3 {
	v.e0 += v2.e0
	v.e1 += v2.e1
	v.e2 += v2.e2

	return *v
}

func (v *Vec3) inPlaceSubtract(v2 Vec3) Vec3 {
	v.e0 -= v2.e0
	v.e1 -= v2.e1
	v.e2 -= v2.e2

	return *v
}

func (v *Vec3) inPlaceMultiply(v2 Vec3) Vec3 {
	v.e0 *= v2.e0
	v.e1 *= v2.e1
	v.e2 *= v2.e2

	return *v
}

func (v *Vec3) inPlaceDivide(v2 Vec3) Vec3 {
	v.e0 /= v2.e0
	v.e1 /= v2.e1
	v.e2 /= v2.e2

	return *v
}

func (v *Vec3) inPlaceMultiplyScalar(t float64) Vec3 {
	v.e0 *= t
	v.e1 *= t
	v.e2 *= t

	return *v
}

func (v *Vec3) inPlaceDivideScalar(t float64) Vec3 {
	v.e0 /= t
	v.e1 /= t
	v.e2 /= t

	return *v
}

func (v Vec3) reflect(normal Vec3) Vec3 {
	v2 := normal.multiplyScalar(v.dot(normal) * 2)

	return v.subtract(v2)
}
