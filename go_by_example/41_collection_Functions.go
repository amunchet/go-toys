package main

import (
	"fmt"
)

// Making my own type

type MyString [] string

// Searches through the string for t.  Returns position or -1 if not found

func (s MyString) Index(t string) int{
	for i,v := range s{
		if v == t{
			return i
		}
	}
	return -1
}

// Include 
func (s MyString) Include(t string) bool{
	return s.Index(t) >= 0
}

// Any

func (s MyString) Any (f func(string) bool) bool{
	for _, v := range s{
		if f(v){
			return true
		}
	}
	return false
}

// All

// Filter

// Map

func main(){
	var a = MyString{"test", "yes", "meow"}
	fmt.Println(a.Index("test"))
	fmt.Println(a.Index("yes"))
	fmt.Println(a.Index("pet"))
}