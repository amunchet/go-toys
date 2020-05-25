package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result"
	}()
	
	func (){
		select {
		case res := <-c1:
			fmt.Println(res)

		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")

		}
	}()
	func (){
		select {
		case res := <-c1:
			fmt.Println(res)

		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")

		}
	}()
	
}
