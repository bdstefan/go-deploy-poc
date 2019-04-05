package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func powerHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	n, err := strconv.Atoi(params["number"])
	exp, errExp := strconv.Atoi(params["exp"])

	if err != nil || errExp != nil || exp <= 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request. Provide a valid number and exponent."))

		return
	}

	fmt.Fprintln(w, "Compute power for all int numbers up to", n)

	compute(n, exp, w)
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The app is up and running! :)")
}

func apiListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All powers computed and stored in Redis")
	retreiveNumbers(w)
}

func main() {
	r.HandleFunc("/power/{number}/{exp}", powerHandler).Methods("GET")
	r.HandleFunc("/liveness", livenessHandler).Methods("GET")

	s := http.Server{
		Addr:    ":3030",
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}
