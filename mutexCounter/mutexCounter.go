package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Mutex Counter
	counter := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go counter.Inc("dio")
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Value of dio:", counter.Value("dio"))
}

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (counter *SafeCounter) Inc(key string) {
	counter.mu.Lock()
	counter.v[key]++
	counter.mu.Unlock()
}

func (counter *SafeCounter) Value(key string) int {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	return counter.v[key]
}
