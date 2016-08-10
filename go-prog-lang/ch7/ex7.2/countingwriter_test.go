package counter

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	const S1 = "Hello"
	const S2 = "Hello World"

	buffer := bytes.NewBuffer([]byte{})
	c1, count1 := CountingWriter(buffer)
	fmt.Fprintf(c1, S1)

	c2, count2 := CountingWriter(buffer)
	fmt.Fprintf(c2, S2)

	if int64(len([]byte(S1))) != *count1 {
		t.Error("Ooops....", *count1)
	}

	if int64(len([]byte(S2))) != *count2 {
		t.Error("Ooops....", *count2)
	}
}
