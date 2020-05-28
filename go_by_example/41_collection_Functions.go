package main

import (
	"fmt"
)

// Making my own type

type MyString []string

// Searches through the string for t.  Returns position or -1 if not found

func (s MyString) Index(t string) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}
	return -1
}

// Include
func (s MyString) Include(t string) bool {
	return s.Index(t) >= 0
}

// Any

func (s MyString) Any(f func(string) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// All
func (s MyString) All(f func(string) bool) bool {
	for _, v := range s {
		if f(v) {
			return false
		}
	}
	return true
}

// Filter
func (s MyString) Filter(f func(string) bool) MyString {
	retval := make([]string, 0)
	for _, i := range s {
		if f(i) {
			retval = append(retval, i)
		}
	}
	return retval
}

// Map
func (s MyString) Map(f func(string) string) MyString {
	retval := make([]string, len(s))
	for i, v := range s {
		retval[i] = f(v)
	}
	return retval
}

func main() {
	var a = MyString{"test", "yes", "meow"}
	fmt.Println(a.Index("test"))
	fmt.Println(a.Index("yes"))
	fmt.Println(a.Index("pet"))
	// Include
	fmt.Println(a.Include("test"))
	fmt.Println(a.Include("pet"))
	// Any
	fmt.Println(a.Any(func(x string) bool {
		return x == "test"
	}))
	// All

	fmt.Println(a.All(func(x string) bool {
		return x == "test"
	}))
	// Filter
	fmt.Println(a.Filter(func(x string) bool {
		return x == "test"
	}))
	// Map

	fmt.Println(a.Map(func(x string) string {
		return x + "test"
	}))
}
