package main

import (
	"math"
	"math/rand"
)

// Pdf represents a probability distribution function.
type Pdf interface {
	value(direction Vec3) float64
	generate() Vec3
}

// CosinePdf is a cosine version of a PDF.
type CosinePdf struct {
	uvw Onb
}

// NewCosinePdf correctly instantiates a CosinePdf.
func NewCosinePdf(w Vec3) CosinePdf {
	uvw := Onb{}
	uvw.buildFromW(w)

	return CosinePdf{uvw}
}

func (cpdf CosinePdf) value(direction Vec3) float64 {
	cosine := direction.unitVector().dot(cpdf.uvw.w())

	if cosine > 0 {
		return cosine / math.Pi
	}

	return 0
}

func (cpdf CosinePdf) generate() Vec3 {
	return cpdf.uvw.local(RandomCosineDirection())
}

// HitablePdf represents a PDF that uses a Hitable object.
type HitablePdf struct {
	hitable Hitable
	o       Vec3
}

func (hPdf HitablePdf) value(direction Vec3) float64 {
	return hPdf.hitable.pdfValue(hPdf.o, direction)
}

func (hPdf HitablePdf) generate() Vec3 {
	return hPdf.hitable.random(hPdf.o)
}

// MixturePdf is a combination of two Pdfs.
type MixturePdf struct {
	pdfs [2]Pdf
}

// NewMixturePdf correctly initializes a MixturePdf.
func NewMixturePdf(p0, p1 Pdf) MixturePdf {
	return MixturePdf{
		[2]Pdf{p0, p1},
	}
}

func (mPdf MixturePdf) value(direction Vec3) float64 {
	return 0.5*mPdf.pdfs[0].value(direction) + 0.5*mPdf.pdfs[1].value(direction)
}

func (mPdf MixturePdf) generate() Vec3 {
	if rand.Float64() < 0.5 {
		return mPdf.pdfs[0].generate()
	}

	return mPdf.pdfs[1].generate()
}
