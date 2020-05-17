package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before lunch")
	default:
		fmt.Println("Afternoon!")
	}
}
