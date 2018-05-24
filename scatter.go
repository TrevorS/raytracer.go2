package main

// Scatter represents a scatter record.
type Scatter struct {
	specularRay Ray
	isSpecular  bool
	attenuation Vec3
	pdf         Pdf
}
