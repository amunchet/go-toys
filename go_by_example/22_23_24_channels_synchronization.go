package main

import (
	"fmt"
	"time"
)

// Function to sync
func worker (done chan bool){
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Print("done")

	done <- true
}

// Channel directions
func ping(pings chan<-string, msg string){
	pings <- msg
}

func pong(pings <- chan string, pongs chan <- string){
	msg:=<-pings
	pongs <- msg
}

func main() {
	msg := make(chan string, 10)

	go func(){
		for i:=0;i<12;i++{
		msg <- fmt.Sprintf("ping - %d", i)
		}
	}()

		go func(){
	for i:=0;i<10;i++{
		fmt.Println(<-msg)
	}
}()

time.Sleep(time.Second * 1)

// Listening channel to sync
done := make(chan bool, 1)

go worker(done)

<- done

// Channel Directions
pings := make(chan string,1)
pongs := make(chan string,1)

ping(pings, "hello there")
pong(pings, pongs)
fmt.Println(<-pongs)

}