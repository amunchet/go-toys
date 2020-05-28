package main

import (
	"fmt"
	"testing"
)

func TestIntLess(t *testing.T) {
	a := IntLess(3, 4)
	if !a {
		t.Error("IntLess(3,4) = ", a, "; wanted true")
	}
}

func TestIntLessMany(t *testing.T) {
	var tests = []struct {
		a, b int
		want bool
	}{
		{1, 1, false},
		{1, 2, true},
		{2, 1, false},
		{-5, 1, true}}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d/%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			a := IntLess(tt.a, tt.b)
			if a != tt.want {
				t.Error("got: ", a, " when we wanted: ", tt.want)
			}
		})
	}
}
