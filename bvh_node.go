package main

import (
	"math/rand"
	"sort"
)

// BVHNode represents a node in a Bounding Volume Hierarchy
type BVHNode struct {
	left  Hitable
	right Hitable
	box   *AABB
}

func (n *BVHNode) newBVHNode(hList *HitableList, time0, time1 float64) *BVHNode {
	list := *hList
	axis := int(3 * rand.Float64())

	if axis == 0 {
		sort.Sort(SortByX(list))
	} else if axis == 1 {
		sort.Sort(SortByY(list))
	} else if axis == 2 {
		sort.Sort(SortByZ(list))
	} else {
		panic("Unexpected axis!")
	}

	length := len(list)

	if length == 1 {
		n.left = list[0]
		n.right = list[0]
	} else if length == 2 {
		n.left = list[0]
		n.right = list[1]
	} else {
		firstHalf := list[:length/2]
		secondHalf := list[length/2:]

		n.left = *n.newBVHNode(&firstHalf, time0, time1)
		n.right = *n.newBVHNode(&secondHalf, time0, time1)
	}

	hasLeftBox, leftBox := n.left.boundingBox(time0, time1)
	hasRightBox, rightBox := n.right.boundingBox(time0, time1)

	if !hasLeftBox || !hasRightBox {
		panic("No BoundingBox in BVHNode constructor")
	}

	n.box = SurroundingBox(*leftBox, *rightBox)

	return n
}

func (n BVHNode) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	didHit := n.box.hit(r, tMin, tMax)

	if didHit {
		left := n.left
		right := n.right

		didHitLeft, leftHit := left.hit(r, tMin, tMax)
		didHitRight, rightHit := right.hit(r, tMin, tMax)

		if didHitLeft && didHitRight {
			if leftHit.t < rightHit.t {
				return true, leftHit
			}

			return true, rightHit
		}

		if didHitLeft {
			return true, leftHit
		}

		if didHitRight {
			return true, rightHit
		}

		return false, nil
	}

	return false, nil
}

func (n BVHNode) boundingBox(t0, t1 float64) (bool, *AABB) {
	return true, n.box
}
