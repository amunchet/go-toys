package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for i := 0; i < 12; i++ {

			c1 <- fmt.Sprintf("C1: %d", i)

			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 12; i++ {

			c2 <- fmt.Sprintf("C2: %d", i)

			time.Sleep(time.Millisecond * 1)
		}
	}()

	for i := 0; i < 2000; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		default:
			fmt.Print(".")
		}
	}

}
