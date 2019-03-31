package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func powerHandler(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")
	n, err := strconv.Atoi(s[2])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Bad request: %s", err)))
		return
	}

	fmt.Fprintln(w, "Compute power for all int numbers up to", n)

	compute(n, w)
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The app is up and running! :)")
}

func main() {
	http.HandleFunc("/power/", powerHandler)
	http.HandleFunc("/liveness", livenessHandler)

	log.Fatal(http.ListenAndServe(":3030", nil))
}
