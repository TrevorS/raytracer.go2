package main

// Material represents different materials hitable objects can be made from.
type Material interface {
	scatter(rayIn Ray, hit Hit) (didScatter bool, attenuation Vec3, scattered Ray)
}

// Lambertian is a diffuse Material.
type Lambertian struct {
	albedo Vec3
}

// Metal is a reflective Material.
type Metal struct {
	albedo Vec3
}

func (l Lambertian) scatter(rayIn Ray, hit Hit) (didScatter bool, attenuation Vec3, scattered Ray) {
	target := hit.p.add(hit.normal).add(RandomInUnitSphere())

	// We could only scatter with some probability and divide albedo by the probability.
	scattered = Ray{hit.p, target.subtract(hit.p)}
	didScatter = true
	attenuation = l.albedo

	return
}

func (m Metal) scatter(rayIn Ray, hit Hit) (didScatter bool, attenuation Vec3, scattered Ray) {
	reflected := rayIn.direction().unitVector().reflect(hit.normal)

	scattered = Ray{hit.p, reflected}
	didScatter = scattered.direction().dot(hit.normal) > 0
	attenuation = m.albedo

	return
}
