package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	Sha256 = "256"
	Sha384 = "384"
	Sha512 = "512"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s [text]:\n", os.Args[0])
		fmt.Printf(" \n")
		flag.PrintDefaults()
	}

	var opt = flag.String("sha", "256", "256 | 384 | 512")
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	s := os.Args[len(os.Args)-1]
	if len(os.Args) == 2 {
		fmt.Printf("256 - %x\n", sha256.Sum256([]byte(s)))
		os.Exit(0)
	} else {
		switch *opt {
		case Sha256:
			fmt.Printf("256 - %x\n", sha256.Sum256([]byte(s)))
		case Sha384:
			fmt.Printf("384 - %x\n", sha512.Sum384([]byte(s)))
		case Sha512:
			fmt.Printf("512 - %x\n", sha512.Sum512([]byte(s)))
		}
	}
}
