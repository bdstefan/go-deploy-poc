package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/bdstefan/go-deploy-poc/nosql"
)

const exp = 5

var computeChan = make(chan string)
var redis = nosql.GetRedisClient()

func publish(n int) {
	for i := 2; i <= n; i++ {
		go computePower(i, exp)
	}
}

func computePower(n int, exp int) {
	value := math.Pow(float64(n), float64(exp))
	key := fmt.Sprintf("%v:%v", n, exp)

	redis.Set(key, value, 300)

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
