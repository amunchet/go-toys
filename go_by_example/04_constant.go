package main

import (
	"fmt"
	"math"
)

const s = 43212

func main() {
	fmt.Println(s)

	fmt.Println(math.Sin(s)) // math.Sin wants a float64, so it auto casts
}
