package main

import (
	"math"
	"math/rand"
)

// Perlin represents a Perlin noise generator.
type Perlin struct {
	randFloat []Vec3
	permX     []int
	permY     []int
	permZ     []int
}

// NewPerlin is a factory for Perlin instances.
func NewPerlin() Perlin {
	return Perlin{
		randFloat: perlinGenerate(),
		permX:     perlinGeneratePerm(),
		permY:     perlinGeneratePerm(),
		permZ:     perlinGeneratePerm(),
	}
}

func (perlin Perlin) noise(p Vec3) float64 {
	u := p.x() - math.Floor(p.x())
	v := p.y() - math.Floor(p.y())
	w := p.z() - math.Floor(p.z())

	i := int(math.Floor(p.x()))
	j := int(math.Floor(p.y()))
	k := int(math.Floor(p.z()))

	c := [2][2][2]Vec3{}

	for di := 0; di < 2; di++ {
		for dj := 0; dj < 2; dj++ {
			for dk := 0; dk < 2; dk++ {
				index := perlin.permX[(i+di)&255] ^
					perlin.permY[(j+dj)&255] ^
					perlin.permZ[(k+dk)&255]

				c[di][dj][dk] = perlin.randFloat[index]
			}
		}
	}

	return PerlinTriLinearInterpolation(c, u, v, w)
}

func (perlin Perlin) turbulence(p Vec3, depth int) float64 {
	acc := 0.0
	weight := 1.0

	tempP := p

	for i := 0; i < depth; i++ {
		acc += weight * perlin.noise(tempP)
		weight *= 0.5
		tempP.inPlaceMultiplyScalar(2)
	}

	return math.Abs(acc)
}

func perlinGenerate() []Vec3 {
	var p []Vec3

	for i := 0; i < 256; i++ {
		vec := Vec3{
			-1 + 2*rand.Float64(),
			-1 + 2*rand.Float64(),
			-1 + 2*rand.Float64(),
		}.unitVector()

		p = append(p, vec)
	}

	return p
}

func permute(p *[]int) {
	for i := len(*p) - 1; i > 0; i-- {
		target := int(rand.Float64() * float64(i+1))
		(*p)[i], (*p)[target] = (*p)[target], (*p)[i]
	}
}

func perlinGeneratePerm() []int {
	var p []int

	for i := 0; i < 256; i++ {
		p = append(p, i)
	}

	permute(&p)

	return p
}
