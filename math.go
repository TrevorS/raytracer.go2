package main

import (
	"math"
	"math/rand"
)

// RandomInUnitSphere returns a random Vector within the unit sphere.
func RandomInUnitSphere() Vec3 {
	for {
		p := Vec3{
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		}.multiplyScalar(2.0).subtract(Vec3{1.0, 1.0, 1.0})

		if p.squaredLength() < 1.0 {
			return p
		}
	}
}

// RandomToSphere :)
func RandomToSphere(radius, distanceSquared float64) Vec3 {
	r1 := rand.Float64()
	r2 := rand.Float64()
	phi := 2 * math.Pi * r1

	z := 1 + r2*(math.Sqrt(1-radius*radius/distanceSquared)-1)

	x := math.Cos(phi) * math.Sqrt(1-z*z)
	y := math.Sin(phi) * math.Sqrt(1-z*z)

	return Vec3{x, y, z}
}

// RandomInUnitDisk returns a random Vector within the unit disk.
func RandomInUnitDisk() Vec3 {
	for {
		p := Vec3{
			rand.Float64(),
			rand.Float64(),
			0,
		}.multiplyScalar(2.0).subtract(Vec3{1.0, 1.0, 0})

		if p.dot(p) < 1.0 {
			return p
		}
	}
}

// Schlick calculates an approximation of reflectivity varied by angle.
func Schlick(cosine, refractiveIndex float64) float64 {
	r0 := (1 - refractiveIndex) / (1 + refractiveIndex)

	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

// PerlinTriLinearInterpolation performs Perlin tri-linear interpolation. :)
func PerlinTriLinearInterpolation(c [2][2][2]Vec3, u, v, w float64) float64 {
	uu := u * u * (3 - 2*u)
	vv := v * v * (3 - 2*v)
	ww := w * w * (3 - 2*w)

	acc := 0.0

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				weightV := Vec3{
					u - float64(i),
					v - float64(j),
					w - float64(k),
				}

				acc += (float64(i)*uu + (1-float64(i))*(1-uu)) *
					(float64(j)*vv + (1-float64(j))*(1-vv)) *
					(float64(k)*ww + (1-float64(k))*(1-ww)) *
					c[i][j][k].dot(weightV)
			}
		}
	}

	return acc
}

// GetSphereUV returns U and V for a Sphere for doing image textures.
func GetSphereUV(p Vec3) (u, v float64) {
	phi := math.Atan2(p.z(), p.x())
	theta := math.Asin(p.y())

	u = 1 - (phi+math.Pi)/(2*math.Pi)
	v = (theta + math.Pi/2) / math.Pi

	return
}

// RandomCosineDirection generates a random cosine direction as a Vec3.
func RandomCosineDirection() Vec3 {
	r1 := rand.Float64()
	r2 := rand.Float64()

	phi := 2 * math.Pi * r1

	x := math.Cos(phi) * 2 * math.Sqrt(r2)
	y := math.Sin(phi) * 2 * math.Sqrt(r2)
	z := math.Sqrt(1 - r2)

	return Vec3{x, y, z}
}
