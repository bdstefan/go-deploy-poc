package main

import "fmt"

var n = 1000000
var computeChan = make(chan int, 20)

func computePower(computeChan chan int, n int, power int) {
	//computeChan <- math.Pow(float64(n), float64(power))

	computeChan <- n * power
}

func publish() {
	for i := 2; i < n; i++ {
		go computePower(computeChan, i, 17)
	}
}

func subscribe() {
	defer close(computeChan)
	for i := 2; i < n; i++ {
		fmt.Println(<-computeChan)
	}
}

//Compute delegate with elegance
func Compute() {
	publish()
	subscribe()
}
