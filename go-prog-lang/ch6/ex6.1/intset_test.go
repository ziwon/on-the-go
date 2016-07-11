package intset

import "testing"
import "fmt"

func TestIntBit(t *testing.T) {
	var x, y IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)

	if x.Len() != 3 {
		t.Error("Failed:", x.Len())
	}

	y.Add(9)
	y.Add(42)

	x.UnionWith(&y)

	if x.Len() != 4 {
		t.Error("Failed:", x.Len())
	}

	if x.Has(9) != true {
		t.Error("Error")
	}

	if x.Has(42) != true {
		t.Error("Error")
	}

	fmt.Println(&x)
	fmt.Println(x.String())
	fmt.Println(x)
}
