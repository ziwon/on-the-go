package main

import (
	"fmt"
)

func main() {
	months := [...]string{
		1: "January", 2: "Februrary", 3: "March",
		4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September",
		10: "October", 11: "November", 12: "December"}

	fmt.Println(months)    // array
	fmt.Println(months[:]) // array

	fmt.Println(months[1:])  // slice
	fmt.Println(months[1:6]) // slice

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)
	fmt.Println(summer)

	endlessSummer := summer[:5] // :20 panic
	fmt.Println(endlessSummer)
}
