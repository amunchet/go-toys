package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("Timed out!")
		os.Exit(0)
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
