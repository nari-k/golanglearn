package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	fmt.Printf("http://localhost:8000/?w=2000&h=2000&t=red&b=green&f=black&c=300")
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")

		cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
		if err != nil || cycles < 1 {
			cycles = 5
		}
		surface(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
