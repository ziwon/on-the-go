package intset

import "testing"
import "fmt"

func TestElems(t *testing.T) {
	var x IntSet

	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(244)

	for _, e := range x.Elems() {
		fmt.Println(e)
	}
}
