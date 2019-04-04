package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/bdstefan/go-deploy-poc/nosql"
)

const exp = 5

var computeChan = make(chan string)
var redis = nosql.GetRedisClient()
var logFile, logErr = os.OpenFile("logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

func publish(n int) {
	for i := 2; i <= n; i++ {
		go computePower(i, exp)
	}
}

func computePower(n int, exp int) {
	key := fmt.Sprintf("%v:%v", n, exp)
	rValue, _ := redis.Get(key).Result()
	value, _ := strconv.Atoi(rValue)

	if value == 0 {
		value = int(math.Pow(float64(n), float64(exp)))
		log.Println(redis.Set(key, value, 0))
	}

	computeChan <- fmt.Sprintf("%v ^ %v = %v", n, exp, value)
}

func displayOutput(n int, w http.ResponseWriter) {
	for i := 2; i <= n; i++ {
		result := <-computeChan
		fmt.Fprintln(w, result)
	}
}

func compute(n int, w http.ResponseWriter) {
	if logErr != nil {
		log.Panic("Log file couldn't be opened.")
		os.Exit(1)
	}

	log.SetOutput(logFile)
	publish(n)
	displayOutput(n, w)
}
