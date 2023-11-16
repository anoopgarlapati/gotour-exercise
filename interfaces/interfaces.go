package main

import (
	"fmt"
	"math"
)

func main() {
	// Interfaces
	var a Abser
	f := MyFloat(-math.Sqrt(23))
	a = f
	fmt.Println(a.Abs())
	v := Vertex{2, 3}
	a = &v
	fmt.Println(a.Abs())

	// Interface values
	a = f
	describeAbser(a)
	a = &v
	describeAbser(a)

	// Interface values with underlying nil values
	var t *T
	var i I
	i = t
	describeI(i)
	i.M()
	i = &T{"Hello World !!"}
	describeI(i)
	i.M()

	// Empty interface
	describe(a)
	describe(i)

	// Type assertions
	var i2 interface{} = "Hello World !!"
	s := i2.(string)
	fmt.Println(s)
	s, ok := i2.(string)
	fmt.Println(s, ok)
	f2, ok2 := i2.(float64)
	fmt.Println(f2, ok2)
	typeAssertionPanicTest()

	// Type switches
	do(23)
	do("Hello World!!")
	do(true)
}

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return math.Abs(float64(f))
	}
	return float64(f)
}

func describeAbser(a Abser) {
	fmt.Printf("%v of type %T\n", a, a)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describeI(i I) {
	fmt.Printf("%v of type %T\n", i, i)
}

func describe(i interface{}) {
	fmt.Printf("%v of type %T\n", i, i)
}

func handlePanic() {
	aPanic := recover()
	if aPanic != nil {
		fmt.Println("Program raised an exception:", aPanic)
	}
}

func typeAssertionPanicTest() {
	defer handlePanic()
	var i interface{} = "Hello World!!"
	f := i.(float64) // throws a panic
	fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Twice of", v, "is", v*2)
	case string:
		fmt.Println(v, "is", len(v), "bytes long")
	default:
		fmt.Println(v, "is of unrecognized type")
	}
}
