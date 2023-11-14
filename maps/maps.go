package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	// Maps
	m := make(map[string]Vertex)
	m["Bengaluru"] = Vertex{12.9716, 77.5946}
	fmt.Println("Coordinates of Bengaluru are:", m["Bengaluru"])

	// Map literals
	m2 := map[string]Vertex{
		"Bengaluru":   {12.9716, 77.5946},
		"BITS Pilani": {28.3588, 75.5880},
	}
	fmt.Println(m2)

	// Mutating Maps
	m3 := make(map[string]int)
	m3["Answer"] = 23
	fmt.Println("Answer:", m3["Answer"])
	m3["Answer"] = 32
	fmt.Println("Answer:", m3["Answer"])
	delete(m3, "Answer")
	fmt.Println("Answer:", m3["Answer"])
	answer, exists := m3["Answer"]
	fmt.Println("Answer:", answer, "Present?", exists)

	// Exercise: Implement WordCount
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	tokens := strings.Fields(s)
	for _, token := range tokens {
		_, exists := result[token]
		if exists {
			result[token]++
		} else {
			result[token] = 1
		}
	}
	return result
}
