package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	// Basic for loop
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println("Sum:", sum1)

	// For loop with no init and no post statements - while loop equivalent
	sum2 := 1
	for sum2 < 100 {
		sum2 += sum2
	}
	fmt.Println("Sum:", sum2)

	// For loop with no post statement
	sum3 := 0
	for i := 0; i < 10; {
		sum3 += i
		i++
	}
	fmt.Println("Sum:", sum3)

	// For loop with no init statement
	sum4 := 1
	for ; sum4 < 100; sum4 += sum4 {
	}
	fmt.Println("Sum: ", sum4)

	// If condition
	fmt.Println("Square root of -4:", sqrt(-4))
	fmt.Println("Square root of 4:", sqrt(4))

	// If with short statement
	fmt.Println("Min of 3 power 2 and 20 is", pow(3, 2, 20))
	fmt.Println("Min of 3 power 3 and 20 is", pow(3, 3, 20))

	// If else with short statement
	pow2(3, 3, 20)

	// Exercise: Sqrt using Newton method
	fmt.Println("Square roots of 10, 50 and 100 as computed by Newton method:", sqrtNewton(10), sqrtNewton(50), sqrtNewton(100))
	fmt.Println("Square roots of 10, 50 and 100 from standard library:", math.Sqrt(10), math.Sqrt(50), math.Sqrt(100))

	// Switch
	fmt.Print("Go is currently running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}

	// Switch with no condition
	hour := time.Now().Hour()
	fmt.Print("It is currently ")
	switch {
	case hour < 4:
		fmt.Println("Late Night")
	case hour < 8:
		fmt.Println("Early Morning")
	case hour < 12:
		fmt.Println("Morning")
	case hour < 16:
		fmt.Println("Afternoon")
	case hour < 20:
		fmt.Println("Evening")
	case hour < 24:
		fmt.Println("Night")
	}

	// Defer, Panic and Recover
	div(3, 0)
	div(3, 2)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g is greater than or equal to %g\n", v, lim)
	}
	return lim
}

func sqrtNewton(x float64) float64 {
	z := float64(1)
	delta := float64(2)
	for delta > 0.0000000001 {
		tmp := z
		z -= (z*z - x) / (2 * z)
		delta = math.Abs(z - tmp)
	}
	return z
}

func div(x, y int) {
	defer handlePanic()
	if y == 0 {
		panic("Cannot divide an integer by zero")
	} else {
		result := float64(x) / float64(y)
		fmt.Printf("%d divided by %d is %g\n", x, y, result)
	}
}

func handlePanic() {
	aPanic := recover()
	if aPanic != nil {
		fmt.Println("Program raised exception:", aPanic)
	}
}
