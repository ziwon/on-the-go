package main

import (
	"fmt"
)

func main() {
	var apples int32 = 1
	var oranges int16 = 2
	// var compote int = apples + oranges not compile

	var compote = int(apples) + int(oranges)
	fmt.Println(compote)

	f := 3.141
	i := int(f)
	fmt.Println(f, i)
	f = 1.99
	fmt.Println(int(f))

	f = 1e100
	i = int(f) // result is implementation-dependent
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdeadbeaf)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '츤'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

	/*
		3
		3.141 3
		1
		438 666 0666
		3735928495 deadbeaf 0xdeadbeaf 0XDEADBEAF
		97 a 'a'
		52772 츤 '츤'
		10 '\n'
	*/
}
