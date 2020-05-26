package main

import (
	"fmt"
	"sort"
)

func main(){
	strs := [] string { "c", "a", "b", "f", "e"}
	sort.Strings(strs)
	fmt.Println(strs)

	fmt.Println("Check if the string are sorted: ", sort.StringsAreSorted(strs))
}