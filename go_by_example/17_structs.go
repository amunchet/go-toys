package main

import "fmt"

type alien struct{
	name string
	age float32

}

func newAlien(name string) *alien{
	a := alien{name:name}
	a.age = 5.3234123
	return &a
}

func main() {
	fmt.Println(alien{"test", 3232.12})

	fmt.Println(newAlien("AlienBob"))

	a := alien{"test", 123.12}
	fmt.Println(a)

	b := &a
	b.name = "TESTSEDTD"

	fmt.Println(a)
	fmt.Println(b)
}