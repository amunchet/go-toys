package main

import (
	"fmt"
)

func IntLess(a, b int) bool {
	return a < b
}

func main() {
	fmt.Println("Hello!")
	fmt.Println("Is 3 less than 4?", IntLess(3, 4))
}
