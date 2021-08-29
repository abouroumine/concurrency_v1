package main

import "fmt"

// Simple Example illustrating how to use Channels
// With Goroutines

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	go add(1, 2, c1)
	go multiple(3, 2, c2)
	go sum(c1, c2, c3)
	fmt.Println(<-c3)
}

func add(a, b int, c chan int) {
	c <- a + b
}

func multiple(a, b int, c chan int) {
	c <- a * b
}

func sum(a, b, c chan int) {
	c <- <-a + <-b
}
