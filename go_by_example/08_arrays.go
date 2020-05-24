package main

import (
	"fmt"
)

func main() {
	var a [5] int //Array of exactly 5 ints
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a[4])

	var b [2][4] int

	fmt.Println(b)
	fmt.Printf("%t", b)
}