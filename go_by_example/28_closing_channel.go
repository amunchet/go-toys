package main

import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func(){
		for {
			j, status := <- jobs
			if status{
				fmt.Println("Received: ", j)
			}else{
				fmt.Println("Done")
				done <- true
				return
			}
		}
	}()

	jobs <- 5
	jobs <- 3
	jobs <- 2

	close(jobs)

	fmt.Println("All sent")
	<- done
}