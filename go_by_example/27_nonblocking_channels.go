package main

import "fmt"

func main() {

	// Without a buffer, it fails
	m := make(chan string)

	a := "adfasd"
	select{
	case m <- a:
		fmt.Println("sent message", a)
	default: 
		fmt.Println("No received msg")
	}

	// With a buffer, it sends immediately
	m = make(chan string,1)

	a = "adfasd"
	select{
	case m <- a:
		fmt.Println("sent message", a)
	default: 
		fmt.Println("No received msg")
	}
}