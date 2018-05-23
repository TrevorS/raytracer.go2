package main

// HitableList is an array of Hitable graphics objects.
type HitableList []Hitable

// SortByX implements sort.Interface for HitableList.
type SortByX HitableList

// Len is required to implement sort.Interface for SortByX.
func (list SortByX) Len() int {
	return len(list)
}

// Swap is required to implement sort.Interface for SortByX.
func (list SortByX) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// Less is required to implement sort.Interface for SortByX.
func (list SortByX) Less(i, j int) bool {
	iHitable := list[i]
	jHitable := list[j]

	iHasBox, iBox := iHitable.boundingBox(0, 0)
	jHasBox, jBox := jHitable.boundingBox(0, 0)

	if !iHasBox || !jHasBox {
		panic("No bounding box in BVHNode constructor")
	}

	return iBox.min.x() < jBox.min.x()
}

// SortByY implements sort.Interface for HitableList.
type SortByY HitableList

// Len is required to implement sort.Interface for SortByY.
func (list SortByY) Len() int {
	return len(list)
}

// Swap is required to implement sort.Interface for SortByY.
func (list SortByY) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// Less is required to implement sort.Interface for SortByY.
func (list SortByY) Less(i, j int) bool {
	iHitable := list[i]
	jHitable := list[j]

	iHasBox, iBox := iHitable.boundingBox(0, 0)
	jHasBox, jBox := jHitable.boundingBox(0, 0)

	if !iHasBox || !jHasBox {
		panic("No bounding box in BVHNode constructor")
	}

	return iBox.min.y() < jBox.min.y()
}

// SortByZ implements sort.Interface for HitableList.
type SortByZ HitableList

// Len is required to implement sort.Interface for SortByZ.
func (list SortByZ) Len() int {
	return len(list)
}

// Swap is required to implement sort.Interface for SortByZ.
func (list SortByZ) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// Less is required to implement sort.Interface for SortByZ.
func (list SortByZ) Less(i, j int) bool {
	iHitable := list[i]
	jHitable := list[j]

	iHasBox, iBox := iHitable.boundingBox(0, 0)
	jHasBox, jBox := jHitable.boundingBox(0, 0)

	if !iHasBox || !jHasBox {
		panic("No bounding box in BVHNode constructor")
	}

	return iBox.min.z() < jBox.min.z()
}

// NewHitableList returns a dynamically sized HitableList.
func NewHitableList(size int) HitableList {
	return make(HitableList, size)
}

func (hList *HitableList) add(hitable Hitable) HitableList {
	*hList = append(*hList, hitable)

	return *hList
}

func (hList HitableList) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	hitAnything := false
	closest := tMax

	var closestHit Hit

	for _, hitable := range hList {
		didHit, hit := hitable.hit(r, tMin, closest)

		if didHit {
			hitAnything = true
			closest = hit.t
			closestHit = *hit
		}
	}

	return hitAnything, &closestHit
}

func (hList HitableList) boundingBox(t0, t1 float64) (hasBox bool, box *AABB) {
	if len(hList) < 1 {
		return false, nil
	}

	doesHaveBox, firstBox := hList[0].boundingBox(t0, t1)

	if !doesHaveBox {
		return false, nil
	}

	box = firstBox

	for i := 1; i < len(hList); i++ {
		doesHaveNextBox, nextBox := hList[i].boundingBox(t0, t1)

		if doesHaveNextBox {
			box = SurroundingBox(*box, *nextBox)
		} else {
			return false, nil
		}
	}

	return true, box
}

func (hList HitableList) pdfValue(o, direction Vec3) float64 {
	return 0.0
}

func (hList HitableList) random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
