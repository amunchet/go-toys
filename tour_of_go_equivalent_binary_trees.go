package main

import (

	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int, isRoot int){
	
	if t.Left != nil{
		Walk(t.Left, ch,0)
	}
	if t.Right != nil{
		Walk(t.Right, ch,0)
	}
	
	ch <- t.Value
	
	if(isRoot == 1){
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool{
	chan1 := make(chan int)
	chan2 := make(chan int)
	
	var slice1 [] int
	var slice2 [] int
	
	go Walk(t1, chan1,1)
	go Walk(t2, chan2,1)
	
	for i:= range chan1 {
		slice1 = append(slice1,i)
	}
	for j:=range chan2{
		slice2 = append(slice2,j)
	}
	
	
	for _,i:=range slice1{
		found := false
		for _,j:=range slice2{
			if i == j{
				found = true
			}
		}
		if found == false{
			return false
		}
		found = false
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch, 1)
	for i:=range ch{
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
