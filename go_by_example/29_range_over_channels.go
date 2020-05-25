package main

import (
	"fmt"
	"time"
)

func main() {

	// Standard Range over channel
	q := make(chan string, 2)
	q <- "one"
	q <- "two"

	close(q)

	for elem := range q {
		fmt.Println(elem)
	}

	// Try in a go routine

	b := make(chan string,5 )
	go func(){
		for i:=0;i<10;i++{
			b <- fmt.Sprintf("Sending %d", i)
			time.Sleep(time.Millisecond * 500)
		}
		close(b)
	}()

	for r := range b{
		fmt.Println(r)
	}
}