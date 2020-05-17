package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before lunch")
	default:
		fmt.Println("Afternoon!")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool!")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Not sure what this type is %T\n", t)

		}
	}

	whatAmI(true)
	whatAmI(7)
	whatAmI("hey")
}
