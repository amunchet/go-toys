package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error){
	if arg == 42{
		return -1, errors.New("we don't deal with 42")
	}
	return arg, nil
}

// Define a custom error

type specialError struct{
	arg int
	prob string
}

// Implementing the Error method

func (err *specialError) Error() string{
	return fmt.Sprintf("%d - %s", err.arg, err.prob)
}

// Normal function using the custom Error

func f2(arg int)(int, error){
	if arg == 42{
		return -1, &specialError{42, "BIG NOPE"}
	}
	return arg, nil
}

func main() {
	fmt.Println(f1(12))
	fmt.Println(f1(42))

	fmt.Println(f2(32))
	fmt.Println(f2(42))
}