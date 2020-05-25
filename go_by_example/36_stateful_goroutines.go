package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Define our structures - this is basically writing what the Mutex structures do for us
type readOp struct{
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Handle the state

	/* This function confused me a for a bit.  I believe Go is nice and tidies up any orphan goroutines at the end of execution.  Otherwise, I don't see how this would stop running */

	state := make(map[int] int)
	go func(){
		for{
			select{
			case read:= <-reads:
				read.resp <-state[read.key]
			case write:= <-writes:
				state[write.key] = write.val
				write.resp <- true

			}

		}
	}()

	// Read from the data
	for r:=0;r<10; r++{
		go func(){
			for{
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int)}
				reads <- read

				<- read.resp
				atomic.AddUint64(&readOps,1)
				time.Sleep(time.Millisecond*10)
			}
		}()
	}

	// Write data
	for r:=0; r<10; r++{
		go func (){
			for  {
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("Readops:", atomic.LoadUint64(&readOps))
	fmt.Println("Writeops:", atomic.LoadUint64(&writeOps))
	fmt.Println("State: ", state)
}