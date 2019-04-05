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

func apiListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All powers computed and stored in Redis")
	retreiveNumbers(w)
}

func main() {
	r.HandleFunc("/power/{number}", powerHandler).Methods("GET")
	r.HandleFunc("/liveness", livenessHandler).Methods("GET")

	s := http.Server{
		Addr:    ":3030",
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}
