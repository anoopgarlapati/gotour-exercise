package main

import (
	"fmt"
	"math"
)

func main() {
	// Methods
	v := Vertex{2, 3}
	fmt.Println(v.Abs())

	// Methods on non-struct types
	f := MyFloat(-math.Sqrt(23))
	fmt.Printf("MyFloat: %v, Absolute value of MyFloat: %g\n", f, f.Abs())

	// Pointer receivers in Methods
	v.Scale(10)
	fmt.Println(v.Abs())

	// Pointers and Functions
	v2 := Vertex{2, 3}
	Scale(&v2, 10)
	fmt.Println(Abs(v2))

	// Methods and pointer indirection
	p := &v
	fmt.Println(p.Abs())
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
