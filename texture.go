package main

import (
	"image"
	"math"
)

// Texture represents a programmatic way of determining the color of a point.
type Texture interface {
	value(u, v float64, p Vec3) Vec3
}

// ConstantTexture is the same color at any point.
type ConstantTexture struct {
	color Vec3
}

func (ct ConstantTexture) value(u, v float64, p Vec3) Vec3 {
	return ct.color
}

// CheckerTexture is a Texture that alternates colors.
type CheckerTexture struct {
	odd  Texture
	even Texture
}

func (ct CheckerTexture) value(u, v float64, p Vec3) Vec3 {
	sines := math.Sin(10*p.x()) * math.Sin(10*p.y()) * math.Sin(10*p.z())

	if sines < 0 {
		return ct.odd.value(u, v, p)
	}

	return ct.even.value(u, v, p)
}

// NoiseTexture is a Texture with Perlin noise.
type NoiseTexture struct {
	noise Perlin
	scale float64
}

// NewNoiseTexture returns a properly initialized NoiseTexture.
func NewNoiseTexture(scale float64) NoiseTexture {
	return NoiseTexture{
		noise: NewPerlin(),
		scale: scale,
	}
}

func (nt NoiseTexture) value(u, v float64, p Vec3) Vec3 {
	turbulenceDepth := 7

	return Vec3{1, 1, 1}.multiplyScalar(nt.noise.turbulence(p.multiplyScalar(nt.scale), turbulenceDepth))
}

// MarbleTexture generates a marbled Texture using Perlin noise.
type MarbleTexture NoiseTexture

// NewMarbleTexture correctly generates a new MarbleTexture.
func NewMarbleTexture(scale float64) MarbleTexture {
	return MarbleTexture{
		noise: NewPerlin(),
		scale: scale,
	}
}

func (mt MarbleTexture) value(u, v float64, p Vec3) Vec3 {
	turbulenceDepth := 7

	return Vec3{1, 1, 1}.multiplyScalar(0.5).multiplyScalar(1 + math.Sin(mt.scale*p.z()+10*mt.noise.turbulence(p, turbulenceDepth)))
}

// ImageTexture uses an Image as a Texture.
type ImageTexture struct {
	data image.Image
	nx   int
	ny   int
}

// NewImageTexture correctly instantiates an ImageTexture.
func NewImageTexture(data image.Image) ImageTexture {
	return ImageTexture{
		data: data,
		nx:   data.Bounds().Max.X,
		ny:   data.Bounds().Max.Y,
	}
}

func (it ImageTexture) value(u, v float64, p Vec3) Vec3 {
	i := int(u * float64(it.nx))
	j := int((1-v)*float64(it.ny) - 0.001)

	if i < 0 {
		i = 0
	}

	if j < 0 {
		j = 0
	}

	if i > it.nx-1 {
		i = it.nx - 1
	}

	if j > it.ny-1 {
		j = it.ny - 1
	}

	r, g, b, _ := it.data.At(i, j).RGBA()

	normR := float64(r) / 65536
	normG := float64(g) / 65536
	normB := float64(b) / 65536

	return Vec3{normR, normG, normB}
}
