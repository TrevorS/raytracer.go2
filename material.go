package main

import (
	"math"
)

// Material represents different materials hitable objects can be made from.
type Material interface {
	scatter(rayIn Ray, hit Hit) (didScatter bool, scatter Scatter)
	scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64
	emitted(rayIn Ray, hit Hit, u, v float64, p Vec3) Vec3
}

// MaterialZero represents a blank material.
type MaterialZero struct {
}

func (mz MaterialZero) scatter(rayIn Ray, hit Hit) (didScatter bool, scatter Scatter) {
	return false, Scatter{}
}

func (mz MaterialZero) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	return 0.0
}

func (mz MaterialZero) emitted(rayIn Ray, hit Hit, u, v float64, p Vec3) Vec3 {
	return Vec3Zero()
}

// Lambertian is a diffuse Material.
type Lambertian struct {
	albedo Texture
}

// NewLambertian returns a Lambertian material.
func NewLambertian(albedo Texture) Lambertian {
	return Lambertian{albedo}
}

func (l Lambertian) scatter(rayIn Ray, hit Hit) (didScatter bool, scatter Scatter) {
	isSpecular := false
	attenuation := l.albedo.value(hit.u, hit.v, hit.p)
	pdf := NewCosinePdf(hit.normal)

	return true, Scatter{Ray{}, isSpecular, attenuation, pdf}
}

func (l Lambertian) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	cosine := hit.normal.dot(scattered.direction().unitVector())

	if cosine < 0 {
		cosine = 0
	}

	return cosine / math.Pi
}

func (l Lambertian) emitted(rayIn Ray, hit Hit, u, v float64, p Vec3) Vec3 {
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

func (m Metal) scatter(rayIn Ray, hit Hit) (didScatter bool, scatter Scatter) {
	reflected := rayIn.direction().unitVector().reflect(hit.normal)

	specularRay := Ray{hit.p, reflected.add(RandomInUnitSphere().multiplyScalar(m.fuzz)), rayIn.time()}
	isSpecular := true
	attenuation := m.albedo
	pdf := PdfZero{}

	return true, Scatter{specularRay, isSpecular, attenuation, pdf}
}

func (m Metal) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	return 0
}

func (m Metal) emitted(rayIn Ray, hit Hit, u, v float64, p Vec3) Vec3 {
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

func (d Dielectric) scatter(rayIn Ray, hit Hit) (didScatter bool, scatter Scatter) {
	return true, Scatter{Ray{}, true, Vec3{}, NewCosinePdf(Vec3Zero())}
}

func (d Dielectric) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	return 0
}

func (d Dielectric) emitted(rayIn Ray, hit Hit, u, v float64, p Vec3) Vec3 {
	return EmitBlack()
}

// DiffuseLight is a material that acts as a diffused light.
type DiffuseLight struct {
	emit Texture
}

func (dl DiffuseLight) scatter(rayIn Ray, hit Hit) (didScatter bool, scatter Scatter) {
	return false, Scatter{Ray{}, true, Vec3{}, NewCosinePdf(Vec3Zero())}
}

func (dl DiffuseLight) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	return 0
}

func (dl DiffuseLight) emitted(rayIn Ray, hit Hit, u, v float64, p Vec3) Vec3 {
	if hit.normal.dot(rayIn.direction()) < 0.0 {
		return dl.emit.value(u, v, p)
	}

	return EmitBlack()
}

// Isotropic has a scattering function that picks a uniform random direction.
type Isotropic struct {
	albedo Texture
}

func (it Isotropic) scatter(rayIn Ray, hit Hit) (didScatter bool, scatter Scatter) {
	return true, Scatter{Ray{}, true, Vec3{}, NewCosinePdf(Vec3Zero())}
}

func (it Isotropic) scatteringPdf(rayIn Ray, hit Hit, scattered Ray) float64 {
	return 0
}

func (it Isotropic) emitted(rayIn Ray, hit Hit, u, v float64, p Vec3) Vec3 {
	return EmitBlack()
}
