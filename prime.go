package main

import (
	"fmt"
	"math"
)

func initChannel(limit int, ch chan<- int) {
	for i := 2; i <= limit; i++ {
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int, i int) {
	for {
		n := <-in
		if n%i != 0 {
			out <- n
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func compute(limit int) {
	ch := make(chan int)      // Create a new channel.
	go initChannel(limit, ch) // Launch Generate goroutine.
	for i := 0; i < limit; i++ {
		n := <-ch
		fmt.Println(n)
		ch1 := make(chan int)
		go filter(ch, ch1, n)
		ch = ch1
	}
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
		if n%i == 0 {
			return false
		}
	}

	return n > 1
}
