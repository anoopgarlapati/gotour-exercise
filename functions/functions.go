package main

import (
	"fmt"
	"math"
)

func main() {
	// Function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("Hypotenuse of right angle triange with sides of length 2 and 3 is:", compute(hypot))
	fmt.Println("2 raised to the power 3 is:", compute(math.Pow))

	// Function closures
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	// Exercise: Pingala series / Fibonacci closure
	ping := pingala()
	for i := 0; i < 10; i++ {
		fmt.Println(ping())
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(2, 3)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func pingala() func() int {
	curr, next := 0, 1
	return func() int {
		tmp := curr
		curr = next
		next += tmp
		return tmp
	}
}
