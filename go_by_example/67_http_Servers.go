package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

	fmt.Fprintf(w, "let's go!\n")
}

func HelloContext(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("Starting handler")
	defer fmt.Println("Handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Sending!\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server said: ", err)
		intErr := http.StatusInternalServerError
		http.Error(w, err.Error(), intErr)
	}

}

func main() {

	go func() {
		fmt.Println("Start on *:9993...")
		http.HandleFunc("/hello", hello)
		http.HandleFunc("/hellotwo", HelloContext)

		http.ListenAndServe(":9993", nil)
	}()

	time.Sleep(time.Second * 20)
	fmt.Println("Done")
}
