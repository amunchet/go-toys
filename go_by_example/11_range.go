package main

import (
	"fmt"
)

func main() {
	nums := [] int { 2,3,4 }

	// Over an array
	for _, num := range nums{

		fmt.Println(num)
	}

	// Try a map
	m := map[string]int {"test" : 5, "temp" : 7}

	for k, num := range m {
		fmt.Println(k, num)
	}

	// Try a string
	b := "teststring"

	for k,v := range b{
		fmt.Println(k,v)
		fmt.Printf("%c\n", v)
	}
}