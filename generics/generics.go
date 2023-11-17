package main

import (
	"fmt"
	"strings"
)

func main() {
	// Type parameters
	si := []int{2, 3, 5, 7}
	fmt.Println(Index(si, 3))
	ss := []string{"Pangolier", "Dark Willow", "Earthshaker"}
	fmt.Println(Index(ss, "Earthshaker"))

	// Generic types
	primes := List[int]{val: 2} // init with first prime 2
	primes.append(3)
	primes.append(5)
	primes.append(7)
	fmt.Println("Primes:", primes.String())
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) append(value T) {
	newNode := &List[T]{val: value}
	if l == nil {
		*l = *newNode
		return
	}
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *List[T]) String() string {
	current := l
	var result strings.Builder
	for current != nil {
		result.WriteString(fmt.Sprintf("%v ", current.val))
		current = current.next
	}
	return result.String()
}
