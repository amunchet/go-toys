package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i := 0; i< 3; i++{
		fmt.Println(from, ":", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go f("TSDSDDD")

	go func(msg string){
		fmt.Println(msg)
	}("TES")

	time.Sleep(time.Second)

	fmt.Println("Done")
}