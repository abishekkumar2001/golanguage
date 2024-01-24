package main

import (
	"fmt"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello cat!")
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/cat", h1)
	http.ListenAndServe(":9004", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
