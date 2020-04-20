package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
		if err != nil && cycles < 1 {
			cycles = 5
		}
		lissajous(w, float64(cycles))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
