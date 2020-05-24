package main

import (
	"fmt"
)
func test(a, b,c int) int{
	return a+b+c
}

func retTwo(a,b,c int) (int,int){
	return a+b+c, a
}

func variadic(nums...int) int{
	retval := 0
	for _,num := range nums{
		retval += num
	}
	return retval
}

func main() {
	//Normal
	fmt.Println(test(1,2,34))

	// Multiple Returns
	fmt.Println(retTwo(1,22,3))

	// Variadic
	fmt.Println(variadic(2,2,1,1,3,4,4,23,2,1))
	a := [] int { 1,4,3,2,1 }
	fmt.Println(variadic(a...))
}