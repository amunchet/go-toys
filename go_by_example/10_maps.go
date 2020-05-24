package main

// Maps are basically associative arrays

import ("fmt")

func main() {
	m := make(map[string]int)

	m["adfsad"] = 234
	m["ad"] = 234
	fmt.Println(m)

	fmt.Println(len(m))

	// Check to see if a key is present in the map

	_, found := m["missing"]

	fmt.Println(found)

	_,found = m["ad"]
	fmt.Println(found)

	// Declare a map
	n := map[string]int{"test" : 5}
	fmt.Println(n)
}