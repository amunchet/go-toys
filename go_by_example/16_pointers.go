package main

import "fmt"

func zeroptr(i *int){
	*i = 0
}

func main() {
	i := 212
	fmt.Println(i)

	zeroptr(&i)

	fmt.Println(i)
	fmt.Println(&i)
}