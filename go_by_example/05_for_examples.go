package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	count := 5
	for count > 0 {
		fmt.Println(count)
		count -= 1
	}
}
