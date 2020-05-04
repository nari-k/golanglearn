package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	fmt.Printf("http://localhost:8000/?w=1000&h=800&co=red")
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")

		p := new(params)
		wid, err := strconv.Atoi(r.URL.Query().Get("w"))
		if err != nil || wid < 1 {
			wid = width
		}
		p.width = float64(wid)
		hi, err := strconv.Atoi(r.URL.Query().Get("h"))
		if err != nil || wid < 1 {
			hi = height
		}
		p.height = float64(hi)
		co := r.URL.Query().Get("co")
		if len(co) == 0 {
			co = "#ffffff"
		}
		p.color = co

		surface(w, p)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
