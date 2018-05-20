package main

import "math"

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
