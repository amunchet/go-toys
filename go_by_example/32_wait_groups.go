package main

import(
	"fmt"
	"sync"
	"time"
)

func worker(id int, work_group *sync.WaitGroup){
	defer work_group.Done()

	fmt.Println("Starting ", id)
	time.Sleep(1 * time.Second)
	fmt.Println("Done ", id)
}

func main() {
	var wait_group sync.WaitGroup

	for i:= 0; i<5; i++{
		wait_group.Add(1)
		go worker(i, &wait_group)
	}
	wait_group.Wait()
}