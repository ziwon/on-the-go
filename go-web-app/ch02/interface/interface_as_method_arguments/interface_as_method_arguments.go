package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "Name: " + h.name + ", Age:" + strconv.Itoa(h.age) + " years, Contacts: " + h.phone
}

func main() {
	Bob := Human{"Bob", 40, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
