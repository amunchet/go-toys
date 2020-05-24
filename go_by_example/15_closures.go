package main

import ("fmt")

func spooky () func() int{
	i := 0
	return func() int{
		i += 4
		return i
	}
}

func main() {

	// Closure - the state of i is retained across multiple function calls

	// I guess it's retained in the context of tempSpooky
	tempSpooky := spooky()

	fmt.Println(tempSpooky())
	fmt.Println(tempSpooky())
	fmt.Println(tempSpooky())
	fmt.Println(tempSpooky())
	fmt.Println(tempSpooky())
	fmt.Println(tempSpooky())
}