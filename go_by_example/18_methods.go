package main

import "fmt"


type measurement struct{
	width, height int
	permimeter int
}

// You can pass in the structure as read only, but most of the time you won't
func (r *measurement) perim(){ 
	r.permimeter = r.width + r.height
	return
}

func main() {
	a:= measurement{15,16, 0}
	fmt.Println(a)

	a.perim()
	a.width = 7

	fmt.Println(a)

	b := &a

	// Go is really nice when it comes to structs and pointers
	// No need to dereference - Go knows what you want
	b.height = 10
	fmt.Println(a)
}