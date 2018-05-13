package main

// Config holds settings for preparing the rendering engine.
type Config struct {
	width    int
	height   int
	samples  int
	from     Vec3
	at       Vec3
	up       Vec3
	fov      float64
	aperture float64
}

func (c Config) aspectRatio() float64 {
	return float64(c.width) / float64(c.height)
}

func (c Config) focusDistance() float64 {
	return c.from.subtract(c.at).length()
}
