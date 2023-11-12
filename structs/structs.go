package main

import "fmt"

func main() {
	// Pointers
	i := 2.0
	p := &i
	fmt.Println("Pointer is referring to", *p)
	*p = 23
	fmt.Println("Pointer is referring to", *p)
	*p = *p / 2
	fmt.Println("Pointer is referring to", *p)

	// Structs
	v := Vertex{2, 1}
	fmt.Println("Vertex's value is", v)
	v.Y = 3
	fmt.Println("Vertex's value is", v)

	// Structs and Pointers
	q := &v
	q.X = 1e2
	fmt.Println("Vertex's value is", v)

	// Struct literals
	var (
		v1 = Vertex{2, 3}
		v2 = Vertex{X: 2}
		v3 = Vertex{}
		r  = &Vertex{2, 3}
	)
	fmt.Println(v1, r, v2, v3)
}

type Vertex struct {
	X, Y int
}
