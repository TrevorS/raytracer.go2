package main

import (
	"math"
	"math/rand"
)

// Material represents different materials hitable objects can be made from.
type Material interface {
	scatter(rayIn Ray, hit Hit) (didScatter bool, albedo Vec3, scattered Ray, pdf float64)
	scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64
	emitted(u, v float64, p Vec3) Vec3
}

// Lambertian is a diffuse Material.
type Lambertian struct {
	albedo Texture
}

// NewLambertian returns a Lambertian material.
func NewLambertian(albedo Texture) Lambertian {
	return Lambertian{albedo}
}

func (l Lambertian) scatter(rayIn Ray, hit Hit) (didScatter bool, albedo Vec3, scattered Ray, pdf float64) {
	var uvw Onb
	uvw.buildFromW(hit.normal)

	direction := uvw.local(RandomCosineDirection())

	scattered = Ray{hit.p, direction.unitVector(), rayIn.time()}
	albedo = l.albedo.value(hit.u, hit.v, hit.p)
	pdf = uvw.w().dot(scattered.direction()) / math.Pi
	didScatter = true

	return
}

func (l Lambertian) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	cosine := hit.normal.dot(scattered.direction().unitVector())

	if cosine < 0 {
		cosine = 0
	}

	return cosine / math.Pi
}

func (l Lambertian) emitted(u, v float64, p Vec3) Vec3 {
	return EmitBlack()
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

func (m Metal) scatter(rayIn Ray, hit Hit) (didScatter bool, albedo Vec3, scattered Ray, pdf float64) {
	reflected := rayIn.direction().unitVector().reflect(hit.normal)

	scattered = Ray{hit.p, reflected.add(RandomInUnitSphere().multiplyScalar(m.fuzz)), rayIn.time()}
	didScatter = scattered.direction().dot(hit.normal) > 0
	albedo = m.albedo

	return
}

func (m Metal) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	return 0
}

func (m Metal) emitted(u, v float64, p Vec3) Vec3 {
	return EmitBlack()
}

// Dielectric is a material that refracts.
type Dielectric struct {
	reflectiveIndex float64
}

// NewDielectric returns a Dielectric material.
func NewDielectric(reflectiveIndex float64) Dielectric {
	return Dielectric{reflectiveIndex}
}

func (d Dielectric) scatter(rayIn Ray, hit Hit) (didScatter bool, albedo Vec3, scattered Ray, pdf float64) {
	didScatter = true
	var outwardNormal Vec3
	var niOverNt float64

	reflected := rayIn.direction().unitVector().reflect(hit.normal)

	albedo = Vec3{1.0, 1.0, 1.0}

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

func (d Dielectric) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	return 0
}

func (d Dielectric) emitted(u, v float64, p Vec3) Vec3 {
	return EmitBlack()
}

// DiffuseLight is a material that acts as a diffused light.
type DiffuseLight struct {
	emit Texture
}

func (dl DiffuseLight) scatter(rayIn Ray, hit Hit) (didScatter bool, albedo Vec3, scattered Ray, pdf float64) {
	return false, Vec3{}, Ray{}, 0.0
}

func (dl DiffuseLight) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	return 0
}

func (dl DiffuseLight) emitted(u, v float64, p Vec3) Vec3 {
	return dl.emit.value(u, v, p)
}

// Isotropic has a scattering function that picks a uniform random direction.
type Isotropic struct {
	albedo Texture
}

func (it Isotropic) scatter(rayIn Ray, hit Hit) (didScatter bool, albedo Vec3, scattered Ray, pdf float64) {
	didScatter = true

	scattered = Ray{
		hit.p,
		RandomInUnitSphere(),
		rayIn.time(),
	}

	albedo = it.albedo.value(hit.u, hit.v, hit.p)

	return
}

func (it Isotropic) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	return 0
}

func (it Isotropic) emitted(u, v float64, p Vec3) Vec3 {
	return EmitBlack()
}
