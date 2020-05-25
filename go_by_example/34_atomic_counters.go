package main

import(
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count uint64

	var weego sync.WaitGroup

	for i:=0; i<50; i++{
		weego.Add(1)
		go func(){
			for j:=0; j<1000; j++{

				/*
				if atomic.LoadUint64(&count) > 5000{
					break
				}
				*/

				atomic.AddUint64(&count,1)
				//count += 1 - This causes a race
			}
			defer weego.Done()
		}()
	}
	weego.Wait()

	fmt.Println("Final count: ", count)
}

