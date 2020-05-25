package main

import (
	"fmt"
	"time"
)

func main() {
	
	// Normal Timer

	timer1 := time.NewTimer(time.Second * 2)

	<- timer1.C

	fmt.Println("Timer 1 fired")

	// Cancelling a timer
	timer2 := time.NewTimer(time.Second * 5)
	go func(){
		<- timer2.C
		fmt.Println("Timer 2 completed")
	}()

	stop2 := timer2.Stop()
	if stop2{
		fmt.Println("Timer 2 was stopped intentionally")
	}

	time.Sleep(2 * time.Second)

	// Tickers

	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	go func(){
		for{
			select{
			case <- done:
				return
			case t := <- ticker.C:
				fmt.Println("Tick at ", t)
			}
		}
	}()


	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("All done!")

}