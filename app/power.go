package main

import (
	"fmt"
	"math"
	"net/http"
)

const exp = 2

var computeChan = make(chan string)

func publish(n int) {
	for i := 2; i <= n; i++ {
		go computePower(i, exp)
	}
}

func computePower(n int, exp int) {
	value := math.Pow(float64(n), float64(exp))
	computeChan <- fmt.Sprintf("%v ^ %v = %v", n, exp, value)
}

func displayOutput(n int, w http.ResponseWriter) {
	for i := 2; i <= n; i++ {
		result := <-computeChan
		fmt.Fprintln(w, result)
	}
}

func compute(n int, w http.ResponseWriter) {
	publish(n)
	displayOutput(n, w)
}
