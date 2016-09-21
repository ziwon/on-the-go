package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func Cat(arg string, f *os.File, opt ...int) (string, error) {
	bufSize := 8192

	if len(opt) > 0 && opt[0] > 0 {
		bufSize = opt[0]
	}

	buf := make([]byte, bufSize)

	for {
		switch nr, er := f.Read(buf); {
		case nr > 0:
			if _, ew := os.Stdout.Write(buf[0:nr]); ew != nil {
				return arg, ew
			}
			continue
		case er != nil && er != io.EOF:
			return arg, er
		}
		break
	}
	return arg, nil
}

func main() {
	check := func(arg string, err error) {
		if err != nil {
			if len(arg) > 0 {
				arg = " " + arg + ":"
			}
			fmt.Fprintf(os.Stderr, "cat:%s %s\n", arg, err)
			os.Exit(1)
		}
	}

	bufSize := flag.Int("n", 512, "buffer size")
	flag.Parse()

	if flag.NArg() == 0 {
		check(Cat("", os.Stdin, *bufSize))
	} else {
		for _, arg := range flag.Args() {
			if arg == "-" {
				check(Cat("", os.Stdin, *bufSize))
			} else if len(arg) == 0 {

			} else if f, err := os.Open(arg); err != nil {
				check(arg, err)
			} else {
				check(Cat(arg, f, *bufSize))
				check(arg, f.Close())
			}
		}
	}

	check("output error", os.Stdout.Close())
}
