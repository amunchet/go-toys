package main

import (
	"fmt"
	"time"
)

func worker(id int,jobs <- chan int,results chan <- int){
	for j := range jobs {
		fmt.Println("Worker: ", id, " - ", j, ": starting")
		time.Sleep(time.Second)
		fmt.Println("Worker: ", id, " - ", j, ": ending")
		results <- j * 5
	}
}

func main() {

	const jobsCount = 15
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)

	// Start up the worker pool
	for w:=0; w<5; w++{
		go worker(w, jobs, results)
	}

	// Send to jobs to jobs channel
	for j:=0;j<jobsCount; j++{
		jobs <- j
	}

	close(jobs)
	fmt.Println("Sent all the data")

	//Read results back
	for j := 0; j< jobsCount; j++ {
		fmt.Println("Read back: ", j)
		<- results
	}
	
}