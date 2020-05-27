package main

import (
	"fmt"
	"strings"
)


// Searches through the string for t.  Returns position or -1 if not found

func Index(vs [] string, t string) int{
	for i,v := range vs{
		if v == t{
			return i
		}
	}
	return -1
}

//
