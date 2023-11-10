package main

import (
	"fmt"
	"math"
	"math/rand"
)

var globalA, globalB int = 2, 3

func main() {
	// Packages
	fmt.Println("A random number is", rand.Intn(99))

	// Imports
	fmt.Printf("Square root of %d is %g\n", 10, math.Sqrt(10))

	// Exported names
	fmt.Println("Pi is an export from math package", math.Pi)

	// Functions
	fmt.Println("Sum of 2 and 3 is", add(2, 3))

	// Functions: Multiple results
	swap1, swap2 := swap("ABC", "XYZ")
	fmt.Println("Swapping ABC XYZ gives us", swap1, swap2)

	// Functions: Named return values
	splitX, splitY := split(3)
	fmt.Println("Splitting 3 into sum of two integers gives", splitX, splitY)

	// Variables
	var localBool, localString = true, "No!!"
	shortBool := false
	fmt.Println("The following are values of global variables", globalA, globalB)
	fmt.Println("The following are values of local variables", localBool, localString, shortBool)

	// Basic types
	var (
		anInt    int        = 23
		aUInt    uint       = 1<<64 - 1
		aBool    bool       = true
		aFloat   float64    = math.Sqrt(23)
		aComplex complex128 = 2 + 3i
		aString  string     = "23"
	)
	fmt.Printf("%v is of type %T\n", anInt, anInt)
	fmt.Printf("%v is of type %T\n", aUInt, aUInt)
	fmt.Printf("%v is of type %T\n", aBool, aBool)
	fmt.Printf("%v is of type %T\n", aFloat, aFloat)
	fmt.Printf("%v is of type %T\n", aComplex, aComplex)
	fmt.Printf("%v is of type %T\n", aString, aString)

	// Numeric type conversions
	fmt.Printf("%v is of type %T\n", int(aFloat), int(aFloat))
	fmt.Printf("%v is of type %T\n", float64(anInt), float64(anInt))
	fmt.Printf("%v is of type %T\n", uint(aFloat), uint(aFloat))

	// Constants
	const (
		Big   = 1 << 23
		Small = Big >> 22
	)
	fmt.Printf("Small %v is of type %T\n", needInt(Small), needInt(Small))
	fmt.Printf("Small %v is of type %T\n", needFloat(Small), needFloat(Small))
	fmt.Printf("Big %v is of type %T\n", needInt(Big), needInt(Big))
	fmt.Printf("Big %v is of type %T\n", needFloat(Big), needFloat(Big))
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum - 1
	y = 1
	return
}

func needInt(x int) int {
	return x
}

func needFloat(x float64) float64 {
	return x
}
