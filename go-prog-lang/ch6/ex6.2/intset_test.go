package intset

import "testing"
import "fmt"

func TestIntBit(t *testing.T) {
	var x, y, z IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)

	if x.Len() != 3 {
		t.Error("Failed:", x.Len())
	}

	y.Add(9)
	y.Add(42)

	x.UnionWith(&y)
	z.AddAll(1, 9, 42, 144)

	if x.Len() != z.Len() {
		t.Error("Failed")
	}

	if x.Len() != 4 {
		t.Error("Failed:", x.Len())
	}

	if x.Has(9) != true {
		t.Error("Failed")
	}

	if x.Has(42) != true {
		t.Error("Failed")
	}

	fmt.Println(&x)
	fmt.Println(x.String())
	fmt.Println(x)
}
