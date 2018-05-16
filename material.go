package main

import (
	"math/rand"
)

// Material represents different materials hitable objects can be made from.
type Material interface {
	scatter(rayIn Ray, hit Hit) (didScatter bool, attenuation Vec3, scattered Ray)
}

// Lambertian is a diffuse Material.
type Lambertian struct {
	albedo Vec3
}

// NewLambertian returns a Lambertian material.
func NewLambertian(albedo Vec3) Lambertian {
	return Lambertian{albedo}
}

func (l Lambertian) scatter(rayIn Ray, hit Hit) (didScatter bool, attenuation Vec3, scattered Ray) {
	target := hit.p.add(hit.normal).add(RandomInUnitSphere())

	// We could only scatter with some probability and divide albedo by the probability.
	scattered = Ray{hit.p, target.subtract(hit.p), rayIn.time()}
	didScatter = true
	attenuation = l.albedo

	return
}

// Metal is a reflective Material.
type Metal struct {
	albedo Vec3
	fuzz   float64
}

// NewMetal returns a Metal with fuzz normalized.
func NewMetal(albedo Vec3, fuzz float64) Metal {
	var f float64

	if fuzz < 1 {
		f = fuzz
	} else {
		f = 1
	}

	return Metal{albedo, f}
}

func (m Metal) scatter(rayIn Ray, hit Hit) (didScatter bool, attenuation Vec3, scattered Ray) {
	reflected := rayIn.direction().unitVector().reflect(hit.normal)

	scattered = Ray{hit.p, reflected.add(RandomInUnitSphere().multiplyScalar(m.fuzz)), rayIn.time()}
	didScatter = scattered.direction().dot(hit.normal) > 0
	attenuation = m.albedo

	return
}

// Dielectric is a material that refracts.
type Dielectric struct {
	reflectiveIndex float64
}

// NewDielectric returns a Dielectric material.
func NewDielectric(reflectiveIndex float64) Dielectric {
	return Dielectric{reflectiveIndex}
}

func (d Dielectric) scatter(rayIn Ray, hit Hit) (didScatter bool, attenuation Vec3, scattered Ray) {
	didScatter = true
	var outwardNormal Vec3
	var niOverNt float64

	reflected := rayIn.direction().unitVector().reflect(hit.normal)

	attenuation = Vec3{1.0, 1.0, 1.0}

	var reflectProb float64
	var cosine float64

	if rayIn.direction().dot(hit.normal) > 0 {
		outwardNormal = hit.normal.negate()
		niOverNt = d.reflectiveIndex
		cosine = d.reflectiveIndex * rayIn.direction().dot(hit.normal) / rayIn.direction().length()
	} else {
		outwardNormal = hit.normal
		niOverNt = 1.0 / d.reflectiveIndex
		cosine = -1.0 * rayIn.direction().dot(hit.normal) / rayIn.direction().length()
	}

	didRefract, refracted := rayIn.direction().refract(outwardNormal, niOverNt)

	if didRefract {
		reflectProb = Schlick(cosine, d.reflectiveIndex)
	} else {
		reflectProb = 1.0
	}

	if rand.Float64() < reflectProb {
		scattered = Ray{hit.p, reflected, rayIn.time()}
	} else {
		scattered = Ray{hit.p, *refracted, rayIn.time()}
	}

	return
}
