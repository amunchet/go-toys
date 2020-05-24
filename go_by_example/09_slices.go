package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 3)
	slice[0] = "abcddd"
	slice[1] = "adsfas"
	slice[2] = "asdfasdf"

	slice = append(slice, "SDFSDFSDF")

	fmt.Println(slice)
	fmt.Println(len(slice))

	// Copy the slice
	test := make([] string, 1)
	copy(test, slice)

	fmt.Println(test)

	// Print partial - a la Python/Swift
	fmt.Println(slice[1:3])
	
}