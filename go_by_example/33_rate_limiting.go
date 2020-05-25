package main

import (
	"fmt"
	"time"
)

func main() {

	// Simple rate limit

	requests := make(chan int, 5)

	for i:=0; i<5; i++{
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<- limiter // Block until receive from limiter - this is where we limit
		fmt.Println("request: ", req, time.Now())

	}

	// Burst rate Limit
	// This is controlled by the Limiter Channel BUFFER

	burstLimiter := make(chan time.Time, 3)

	for i := 0; i<3; i++{
		burstLimiter <- time.Now()
	}

	go func(){
		for t := range time.Tick(200 * time.Millisecond){ // Every 200 milliseconds, try to put another entry into burstLimiter of time Now
			burstLimiter <- t
		}
	}()

	burst_requests := make(chan int, 7)
	for i := 0; i<7; i++{
		burst_requests <- i
	}

	close(burst_requests)

	// So what happens is the first 3 fire instantly, while the remainder are confined to the 200 ms limit, until another time is added to the limiter channel
	for req := range burst_requests {
		fmt.Println(<- burstLimiter) // It's popping out a time
		fmt.Println("Request: ", req, time.Now())
	}


}