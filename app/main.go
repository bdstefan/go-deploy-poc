package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

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

func apiListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All powers computed and stored in Redis")
	retreiveNumbers(w)
}

func main() {
	//http.HandleFunc("/power/", powerHandler)
	//http.HandleFunc("/liveness", livenessHandler)
	//http.HandleFunc("/api/list", apiListHandler)

	//log.Fatal(http.ListenAndServe(":3030", nil))

	s := http.Server{
		Addr:    ":3030",
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}
