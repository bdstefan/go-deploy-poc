package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bdstefan/go-deploy-poc/nosql"
)

//Number struct store the number and its exponent
type Number struct {
	number   string
	exponent string
}

var numbers []Number

func retreiveNumbers(w http.ResponseWriter) {
	var redis = nosql.GetRedisClient()
	keys, err := redis.Do("KEYS", "*").String()

	if err != nil {
		log.Println("Couldn't fetch redis keys")
		os.Exit(1)
	}

	for _, key := range keys {
		fmt.Fprintln(w, key)
	}
}
