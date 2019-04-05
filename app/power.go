package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bdstefan/go-deploy-poc/nosql"
)

var computeChan = make(chan string)
var redis = nosql.GetRedisClient()
var logFile, logErr = os.OpenFile("logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

func publish(n int, exp int) {
	for i := 2; i <= n; i++ {
		go computePower(Exponent{base: i, exponent: exp})
	}
}

func computePower(e Exponent) {
	key := fmt.Sprintf("%v:%v", e.base, e.exponent)
	rValue, _ := redis.Get(key).Result()
	value, _ := strconv.Atoi(rValue)

	if value == 0 {
		value = e.power()
		log.Println(redis.Set(key, value, 0))
	}

	computeChan <- fmt.Sprintf("%v ^ %v = %v", e.base, e.exponent, value)
}

func displayOutput(n int, w http.ResponseWriter) {
	for i := 2; i <= n; i++ {
		result := <-computeChan
		fmt.Fprintln(w, result)
	}
}

func compute(n int, exp int, w http.ResponseWriter) {
	if logErr != nil {
		panic("Log file couldn't be opened.")
	}

	log.SetOutput(logFile)

	publish(n, exp)
	displayOutput(n, w)
}
