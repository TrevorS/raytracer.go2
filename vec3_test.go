package main

import "testing"

func TestGet(t *testing.T) {
	expected := []float64{2.0, 4.0, 6.0}

	v := Vec3{
		expected[0],
		expected[1],
		expected[2],
	}

	for e := 0; e < len(expected); e++ {
		actual := v.get(e)
		if actual != expected[e] {
			t.Errorf("did not match, %v != %v", actual, expected[e])
		}
	}
}
