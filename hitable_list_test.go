package main

import (
	"sort"
	"testing"
)

type mockHitable struct {
	x, y, z float64
}

func (mh mockHitable) hit(r Ray, tMin, tMax float64) (bool, *Hit) {
	return false, nil
}

func (mh mockHitable) boundingBox(t0, t1 float64) (bool, *AABB) {
	return true, &AABB{
		Vec3{
			mh.x,
			mh.y,
			mh.z,
		},
		Vec3{
			mh.x,
			mh.y,
			mh.z,
		},
	}
}

func testOrder(actual, expected HitableList, t *testing.T) {
	for i := 0; i < len(actual); i++ {
		a := actual[i]
		e := expected[i]

		if a != e {
			t.Errorf("Sort failed: %v != %v", a, e)
		}
	}
}

func TestSortByX(t *testing.T) {
	low := mockHitable{0, 9, 9}
	mid := mockHitable{1, 9, 9}
	high := mockHitable{2, 9, 9}

	actual := HitableList{mid, high, low}
	expected := HitableList{low, mid, high}

	sort.Sort(SortByX(actual))

	testOrder(actual, expected, t)
}

func TestSortByY(t *testing.T) {
	low := mockHitable{9, 2, 9}
	mid := mockHitable{9, 4, 9}
	high := mockHitable{9, 6, 9}

	actual := HitableList{mid, high, low}
	expected := HitableList{low, mid, high}

	sort.Sort(SortByY(actual))

	testOrder(actual, expected, t)
}

func TestSortByZ(t *testing.T) {
	low := mockHitable{9, 9, 3}
	mid := mockHitable{9, 9, 6}
	high := mockHitable{9, 9, 9}

	actual := HitableList{mid, high, low}
	expected := HitableList{low, mid, high}

	sort.Sort(SortByZ(actual))

	testOrder(actual, expected, t)
}
