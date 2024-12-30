package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutines
	go say("world")
	say("hello")

	time.Sleep(500 * time.Millisecond)

	// Channels
	c := make(chan int)
	s := []int{2, 3, 5, 7, 9, 11}
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	// Buffered Channels
	bc := make(chan int, 2)
	bc <- 2
	bc <- 3
	fmt.Println(<-bc, <-bc)

	// Close
	close(c)
	_, ok := <-c
	fmt.Println("Channel status:", ok)

	// Range and Close
	pc := make(chan int, 10)
	go pingala(cap(pc), pc)
	for i := range pc {
		fmt.Println(i)
	}

	// Select
	pc2, q := make(chan int), make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-pc2)
		}
		q <- 0
	}()
	pingala2(pc2, q)

	// Default Selection
	tick, boom := time.Tick(100*time.Millisecond), time.After(300*time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func pingala(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func pingala2(c, q chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-q:
			fmt.Println("Completed the requested Pingala series")
			return
		}
	}
}
