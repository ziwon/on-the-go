package bytecounter

import (
	"fmt"
	"testing"
)

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	c.Write([]byte("hello"))

	if c != 5 {
		t.Error("Ooos...")
	}

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)

	if c != 12 {
		t.Error("Shit...")
	}
}
