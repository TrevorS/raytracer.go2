package main

import (
	"math/rand"
)

// Perlin represents a Perlin noise generator.
type Perlin struct {
	randFloat []float64
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
	i := int(4*p.x()) & 255
	j := int(4*p.y()) & 255
	k := int(4*p.z()) & 255

	index := perlin.permX[i] ^ perlin.permY[j] ^ perlin.permZ[k]

	return perlin.randFloat[index]
}

func perlinGenerate() []float64 {
	var p []float64

	for i := 0; i < 256; i++ {
		p = append(p, rand.Float64())
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
