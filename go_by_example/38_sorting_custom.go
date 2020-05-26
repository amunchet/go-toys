package main

import (
	"fmt"
	"sort"
)

type byLength [] string // Define custom type to implement sort interface on

// These are the implements needed for the sort interface
func (s byLength) Len() int{
	return len(s)
}

func (s byLength) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i,j int) bool{
	return len(s[i]) > len(s[j])
}

func main() {
	fruits := []string{"test", "meow", "longerstring", "longeststringpossible"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}