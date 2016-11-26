package eval

import (
	"bytes"
	"fmt"
	"strconv"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', -1, 64)
}

func (u unary) String() string {
	return fmt.Sprintf("%s%s", string(u.op), u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("%s %s %s", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
	b := new(bytes.Buffer)
	b.WriteString(c.fn)
	b.WriteString("(")
	for i, a := range c.args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(a.String())
	}
	b.WriteString(")")
	return b.String()
}
