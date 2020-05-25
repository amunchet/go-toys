package main

import (
	"fmt"	
)

// Creation of interface type
type geometry interface{
	area() float64
	perim() float64
}

// Struct for the interface to be implemented
type rect struct{
	width, height float64
}

func (r rect) area() float64{
	return r.width * r.height
}

func (r rect) perim() float64{
	return r.width + r.height
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r:=rect{3,4}
	measure(r)
}