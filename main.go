package main

func main() {
	config := Config{
		width:     640,
		height:    480,
		samples:   100,
		from:      Vec3{278, 278, -800},
		at:        Vec3{278, 278, 0},
		up:        Vec3{0, 1, 0},
		fov:       40.0,
		aperture:  0.00,
		filename:  "output.png",
		timeStart: 0,
		timeEnd:   1,
	}

	world := CornellSmoke(config)

	camera := NewCamera(
		config.from,
		config.at,
		config.up,
		config.fov,
		config.aspectRatio(),
		config.aperture,
		config.focusDistance(),
		config.timeStart,
		config.timeEnd,
	)

	framebuffer := Render(camera, world, config)

	WriteImage(&framebuffer, config)
}
