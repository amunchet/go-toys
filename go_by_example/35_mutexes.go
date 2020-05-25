package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int] int)
	var mutex = &sync.Mutex{}

	var readCount uint64
	var writeCount uint64 

	// Reading and totalling
	for i:=0; i<100; i++{
		go func(){
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				atomic.AddUint64(&readCount, 1)
				time.Sleep(time.Millisecond)
			}

		}()
	}

	// Writing and totalling
	for w := 0; w< 100; w++ {
		go func(){
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			atomic.AddUint64(&writeCount, 1)
			time.Sleep(time.Millisecond)
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("Total reads: ", atomic.LoadUint64(&readCount))
	fmt.Println("Total writes: ", atomic.LoadUint64(&writeCount))

	mutex.Lock()
	fmt.Println("State: ", state)
	mutex.Unlock()
}