package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Path[1:])
	fmt.Fprintln(w, "Compute power for all ints up to", n)
	compute(n, w)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
