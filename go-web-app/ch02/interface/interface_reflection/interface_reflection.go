package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	p := reflect.ValueOf(&x)
	u := p.Elem()
	u.SetFloat(7.1)
	fmt.Println("value:", u)

	var y uint8 = 'y'
	w := reflect.ValueOf(y)
	fmt.Println("type:", w.Type())
	fmt.Println("kind is uinit8:", w.Kind() == reflect.Uint8)
	fmt.Println("value:", uint8(w.Uint()))
}
