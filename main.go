package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Path[1:])

	if err != nil {
		log.Fatal("Dude, I asked your for a valid Int and I got: %s", err)
		os.Exit(255)
	}

	fmt.Fprintf(w, "Prime numbers up to %v.\n", n)
	computePrimes(n)
}

func main() {
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":3030", nil))
	runtime.GOMAXPROCS(4)
	Compute()
}
