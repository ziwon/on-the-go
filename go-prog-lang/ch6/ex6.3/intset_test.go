package intset

import "testing"
import "fmt"

func TestIntBit(t *testing.T) {
	var x, y IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(144)
	y.Add(8)
	y.Add(3)

	x.IntersectWith(&y)
	fmt.Println("===========")
	fmt.Println(x.String())
}
