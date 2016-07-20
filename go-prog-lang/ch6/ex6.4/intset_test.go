package intset

import "testing"

func TestElems(t *testing.T) {
	y := []int{1, 2, 3, 144, 4095, 8129, 1048575}

	var x IntSet
	for _, e := range y {
		x.Add(e)
	}

	for i, e := range x.Elems() {
		if e != y[i] {
			t.Error("Oops...")
		}
	}
}
