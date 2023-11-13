package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/pic"
)

func main() {
	// Arrays
	primes := [4]int{2, 3, 5, 7}
	fmt.Println("First four prime numbers are:", primes)

	// Slices
	primes23 := primes[0:2]
	fmt.Println("First two primes are:", primes23)

	// Slices refernce to arrays
	greek := [3]string{"Alpha", "Beta", "Gamma"}
	fmt.Println("First three Greek alphabets in order", greek)
	slice01 := greek[0:2]
	slice01[0] = "Omicron"
	slice01[1] = "Kappa"
	fmt.Println("Three Greek alphabets not in any particular order", greek)

	// Slice literals
	isPrime := []struct {
		i int
		b bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
	}
	fmt.Println("First five numbers and whether they are prime:", isPrime)

	// Slice defaults
	primesAgain := []int{2, 3, 5, 7}
	primesAgain23 := primesAgain[:2]
	primesAgain57 := primesAgain[2:]
	fmt.Println("The first two primes are:", primesAgain23)
	fmt.Println("The next two primes are:", primesAgain57)

	// Slice length and capacity
	primesAgainSlice := primesAgain[:2]
	printSlice(primesAgainSlice)
	primesAgainSlice = primesAgain[:3]
	printSlice(primesAgainSlice)

	// Slice with make
	a := make([]int, 5)
	printSlice(a)
	b := make([]int, 0, 5)
	printSlice(b)
	c := b[:2]
	printSlice(c)

	// Slices of slices - 2D slice
	// Tic Tac Toe
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[0][2] = "O"
	board[2][2] = "X"
	board[1][1] = "O"
	board[2][0] = "X"
	board[1][0] = "O"
	board[2][1] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Slice - append
	primes23Array := [2]int{2, 3}
	primesSlice23 := primes23Array[:]
	primesSlice2357 := append(primesSlice23, 5, 7)
	printSlice(primesSlice2357)
	printSlice(primes23Array[:])

	// Range
	for i, prime := range primesSlice2357 {
		fmt.Printf("Prime no. %d is %d\n", i+1, prime)
	}
	for _, prime := range primesSlice2357 {
		fmt.Println("Prime:", prime)
	}
	for i := range primesSlice2357 {
		fmt.Println("Prime:", primesSlice2357[i])
	}

	// Exercise: 8-bit Pic using Slices
	pic.Show(Pic)
}

func printSlice(slice []int) {
	fmt.Printf("Length: %d, Capacity: %d, Elements: %v\n", len(slice), cap(slice), slice)
}

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := range result {
		result[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			result[i][j] = uint8(i ^ j)
		}
	}
	return result
}
