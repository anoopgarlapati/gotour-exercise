package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	// Walk
	t1, t2 := tree.New(1), tree.New(1)
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	fmt.Println("Tree 1:", t1)
	fmt.Println(t1)
	for v := range c1 {
		fmt.Println(v)
	}
	fmt.Println("Tree 2:", t2)
	for v := range c2 {
		fmt.Println(v)
	}

	// Same
	fmt.Println("Tree(1) and Tree(1) are same?")
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println("Tree(1) and Tree(2) are same?")
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

func Walk(tr *tree.Tree, c chan int) {
	defer close(c)
	walkInt(tr, c)
}

func walkInt(tr *tree.Tree, c chan int) {
	if tr.Left != nil {
		walkInt(tr.Left, c)
	}
	c <- tr.Value
	if tr.Right != nil {
		walkInt(tr.Right, c)
	}
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}
