package intset

import "testing"
import "fmt"

func TestIntersectWith(t *testing.T) {
	var x, y IntSet

	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(244)

	y.Add(3)
	y.Add(4)
	y.Add(244)

	x.IntersectWith(&y)
	if !x.Has(3) || !x.Has(244) || x.Len() != 2 {
		t.Error("Fail")
	} else {
		fmt.Printf("IntersectWith: %s\n", x.String())
	}
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet

	x.Add(1)
	x.Add(2)
	x.Add(300)

	y.Add(300)
	y.Add(4)

	x.DifferenceWith(&y)
	if !x.Has(1) || !x.Has(2) || x.Len() != 2 {
		t.Error("Fail")
	} else {
		fmt.Printf("DifferenceWith: %s\n", x.String())
	}
}

func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet

	x.Add(1)
	x.Add(2)
	x.Add(300)

	y.Add(300)
	y.Add(4)

	x.SymmetricDifference(&y)
	if !x.Has(1) || !x.Has(2) || !x.Has(4) || x.Len() != 3 {
		t.Error("Fail")
	} else {
		fmt.Printf("SymmetricDifference: %s\n", x.String())
	}

}
